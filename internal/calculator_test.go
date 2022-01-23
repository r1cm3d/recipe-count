package internal

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	cases := []struct {
		name  string
		inRec []Record
		inFil Filter
		want  Aggregation
	}{
		{"Happy path", happyPathRecords, regularFilter, happyPathAggregation},
		{"Empty", emptyRecords, regularFilter, emptyAggregation},
		{"Tied post codes", tiedPostcodesRecords, regularFilter, tiedPostCodesAggregation},
		{"Not found names", happyPathRecords, notFoundNamesFilter, notFoundNamesAggregation},
		{"Not found postcode", happyPathRecords, notFoundPostcodeFilter, notFoundPostcodeAggregation},
		{"Invalid range filter", happyPathRecords, invalidRangeFilter, invalidRangeAggregation},
		{"Duplicated name matches", duplicatedNameMatchesRecords, regularFilter, duplicatedNameMatchesAggregation},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			summaryCalculator := NewSummaryCalculator(c.inFil)

			for _, r := range c.inRec {
				summaryCalculator.Calculate(r)
			}
			got := summaryCalculator.Aggregate()

			if !reflect.DeepEqual(c.want, got) {
				t.Errorf("%s, want: %v, got: %v", c.name, c.want, got)
			}
		})
	}
}

var (
	regularFilter = Filter{
		Postcode:  "10120",
		TimeRange: "10AM - 3PM",
		Recipes:   []string{"Potato", "Veggie", "Mushroom"},
	}
	happyPathAggregation = Aggregation{
		UniqueRecipeName: 9,
		RecipeCount: []RecipeCount{
			{"American One-Pan Mushroom", 1},
			{"Cherry Balsamic Pork Chops", 2},
			{"Chicken Sausage Pizzas", 1},
			{"Creamy Dill Chicken", 1},
			{"Grilled Cheese and Veggie Jumble", 1},
			{"Hot Honey Barbecue Chicken Legs", 1},
			{"One-Pan Orzo Italiano", 1},
			{"Speedy Steak Fajitas", 1},
			{"Tex-Mex Tilapia", 1},
		},
		BusiestPostcode: BusiestPostcode{
			"10224",
			2,
		},
		PostcodeAndTimeCount: PostcodeAndTimeCount{
			Postcode:      "10120",
			From:          "10AM",
			To:            "3PM",
			DeliveryCount: 1,
		},
		NameMatches: []string{"American One-Pan Mushroom", "Grilled Cheese and Veggie Jumble"},
	}
	happyPathRecords = []Record{
		{
			Postcode: "10224",
			Recipe:   "Creamy Dill Chicken",
			Delivery: "Wednesday 1AM - 7PM",
		},
		{
			Postcode: "10224",
			Recipe:   "Speedy Steak Fajitas",
			Delivery: "Thursday 7AM - 5PM",
		},
		{
			Postcode: "10120",
			Recipe:   "Cherry Balsamic Pork Chops",
			Delivery: "Thursday 10AM - 2PM",
		},
		{
			Postcode: "10186",
			Recipe:   "Cherry Balsamic Pork Chops",
			Delivery: "Saturday 1AM - 8PM",
		},
		{
			Postcode: "10163",
			Recipe:   "Hot Honey Barbecue Chicken Legs",
			Delivery: "Wednesday 7AM - 5PM",
		},
		{
			Postcode: "10213",
			Recipe:   "Tex-Mex Tilapia",
			Delivery: "Friday 8AM - 7PM",
		},
		{
			Postcode: "10137",
			Recipe:   "One-Pan Orzo Italiano",
			Delivery: "Wednesday 4AM - 7PM",
		},
		{
			Postcode: "10180",
			Recipe:   "Chicken Sausage Pizzas",
			Delivery: "Saturday 6AM - 7PM",
		},
		{
			Postcode: "10127",
			Recipe:   "Grilled Cheese and Veggie Jumble",
			Delivery: "Saturday 8AM - 1PM",
		},
		{
			Postcode: "10148",
			Recipe:   "American One-Pan Mushroom",
			Delivery: "Saturday 10AM - 4PM",
		}}
	emptyAggregation = Aggregation{
		UniqueRecipeName:     0,
		RecipeCount:          nil,
		BusiestPostcode:      BusiestPostcode{},
		PostcodeAndTimeCount: PostcodeAndTimeCount{},
		NameMatches:          nil,
	}
	emptyRecords             []Record
	tiedPostCodesAggregation = Aggregation{
		UniqueRecipeName: 9,
		RecipeCount: []RecipeCount{
			{"American One-Pan Mushroom", 1},
			{"Cherry Balsamic Pork Chops", 2},
			{"Chicken Sausage Pizzas", 1},
			{"Creamy Dill Chicken", 1},
			{"Grilled Cheese and Veggie Jumble", 1},
			{"Hot Honey Barbecue Chicken Legs", 1},
			{"One-Pan Orzo Italiano", 1},
			{"Speedy Steak Fajitas", 1},
			{"Tex-Mex Tilapia", 1},
		},
		BusiestPostcode: BusiestPostcode{
			"10224",
			2,
		},
		PostcodeAndTimeCount: PostcodeAndTimeCount{
			Postcode:      "10120",
			From:          "10AM",
			To:            "3PM",
			DeliveryCount: 1,
		},
		NameMatches: []string{"American One-Pan Mushroom", "Grilled Cheese and Veggie Jumble"},
	}
	tiedPostcodesRecords = []Record{
		{
			Postcode: "10224",
			Recipe:   "Creamy Dill Chicken",
			Delivery: "Wednesday 1AM - 7PM",
		},
		{
			Postcode: "10224",
			Recipe:   "Speedy Steak Fajitas",
			Delivery: "Thursday 7AM - 5PM",
		},
		{
			Postcode: "10120",
			Recipe:   "Cherry Balsamic Pork Chops",
			Delivery: "Thursday 10AM - 2PM",
		},
		{
			Postcode: "10186",
			Recipe:   "Cherry Balsamic Pork Chops",
			Delivery: "Saturday 1AM - 8PM",
		},
		{
			Postcode: "10163",
			Recipe:   "Hot Honey Barbecue Chicken Legs",
			Delivery: "Wednesday 7AM - 5PM",
		},
		{
			Postcode: "999999",
			Recipe:   "Tex-Mex Tilapia",
			Delivery: "Friday 8AM - 7PM",
		},
		{
			Postcode: "999999",
			Recipe:   "One-Pan Orzo Italiano",
			Delivery: "Wednesday 4AM - 7PM",
		},
		{
			Postcode: "10180",
			Recipe:   "Chicken Sausage Pizzas",
			Delivery: "Saturday 6AM - 7PM",
		},
		{
			Postcode: "10127",
			Recipe:   "Grilled Cheese and Veggie Jumble",
			Delivery: "Saturday 8AM - 1PM",
		},
		{
			Postcode: "10148",
			Recipe:   "American One-Pan Mushroom",
			Delivery: "Saturday 10AM - 4PM",
		}}
	notFoundNamesFilter = Filter{
		Postcode:  "10120",
		TimeRange: "10AM - 3PM",
		Recipes:   []string{"Beans", "Rice"},
	}
	notFoundNamesAggregation = Aggregation{
		UniqueRecipeName: 9,
		RecipeCount: []RecipeCount{
			{"American One-Pan Mushroom", 1},
			{"Cherry Balsamic Pork Chops", 2},
			{"Chicken Sausage Pizzas", 1},
			{"Creamy Dill Chicken", 1},
			{"Grilled Cheese and Veggie Jumble", 1},
			{"Hot Honey Barbecue Chicken Legs", 1},
			{"One-Pan Orzo Italiano", 1},
			{"Speedy Steak Fajitas", 1},
			{"Tex-Mex Tilapia", 1},
		},
		BusiestPostcode: BusiestPostcode{
			"10224",
			2,
		},
		PostcodeAndTimeCount: PostcodeAndTimeCount{
			Postcode:      "10120",
			From:          "10AM",
			To:            "3PM",
			DeliveryCount: 1,
		},
		NameMatches: nil,
	}
	notFoundPostcodeFilter = Filter{
		Postcode:  "666",
		TimeRange: "10AM - 3PM",
		Recipes:   []string{"Potato", "Veggie", "Mushroom"},
	}
	notFoundPostcodeAggregation = Aggregation{
		UniqueRecipeName: 9,
		RecipeCount: []RecipeCount{
			{"American One-Pan Mushroom", 1},
			{"Cherry Balsamic Pork Chops", 2},
			{"Chicken Sausage Pizzas", 1},
			{"Creamy Dill Chicken", 1},
			{"Grilled Cheese and Veggie Jumble", 1},
			{"Hot Honey Barbecue Chicken Legs", 1},
			{"One-Pan Orzo Italiano", 1},
			{"Speedy Steak Fajitas", 1},
			{"Tex-Mex Tilapia", 1},
		},
		BusiestPostcode: BusiestPostcode{
			"10224",
			2,
		},
		PostcodeAndTimeCount: PostcodeAndTimeCount{},
		NameMatches:          []string{"American One-Pan Mushroom", "Grilled Cheese and Veggie Jumble"},
	}
	invalidRangeFilter = Filter{
		Postcode:  "10120",
		TimeRange: "",
		Recipes:   []string{"Potato", "Veggie", "Mushroom"},
	}
	invalidRangeAggregation = Aggregation{
		UniqueRecipeName: 9,
		RecipeCount: []RecipeCount{
			{"American One-Pan Mushroom", 1},
			{"Cherry Balsamic Pork Chops", 2},
			{"Chicken Sausage Pizzas", 1},
			{"Creamy Dill Chicken", 1},
			{"Grilled Cheese and Veggie Jumble", 1},
			{"Hot Honey Barbecue Chicken Legs", 1},
			{"One-Pan Orzo Italiano", 1},
			{"Speedy Steak Fajitas", 1},
			{"Tex-Mex Tilapia", 1},
		},
		BusiestPostcode: BusiestPostcode{
			"10224",
			2,
		},
		PostcodeAndTimeCount: PostcodeAndTimeCount{},
		NameMatches:          []string{"American One-Pan Mushroom", "Grilled Cheese and Veggie Jumble"},
	}
	duplicatedNameMatchesRecords = []Record{
		{
			Postcode: "10224",
			Recipe:   "Creamy Dill Chicken",
			Delivery: "Wednesday 1AM - 7PM",
		},
		{
			Postcode: "10224",
			Recipe:   "Speedy Steak Fajitas",
			Delivery: "Thursday 7AM - 5PM",
		},
		{
			Postcode: "10120",
			Recipe:   "Cherry Balsamic Pork Chops",
			Delivery: "Thursday 10AM - 2PM",
		},
		{
			Postcode: "10186",
			Recipe:   "Cherry Balsamic Pork Chops",
			Delivery: "Saturday 1AM - 8PM",
		},
		{
			Postcode: "10163",
			Recipe:   "Hot Honey Barbecue Chicken Legs",
			Delivery: "Wednesday 7AM - 5PM",
		},
		{
			Postcode: "10213",
			Recipe:   "Tex-Mex Tilapia",
			Delivery: "Friday 8AM - 7PM",
		},
		{
			Postcode: "10137",
			Recipe:   "One-Pan Orzo Italiano",
			Delivery: "Wednesday 4AM - 7PM",
		},
		{
			Postcode: "10180",
			Recipe:   "Chicken Sausage Pizzas",
			Delivery: "Saturday 6AM - 7PM",
		},
		{
			Postcode: "10127",
			Recipe:   "Grilled Cheese and Veggie Jumble",
			Delivery: "Saturday 8AM - 1PM",
		},
		{
			Postcode: "10148",
			Recipe:   "American One-Pan Mushroom",
			Delivery: "Saturday 10AM - 4PM",
		},
		{
			Postcode: "666",
			Recipe:   "American One-Pan Mushroom",
			Delivery: "Saturday 8AM - 8PM",
		},}

	duplicatedNameMatchesAggregation = Aggregation{
		UniqueRecipeName: 9,
		RecipeCount: []RecipeCount{
			{"American One-Pan Mushroom", 2},
			{"Cherry Balsamic Pork Chops", 2},
			{"Chicken Sausage Pizzas", 1},
			{"Creamy Dill Chicken", 1},
			{"Grilled Cheese and Veggie Jumble", 1},
			{"Hot Honey Barbecue Chicken Legs", 1},
			{"One-Pan Orzo Italiano", 1},
			{"Speedy Steak Fajitas", 1},
			{"Tex-Mex Tilapia", 1},
		},
		BusiestPostcode: BusiestPostcode{
			"10224",
			2,
		},
		PostcodeAndTimeCount: PostcodeAndTimeCount{
			Postcode:      "10120",
			From:          "10AM",
			To:            "3PM",
			DeliveryCount: 1,
		},
		NameMatches: []string{"American One-Pan Mushroom", "Grilled Cheese and Veggie Jumble"},
	}
)
