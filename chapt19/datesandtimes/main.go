package main

import (
	"fmt"
	"time"
)

func PrintTime(label string, t *time.Time) {
	Printfln("%s: Day: %v: Month: %v Year: %v",
		label, t.Day(), t.Month(), t.Year())
}

func timeValues() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	fmt.Println("\ncurrent:", current, "\nspecific:", specific, "\nunix:", unix)
	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)
}

// Changes made for Formattin Times as Strings
func PrintTimeFormated(label string, t *time.Time) {
	// layout := "Day: 02 Month: Jan Year: 2006"
	// fmt.Println(label, t.Format(layout))

	// RFC8222 example
	fmt.Println(label, t.Format(time.RFC822Z))
}

func formatTimes() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTimeFormated("Current", &current)
	PrintTimeFormated("Specific", &specific)
	PrintTimeFormated("UNIX", &unix)
}

func parseTimeValues() {
	layout := "2006-Jan-02"
	dates := []string{
		"1995-Jun-09",
		"2015-Jun-02",
	}
	for _, d := range dates {
		time, err := time.Parse(layout, d)
		if err == nil {
			PrintTimeFormated("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
}

func parsePredefined() {
	//layout := "2006-Jan-02"
	dates := []string{
		"09 Jun 95 00:00 GMT",
		"02 Jun 15 00:00 GMT",
	}
	for _, d := range dates {
		time, err := time.Parse(time.RFC822, d)
		if err == nil {
			PrintTimeFormated("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
}

func specPrasinfLocation() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nycerr := time.LoadLocation("America/New_York")
	if lonerr == nil && nycerr == nil {
		nolocation, _ := time.Parse(layout, date)
		londonTime, _ := time.ParseInLocation(layout, date, london)
		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
		PrintTimeFormated("No location:", &nolocation)
		PrintTimeFormated("London:", &londonTime)
		PrintTimeFormated("New York:", &newyorkTime)
	} else {
		fmt.Println(lonerr.Error(), nycerr.Error())
	}
}

func usingLocalLocation() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nycerr := time.LoadLocation("America/New_York")
	local, _ := time.LoadLocation("Local")
	if lonerr == nil && nycerr == nil {
		nolocation, _ := time.Parse(layout, date)
		londonTime, _ := time.ParseInLocation(layout, date, london)
		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
		localTime, _ := time.ParseInLocation(layout, date, local)
		PrintTimeFormated("No location:", &nolocation)
		PrintTimeFormated("London:", &londonTime)
		PrintTimeFormated("New York:", &newyorkTime)
		PrintTimeFormated("Local:", &localTime)
	} else {
		fmt.Println(lonerr.Error(), nycerr.Error())
	}
}

func directTimeZones() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london := time.FixedZone("BST", 1*60*60)
	newyork := time.FixedZone("EDT", -4*60*60)
	local := time.FixedZone("Local", 0)
	//if (lonerr == nil && nycerr == nil) {
	nolocation, _ := time.Parse(layout, date)
	londonTime, _ := time.ParseInLocation(layout, date, london)
	newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
	localTime, _ := time.ParseInLocation(layout, date, local)
	PrintTimeFormated("No location:", &nolocation)
	PrintTimeFormated("London:", &londonTime)
	PrintTimeFormated("New York:", &newyorkTime)
	PrintTimeFormated("Local:", &localTime)
	// } else {
	//     fmt.Println(lonerr.Error(), nycerr.Error())
	// }
}

func manipTimeVal() {
	t, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
	fmt.Println("Value:", t)
	if err == nil {
		Printfln("After?: %v", t.After(time.Now()))
		Printfln("Before?: %v", t.Before(time.Now()))
		Printfln("Round: %v", t.Round(time.Hour))
		Printfln("Truncate: %v", t.Truncate(time.Hour))
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println("\nMore comparing:")
	t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
	t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")
	Printfln("Equal Method: %v", t1.Equal(t2))
	Printfln("Equality Operator: %v", t1 == t2)
}

func representDurations() {
	var d time.Duration = time.Hour + (30 * time.Minute)
	fmt.Println("Direct variable d:", d)
	Printfln("Hours: %v", d.Hours())
	Printfln("Mins: %v", d.Minutes())
	Printfln("Seconds: %v", d.Seconds())
	Printfln("Millseconds: %v", d.Milliseconds())
	rounded := d.Round(time.Hour)
	Printfln("Rounded Hours: %v", rounded.Hours())
	Printfln("Rounded Mins: %v", rounded.Minutes())
	trunc := d.Truncate(time.Hour)
	Printfln("Truncated  Hours: %v", trunc.Hours())
	Printfln("Rounded Mins: %v", trunc.Minutes())
}

func createDurationsRelativeTime() {
	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}
	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)
	Printfln("Future: %v", toYears(time.Until(future)))
	Printfln("Past: %v", toYears(time.Since(past)))
}

func createDurFromStr() {
	d, err := time.ParseDuration("1h30m")
	if err == nil {
		Printfln("Hours: %v", d.Hours())
		Printfln("Mins: %v", d.Minutes())
		Printfln("Seconds: %v", d.Seconds())
		Printfln("Millseconds: %v", d.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}
}

func main() {
	Printfln("Hello, Dates and Times")

	fmt.Println("\nRepresenting Dates and Times")
	timeValues()

	fmt.Println("\nFormatting time as strings:")
	formatTimes()

	fmt.Println("\nParsing Time Values from Strings:")
	parseTimeValues()

	fmt.Println("\nUsing Predefined Date Layouts:")
	parsePredefined()

	fmt.Println("\nSpecifying a Parsing Location:")
	specPrasinfLocation()

	fmt.Println("\nUsing the Local Location:")
	usingLocalLocation()

	fmt.Println("\nSpecifying Time Zones Directly:")
	directTimeZones()

	fmt.Println("\nManipulating Time Values:")
	manipTimeVal()

	fmt.Println("\nRepresenting Durations:")
	representDurations()

	fmt.Println("\nCreating Durations Relative to a Time:")
	createDurationsRelativeTime()

	fmt.Println("\nCreating Durations from Strings:")
	createDurFromStr()
}
