package internal

import (
	"sort"
	"strings"
)

type (
	// Calculator is the interface that calculates the aggregation given a Record.
	Calculator interface {
		Calculate(r Record)
	}
	// Filter is the information needed to matches PostcodeAndTimeCount and NamesMatches.
	Filter struct {
		Postcode  string
		TimeRange string
		Recipes   []string
	}
	// SummaryCalculator is a single thread implementation of the calculator. It keeps all state into its unexported
	// structures. It MUST NOT be used in concurrent environments without proper synchronization. Besides that, all
	// properties that are used as cache are mutable and might have unpredictable behavior in concurrent environments.
	SummaryCalculator struct {
		Filter
		uniqueRecipesCache   map[string]int
		busiestPostcode      map[string]int
		postcodeAndTimeCount PostcodeAndTimeCount
		nameMatchesCache     []string
	}
)

// NewSummaryCalculator creates SummaryCalculator given a filter. It is important use this function instead
// creating a SummaryCalculator directly.
func NewSummaryCalculator(filter Filter) SummaryCalculator {
	return SummaryCalculator{filter, make(map[string]int), make(map[string]int),
		PostcodeAndTimeCount{}, nil,
	}
}

// Aggregate applies sorting uniqueRecipesCache and get the busiestPostcode applying a count sorting and getting
// the postcode with more appearances in the input JSON file. If two postcode are tied with the same count, it gets the
// one with lower number. E.g. 666 has 6 appearances as 10212 does, it will choose 666.
func (s SummaryCalculator) Aggregate() Aggregation {
	recipeCount := s.sumRecipes()
	busiestPostcode := s.calcBusiestPostCode()
	return Aggregation{
		UniqueRecipeName:     len(s.uniqueRecipesCache),
		RecipeCount:          recipeCount,
		BusiestPostcode:      busiestPostcode,
		PostcodeAndTimeCount: s.postcodeAndTimeCount,
		NameMatches:          s.nameMatchesCache,
	}
}

// Calculate adds Record information in its caches according functional requirements. It was designed to be used
// in a single thread environment and using it in a concurrent environment might causes unpredictable behavior.
func (s *SummaryCalculator) Calculate(r Record) {
	s.uniqueRecipesCache[r.Recipe]++
	s.busiestPostcode[r.Postcode]++
	s.addToNamesMatches(r)
	s.filterRecipeAccordingFilter(r)
}

func (s *SummaryCalculator) addToNamesMatches(r Record) {
	for _, fr := range s.Filter.Recipes {
		lowerFilterRecipe := strings.ToLower(fr)
		lowerRecordRecipe := strings.ToLower(r.Recipe)

		if strings.Contains(lowerRecordRecipe, lowerFilterRecipe) && !s.isNameAlreadyInserted(r.Recipe) {
			s.insertSorted(r.Recipe)
		}
	}
}

// filterRecipeAccordingFilter does not take weekday in consideration to perform the query
func (s *SummaryCalculator) filterRecipeAccordingFilter(r Record) {
	if s.Filter.Postcode != r.Postcode || !r.DeliveredBetween(s.Filter.TimeRange) {
		return
	}

	s.postcodeAndTimeCount.DeliveryCount++
	s.postcodeAndTimeCount.Postcode = r.Postcode

	from, err := GetBeginHourStr(s.Filter.TimeRange)
	if err != nil {
		return
	}
	s.postcodeAndTimeCount.From = from[0]

	to, err := GetEndHourStr(s.Filter.TimeRange)
	if err != nil {
		return
	}
	s.postcodeAndTimeCount.To = to[0]
}

func (s SummaryCalculator) sumRecipes() []RecipeCount {
	keys := make([]string, 0, len(s.uniqueRecipesCache))
	for k := range s.uniqueRecipesCache {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sortedRecipes []RecipeCount
	for _, k := range keys {
		rc := RecipeCount{
			Recipe: k,
			Count:  s.uniqueRecipesCache[k],
		}

		sortedRecipes = append(sortedRecipes, rc)
	}

	return sortedRecipes
}

func (s SummaryCalculator) calcBusiestPostCode() BusiestPostcode {
	var sortedBusiestPostCodes []BusiestPostcode
	for k, v := range s.busiestPostcode {
		sortedBusiestPostCodes = append(sortedBusiestPostCodes, BusiestPostcode{
			Postcode:      k,
			DeliveryCount: v,
		})
	}

	// Sort by most delivery count and then lower postcode
	sort.Slice(sortedBusiestPostCodes, func(i, j int) bool {
		iDC, jDC := sortedBusiestPostCodes[i].DeliveryCount, sortedBusiestPostCodes[j].DeliveryCount
		iPC, jPC := sortedBusiestPostCodes[i].Postcode, sortedBusiestPostCodes[j].Postcode

		return (iDC == jDC && iPC < jPC) || iDC > jDC
	})

	if len(sortedBusiestPostCodes) > 0 {
		return sortedBusiestPostCodes[0]
	}

	return BusiestPostcode{}
}


func (s *SummaryCalculator) insertSorted(name string) {
	i := sort.SearchStrings(s.nameMatchesCache, name)
	s.nameMatchesCache = append(s.nameMatchesCache, "")
	copy(s.nameMatchesCache[i+1:], s.nameMatchesCache[i:])
	s.nameMatchesCache[i] = name
}

func (s *SummaryCalculator) isNameAlreadyInserted(e string) bool {
	for _, a := range s.nameMatchesCache {
		if a == e {
			return true
		}
	}
	return false
}