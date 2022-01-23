package main

import (
	"fmt"
	"github.com/hellofreshdevtests/r1cm3d-recipe-count-test-2020/internal"
	"github.com/thatisuday/clapper"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	defaultPostcode  = "10120"
	defaultTimeRange = "10AM - 3PM"
	defaultNames     = "Potato,Veggie,Mushroom"
	filter           internal.Filter
	file             string
	isVerbose        bool
)

//Example of use:
//./recipe-aggregator -f 'test/hf_test_calculation_fixtures.json' -r 'Friday 10AM - 2PM' -p '10021' -n 'Veggie,Potato'
//
//Only filename parameter is required. The filter (postcode, timerange and names) are optional.
//
//List of parameters:
//
//Name      | Type   | Name        | Shortname | Example 									| Required | Default
//Filename  | string | --filename  | -f        | "test/hf_test_calculation_fixtures.json" | true     | NA
//Postcode  | string | --postcode  | -p        | '10021'                                  | false    | '10120'
//Timerange | string | --timerange | -r        | 'Friday 10AM - 2PM'                      | false    | '10AM - 3PM'
//Names     | string | --names     | -n        | 'Veggie,Potato'                          | false    | 'Potato,Veggie,Mushroom'
//Verbose   | flag   | --verbose   | -v        | NA                                       | false    | NA
//Help      | flag   | --help      | -h        | NA                                       | false    | NA
func main() {
	loadArgs()
	start := time.Now()
	if isVerbose {
		fmt.Printf("Input\nFile: %v\nFilter: %v\n", file, filter)
	}

	calculator := internal.NewSummaryCalculator(filter)
	internal.Parse(file, &calculator, isVerbose)
	aggregation := calculator.Aggregate()
	fmt.Printf(internal.ConsoleClear)
	fmt.Println(aggregation)

	duration := time.Since(start)
	if isVerbose {
		fmt.Println(duration)
	}
}

func loadArgs() {
	var m = make(map[string]string)
	const (
		filepath = "filepath"
		postcode = "postcode"
		timeRange = "timerange"
		names = "names"
		verbose = "verbose"
		help = "help"
	)

	registry := clapper.NewRegistry()
	rootCommand, _ := registry.Register("")
	rootCommand.AddFlag(filepath, "f", false, "")
	rootCommand.AddFlag(postcode, "p", false, defaultPostcode)
	rootCommand.AddFlag(timeRange, "r", false, defaultTimeRange)
	rootCommand.AddFlag(names, "n", false, defaultNames)
	rootCommand.AddFlag(verbose, "v", true, "")
	rootCommand.AddFlag(help, "h", true, "")

	command, err := registry.Parse(os.Args[1:])

	if err != nil {
		printHelpAndExit()
	}

	for flagName, flagValue := range command.Flags {
		if flagValue.Value != "" {
			m[flagName] = flagValue.Value
		} else {
			m[flagName] = flagValue.DefaultValue
		}
	}

	if ok, _ := strconv.ParseBool(m[help]); ok {
		printHelpAndExit()
	}

	verb, err := strconv.ParseBool(m[verbose])
	isVerbose = err == nil && verb

	filter = internal.Filter{
		Postcode:  m[postcode],
		TimeRange: m[timeRange],
		Recipes:   strings.Split(m[names], ","),
	}

	file = m[filepath]
	if file == "" {
		printHelpAndExit()
	}
}

func printHelpAndExit() {
	fmt.Println(exampleOfUsage)
	os.Exit(0)
}
const exampleOfUsage = `
Example of use:
	./recipe-aggregator -f 'test/hf_test_calculation_fixtures.json' -r 'Friday 10AM - 2PM' -p '10021' -n 'Veggie,Potato'

Only filename parameter is required. The filter (postcode, timerange and names) are optional.

List of parameters:

Name      | Type   | Name        | Shortname | Example 									| Required | Default  
Filename  | string | --filename  | -f        | "test/hf_test_calculation_fixtures.json" | true     | NA
Postcode  | string | --postcode  | -p        | '10021'                                  | false    | '10120'
Timerange | string | --timerange | -r        | 'Friday 10AM - 2PM'                      | false    | '10AM - 3PM'
Names     | string | --names     | -n        | 'Veggie,Potato'                          | false    | 'Potato,Veggie,Mushroom'
Verbose   | flag   | --verbose   | -v        | NA                                       | false    | NA
Help      | flag   | --help      | -h        | NA                                       | false    | NA`