package internal

import (
	"regexp"
	"strconv"
)

// Record represents each Record of the delivered recipes list that are into the input JSON file.
type Record struct {
	Postcode string
	Recipe   string
	Delivery string
}

// DeliveredBetween receives a range in the contract format: "{H}AM - {H}PM" and checks if the delivery time range
// of the current Record matches with the input criteria.
// If any format error occurs it returns false.
func (r Record) DeliveredBetween(timeRange string) bool {
	begin, err := r.getBeginHour(timeRange)
	if err != nil {
		return false
	}

	end, err := r.getEndHour(timeRange)
	if err != nil {
		return false
	}

	return r.beginHour() <= begin && r.endHour() <= end
}

// IsValid checks all properties of the current record according functional requirements. If any checked fails
// it returns false.
func (r Record) IsValid() bool {
	recipeRegex := `^[^\s]+\s+(1[0-2]|0?[1-9])[Aa][Mm]\s+\-\s+(1[0-2]|0?[1-9])[Pp][Mm]`
	maxCharacterRecipeAllowed := 100
	maxCharacterPostcodeAllowed := 10
	validDelivery, err := regexp.MatchString(recipeRegex, r.Delivery)
	validRecipe := len(r.Recipe) <= maxCharacterRecipeAllowed && len(r.Recipe) > 0
	validPostcode := len(r.Postcode) <= maxCharacterPostcodeAllowed && len(r.Postcode) > 0

	return  validRecipe &&
		validPostcode &&
		err == nil &&
		validDelivery
}

func (r Record) beginHour() int {
	begin, _ := r.getBeginHour(r.Delivery)

	return begin
}

func (r Record) endHour() int {
	end, _ := r.getEndHour(r.Delivery)

	return end
}

func (r Record) getBeginHour(timeRange string) (int, error) {
	bhStr, err := GetBeginHourStr(timeRange)
	if err != nil {
		return 0, err
	}

	hour, err := strconv.Atoi(bhStr[1])
	if err != nil {
		return 0, err
	}

	return hour, nil
}

func (r Record) getEndHour(timeRange string) (int, error) {
	diffTo24Hour := 12

	endHour, err := GetEndHourStr(timeRange)
	if err != nil {
		return 0, err
	}
	hour, err := strconv.Atoi(endHour[1])
	if err != nil {
		return 0, err
	}

	return hour + diffTo24Hour, nil
}
