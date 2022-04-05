package main

import (
	"fmt"
	"strings"
)

func main() {
	username := " Alice"
	trimmed := strings.TrimSpace(username)
	fmt.Println("Trimmed:", ">>"+trimmed+"<<")

	description := "A boat for one person"
	trimmed2 := strings.Trim(description, "Asno ")
	fmt.Println("Trimmed:", trimmed2)

	fmt.Println("\nTrimming substrings:")
	prefixTrimmed := strings.TrimPrefix(description, "A boat ")
	wrongPrefix := strings.TrimPrefix(description, "A hat ")
	fmt.Println("Trimmed:", prefixTrimmed)
	fmt.Println("Not trimmed:", wrongPrefix)

	fmt.Println("\nTrimming with custom funcs:")
	trimmer := func(r rune) bool {
		return r == 'A' || r == 'n'
	}
	trimmed3 := strings.TrimFunc(description, trimmer)
	fmt.Println("Trimmed:", trimmed3)
}
