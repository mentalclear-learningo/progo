package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "true"
	val2 := "false"
	val3 := "not true"
	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)
	fmt.Println("Bool 1", bool1, b1err)
	fmt.Println("Bool 2", bool2, b2err)
	fmt.Println("Bool 3", bool3, b3err)

	// Listing 5-18 error checking
	fmt.Println("Error checking:")
	errorChecking()

	// Listing 5-19
	fmt.Println("Error checking with if's init:")
	errorCheckingIfInit()
}

func errorChecking() {
	val1 := "0"
	bool1, b1err := strconv.ParseBool(val1)
	if b1err == nil {
		fmt.Println("Parsed value:", bool1)
	} else {
		fmt.Println("Cannot parse", val1)
	}
}

func errorCheckingIfInit() {
	val1 := "0"
	if bool1, b1err := strconv.ParseBool(val1); b1err == nil {
		fmt.Println("Parsed value:", bool1)
	} else {
		fmt.Println("Cannot parse", val1)
	}
}
