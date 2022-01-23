package internal

import (
	"encoding/json"
)

type (
	// NamesMatches groups recipe names that matched filter slice names criteria.
	NamesMatches []string
	// RecipeCount is an abstraction for each distinct recipe that is in the input JSON file.
	RecipeCount struct {
		Recipe string `json:"recipe"`
		Count  int `json:"count"`
	}
	// BusiestPostcode represents the postcode with more appearances in the input JSON file.
	BusiestPostcode struct {
		Postcode      string `json:"postcode"`
		DeliveryCount int `json:"delivery_count"`
	}
	// PostcodeAndTimeCount counts how many times the recipes that matches filter criteria appears in the input JSON file.
	PostcodeAndTimeCount struct {
		Postcode      string `json:"postcode"`
		From          string `json:"from"`
		To            string `json:"to"`
		DeliveryCount int `json:"delivery_count"`
	}
	// Aggregation groups all information needed in output file.
	Aggregation struct {
		UniqueRecipeName     int           `json:"unique_recipe_count"`
		RecipeCount          []RecipeCount `json:"count_per_recipe"`
		BusiestPostcode      `json:"busiest_postcode"`
		PostcodeAndTimeCount `json:"count_per_postcode_and_time"`
		NameMatches          NamesMatches `json:"match_by_name"`
	}
)

func (a Aggregation) String() string {
	str, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(str)
}
