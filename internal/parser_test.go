package internal

import (
	"os"
	"reflect"
	"testing"
)

func TestParseIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	createFile(fixture)
	defer removeFile()
	mc := mockCalculator{results: []Record{}}
	want := []Record{
		{"10224", "Creamy Dill Chicken", "Wednesday 1AM - 7PM"},
		{"10208", "Speedy Steak Fajitas", "Thursday 7AM - 5PM"},
		{"10120", "Cherry Balsamic Pork Chops", "Thursday 7AM - 9PM"},
	}
	parsedWant := 3
	ignoredWant := 5

	parsedGot, ignoredGot := Parse(stubFile, &mc, true)

	if !reflect.DeepEqual(mc.results, want) || parsedGot != parsedWant || ignoredGot != ignoredWant  {
		t.Errorf("Error at Parse function, resultsWant: %v, resultsGot: %v, parsedWant: %v, parsedGot: %v\", ignoredWant: %v, ignoredGot: %v\"",
			want, mc.results, parsedWant, parsedGot, ignoredWant, ignoredGot)
	}
}

type mockCalculator struct {
	results []Record
}

func (m *mockCalculator) Calculate(r Record) {
	m.results = append(m.results, r)
}

func createFile(content string) {
	f, err := os.Create(stubFile)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
	f.Sync()
}

func removeFile() {
	if err := os.Remove(stubFile); err != nil {
		panic(err)
	}
}

const (
	stubFile = "tf"
	fixture  = `[
{
  "postcode": "10224",
  "recipe": "Creamy Dill Chicken",
  "delivery": "Wednesday 1AM - 7PM"
},
{
  "postcode": "10208",
  "recipe": "Speedy Steak Fajitas",
  "delivery": "Thursday 7AM - 5PM"
},
{
  "postcode": "10120",
  "recipe": "Cherry Balsamic Pork Chops",
  "delivery": "Thursday 7AM - 9PM"
},
{
  "postcode": "10224",
  "recipe": "Creamy Dill Chicken",
  "delivery": "1AM - 7PM"
},
{
  "postcode": "10224",
  "recipe": "KKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK",
  "delivery": "Thursday 1AM - 7PM"
},
{
  "postcode": "10224196412",
  "recipe": "Creamy Dill Chicken",
  "delivery": "Thursday 1AM - 7PM"
},
{
  "postcode": "",
  "recipe": "",
  "delivery": "Thursday 1AM - 7PM"
},
{
  "postcode": "10224",
  "recipe": "Cherry Balsamic Pork Chops",
  "delivery": ""
}]`
)
