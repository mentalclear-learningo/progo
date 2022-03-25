package main

import (
	"fmt"
	"strconv"
)

func main() {
	parseBin()
	fmt.Println("Int convinience example:")
	intConvinience()
	fmt.Println("Atoi example:")
	atoiExample()
}

func parseBin() {
	// val1 := "100"
	// int1, int1err := strconv.ParseInt(val1, 2, 8)  // 4
	val1 := "0b1100100"
	int1, int1err := strconv.ParseInt(val1, 0, 8)
	if int1err == nil {
		smallInt := int8(int1)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}

/* Base prefixes for Numeric Strings
0b - Binary
0o - Octal
0x - Hex
*/

func intConvinience() {
	val1 := "100"
	int1, int1err := strconv.ParseInt(val1, 10, 0)
	fmt.Printf("%T\n", int1)
	if int1err == nil {
		var intResult int = int(int1)
		fmt.Println("Parsed value:", intResult)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}

func atoiExample() {
	val1 := "100"
	int1, int1err := strconv.Atoi(val1)
	fmt.Printf("%T\n", int1)
	if int1err == nil {
		var intResult int = int1
		fmt.Println("Parsed value:", intResult)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}
