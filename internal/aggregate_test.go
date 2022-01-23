package internal

import (
	"testing"
)

func TestAggregationString(t *testing.T) {
	agg := Aggregation{
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
	want := `{
    "unique_recipe_count": 9,
    "count_per_recipe": [
        {
            "recipe": "American One-Pan Mushroom",
            "count": 1
        },
        {
            "recipe": "Cherry Balsamic Pork Chops",
            "count": 2
        },
        {
            "recipe": "Chicken Sausage Pizzas",
            "count": 1
        },
        {
            "recipe": "Creamy Dill Chicken",
            "count": 1
        },
        {
            "recipe": "Grilled Cheese and Veggie Jumble",
            "count": 1
        },
        {
            "recipe": "Hot Honey Barbecue Chicken Legs",
            "count": 1
        },
        {
            "recipe": "One-Pan Orzo Italiano",
            "count": 1
        },
        {
            "recipe": "Speedy Steak Fajitas",
            "count": 1
        },
        {
            "recipe": "Tex-Mex Tilapia",
            "count": 1
        }
    ],
    "busiest_postcode": {
        "postcode": "10224",
        "delivery_count": 2
    },
    "count_per_postcode_and_time": {
        "postcode": "10120",
        "from": "10AM",
        "to": "3PM",
        "delivery_count": 1
    },
    "match_by_name": [
        "American One-Pan Mushroom",
        "Grilled Cheese and Veggie Jumble"
    ]
}`

	got := agg.String()

	if want != got {
		t.Errorf("Error at Aggregate.String(), want: %v, got: %v", want, got)
	}
}
