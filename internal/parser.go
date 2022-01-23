package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
)

// ConsoleClear is a constant that "cleans" the console. It is used only in verbose mode for debug purposes.
// It was tested in a Linux environment.
const ConsoleClear = "\033[H\033[2J"

// GetBeginHourStr returns a slice with the matched begin hour.
// E.g. "10AM - 3PM":  [0] == "10AM, [1] == "10"
// It returns an error if regexp.Compile fails or re.FindStringSubmatch does not found anything.
func GetBeginHourStr(timeRange string) ([]string, error) {
	re, err := regexp.Compile(`(1[0-2]|0?[1-9])(?:[Aa][Mm])`)
	if err != nil {
		return nil, err
	}

	matches := re.FindStringSubmatch(timeRange)
	if len(matches) <= 0 {
		return nil, fmt.Errorf("begin hour not found at %s", timeRange)
	}

	return matches, nil
}

// GetEndHourStr returns a slice with the matched end hour.
// E.g. "10AM - 3PM":  [0] == "3PM, [0] == "3"
// It returns an error if regexp.Compile fails or re.FindStringSubmatch does not found anything.
func GetEndHourStr(timeRange string) ([]string, error) {
	re, err := regexp.Compile(`(1[0-2]|0?[1-9])(?:[Pp][Mm])`)
	if err != nil {
		return nil, err
	}

	matches := re.FindStringSubmatch(timeRange)
	if len(matches) <= 0 {
		return nil, fmt.Errorf("end hour not found at %s", timeRange)
	}

	return matches, nil
}

// Parse opens a file given filepath, decodes it and apply Calculator.calculate() for each parsed record.
// It panics if any parse error happens. For instance: an invalid JSON.
// It returns parsed and ignored:
// - parsed is a count with all successful parsed records;
// - ignored contains all invalid records that were ignored;
func Parse(filepath string, calc Calculator, isVerbose bool) (parsed int, ignored int) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", filepath, err.Error())
	}

	r := bufio.NewReader(f)
	d := json.NewDecoder(r)

	nextToken(d)
	parsed = 0
	ignored = 0
	i := 0
	for d.More() {
		r := &Record{}
		if err := d.Decode(r); err != nil {
			log.Fatalf("Error to decode %v", err.Error())
		}

		i++
		if !r.IsValid() {
			ignored++
			logCount(isVerbose, i, parsed, ignored)
			continue
		}

		calc.Calculate(*r)
		parsed++
		logCount(isVerbose, i, parsed, ignored)
	}
	nextToken(d)

	return
}

func nextToken(d *json.Decoder) {
	if _, err := d.Token(); err != nil {
		panic(fmt.Errorf("function Token() encountered an unexpected delimiter in the input %v", err.Error()))
	}
}

func logCount(isVerbose bool, rc, pc, ic int) {
	if !isVerbose {
		return
	}

	fmt.Print(ConsoleClear)
	fmt.Printf("Record: %d", rc)
	fmt.Printf("\nParsed: %d", pc)
	fmt.Printf("\nIgnored: %d", ic)
}
