package internal

import "testing"

func TestDeliveredBetween(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"Valid time range - YES", "10AM - 3PM", true},
		{"Valid time range - NO", "9AM - 2PM", false},
		{"Invalid time range - EMPTY", "", false},
		{"Invalid time range - GARBAGE", "ZAMBAS", false},
	}
	r := Record{
		Postcode: "10120",
		Recipe:   "Cherry Balsamic Pork Chops",
		Delivery: "10AM - 2PM",
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := r.DeliveredBetween(c.in); got != c.want {
				t.Errorf("%s, want: %v, got: %v", c.name, c.want, got)
			}
		})
	}
}
