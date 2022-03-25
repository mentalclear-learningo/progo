package main

import (
	"fmt"
	"strconv"
)

func main() {
	expFromatBools()
	fmt.Println("Example Ints Formatting:")
	exmplIntValues()
	fmt.Println("Example Ints Conv:")
	exmplIntConv()
	fmt.Println("Example Float Point:")
	exmplFloatPoint()
}

func expFromatBools() {
	val1 := true
	val2 := false
	str1 := strconv.FormatBool(val1)
	str2 := strconv.FormatBool(val2)
	fmt.Println("Formatted value 1: " + str1)
	fmt.Println("Formatted value 2: " + str2)
}

func exmplIntValues() {
	val := 275
	base10String := strconv.FormatInt(int64(val), 10)
	base2String := strconv.FormatInt(int64(val), 2)
	fmt.Println("Base 10: " + base10String)
	fmt.Println("Base 2: " + base2String)
}

func exmplIntConv() {
	val := 275
	base10String := strconv.Itoa(val)
	base2String := strconv.FormatInt(int64(val), 2)
	fmt.Println("Base 10: " + base10String)
	fmt.Println("Base 2: " + base2String)
}

func exmplFloatPoint() {
	val := 49.95
	Fstring := strconv.FormatFloat(val, 'f', 2, 64)
	Estring := strconv.FormatFloat(val, 'e', -1, 64)
	fmt.Println("Format F: " + Fstring)
	fmt.Println("Format E: " + Estring)
}
