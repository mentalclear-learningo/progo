package main

import (
	"fmt"
	"strings"
)

func splittingString() {
	description := "A boat for one person"
	splits := strings.Split(description, " ")
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
	splitsAfter := strings.SplitAfter(description, " ")
	for _, x := range splitsAfter {
		fmt.Println("SplitAfter >>" + x + "<<")
	}
}

func restrictResultsNumber() {
	description := "A boat for one person"
	splits := strings.SplitN(description, " ", 3)
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
	// splitsAfter := strings.SplitAfter(description, " ")
	// for _, x := range splitsAfter {
	//     fmt.Println("SplitAfter >>" + x + "<<")
	// }
}

func splittinOnWhitespace() {
	description := "This  is  double  spaced"
	splits := strings.SplitN(description, " ", 3)
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}

	// Fields func:
	splits2 := strings.Fields(description)
	for _, x := range splits2 {
		fmt.Println("Field >>" + x + "<<")
	}
}

func splittingUsingCustom() {
	description := "This  is  double  spaced"
	splitter := func(r rune) bool {
		return r == ' '
	}
	splits := strings.FieldsFunc(description, splitter)
	for _, x := range splits {
		fmt.Println("Field >>" + x + "<<")
	}
}

func main() {
	fmt.Println("Splitting Strings:")
	splittingString()
	fmt.Println("\nRestricting the Number of Results:")
	restrictResultsNumber()
	fmt.Println("\nSplitting on Whitespace Characters:")
	splittinOnWhitespace()
	fmt.Println("\nSplitting Using a Custom Function to Split Strings:")
	splittingUsingCustom()
}
