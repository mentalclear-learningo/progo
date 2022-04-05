package main

import (
	"fmt"
	//"strings"
	"regexp"
)

func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

func withFindStringIndex() {
	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description := "Kayak. A boat for one person."
	firstIndex := pattern.FindStringIndex(description)
	allIndices := pattern.FindAllStringIndex(description, -1)
	fmt.Println("First index", firstIndex[0], "-", firstIndex[1],
		"=", getSubstring(description, firstIndex))
	for i, idx := range allIndices {
		fmt.Println("Index", i, "=", idx[0], "-",
			idx[1], "=", getSubstring(description, idx))
	}

	// Just matches:
	firstMatch := pattern.FindString(description)
	allMatches := pattern.FindAllString(description, -1)
	fmt.Println("First match:", firstMatch)
	for i, m := range allMatches {
		fmt.Println("Match", i, "=", m)
	}
}

func main() {
	description := "A boat for one person"
	match, err := regexp.MatchString("[A-z]oat", description)
	if err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}

	// Compile patter to be reused later
	pattern, compileErr := regexp.Compile("[A-z]oat")
	question := "Is that a goat?"
	preference := "I like oats"
	if compileErr == nil {
		fmt.Println("Description:", pattern.MatchString(description))
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Preference:", pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}

	fmt.Println("\nThe FindStringIndex and FindAllStringIndex methods:")
	withFindStringIndex()

	fmt.Println("\nSplitting Strings Using a Regular Expression:")
	pattern2 := regexp.MustCompile(" |boat|one")
	split := pattern2.Split(description, -1)
	for _, s := range split {
		if s != "" {
			fmt.Println("Substring:", s)
		}
	}

	fmt.Println("\nUsing Subexpressions:")
	pattern3 := regexp.MustCompile("A [A-z]* for [A-z]* person")
	str := pattern3.FindString(description)
	fmt.Println("Match:", str)

	fmt.Println("More sub expressions:")
	pattern4 := regexp.MustCompile("A ([A-z]*) for ([A-z]*) person")
	subs := pattern4.FindStringSubmatch(description)
	for _, s := range subs {
		fmt.Println("Match:", s)
	}

	// fmt.Println("\nUsing named subexpressions:")
	// pattern5 := regexp.MustCompile(
	// 	"A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")

	// subs2 := pattern5.FindStringSubmatch(description)
	// for _, name11 := range []string{"type", "capacity"} {
	// 	fmt.Println(name11, "=", subs2[pattern.SubexpIndex(name11)])
	// }
	fmt.Println("\nReplacing Substrings Using a Regular Expression:")
	pattern6 := regexp.MustCompile(
		"A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description2 := "Kayak. A boat for one person."
	template := "(type: ${type}, capacity: ${capacity})"
	replaced := pattern6.ReplaceAllString(description2, template)
	fmt.Println(replaced)

	fmt.Println("\nReplacing Matched Content with a Function:")
	pattern7 := regexp.MustCompile(
		"A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description3 := "Kayak. A boat for one person."
	replacedOneMore := pattern7.ReplaceAllStringFunc(description3, func(s string) string {
		return "This is the replacement content"
	})
	fmt.Println(replacedOneMore)
}
