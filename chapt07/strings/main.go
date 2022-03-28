package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Understanding the Dual Nature of Strings:")
	dualStringNature()
	fmt.Println("Euro sgin pitfall:")
	euroSignPitfall()
	fmt.Println("Converting a String to Runes:")
	convertStringsToRunes()
	fmt.Println("\nEnumerating Strings:")
	enumerStrings()
}

func dualStringNature() {
	var price string = "$48.95"
	// var currency byte = price[0]
	var currency string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Currency byte:", currency, "Actual string:", string(currency))
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

func euroSignPitfall() {
	var price string = "€48.95"
	var currency string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Length:", len(price))
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

func convertStringsToRunes() {
	var price []rune = []rune("€48.95")
	var currency string = string(price[0])
	var amountString string = string(price[1:]) // Slicing a rune creates slice of runes
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Length:", len(price))
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

func enumerStrings() {
	var price = "€48.95"
	for index, char := range price {
		fmt.Println(index, char, string(char))
	}
	fmt.Println("After explicit byte conversion:")
	for index, singleByte := range []byte(price) {
		fmt.Println(index, singleByte)
	}
}
