package main

import (
	"fmt"
	"strconv"
)

func main() {
	exampleParsInt()
	exampleParsIntFail()
	fmt.Println("Convert example:")
	exampleParsIntConvert()
}

func exampleParsInt() {
	val1 := "100"
	int1, int1err := strconv.ParseInt(val1, 0, 8)
	if int1err == nil {
		fmt.Println("Parsed value:", int1)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}

func exampleParsIntFail() {
	val1 := "500"
	int1, int1err := strconv.ParseInt(val1, 0, 8)
	if int1err == nil {
		fmt.Println("Parsed value:", int1)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}

func exampleParsIntConvert() {
	val1 := "100"
	int1, int1err := strconv.ParseInt(val1, 0, 8)
	if int1err == nil {
		smallInt := int8(int1)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}
