package main

import (
	"fmt"
	"strconv"
)

func main() {
	baseListing64()
	fmt.Println("Using If statement")
	usingIfStatement()
	fmt.Println("Using Else keyword:")
	usingElse()
	fmt.Println("Understanding if Statement Scope:")
	ifStatementScope()
	fmt.Println("Initialization Statement with an if Statement:")
	initWithIfStatement()
}

func baseListing64() {
	kayakPrice := 275.00
	fmt.Println("Price", kayakPrice)
}

func usingIfStatement() {
	kayakPrice := 275.00
	if kayakPrice > 100 {
		fmt.Println("Price is greater than 100")
	}
}

func usingElse() {
	kayakPrice := 275.00
	if kayakPrice > 500 {
		fmt.Println("Price is greater than 500")
	} else if kayakPrice < 100 {
		fmt.Println("Price is less than 100")
	} else {
		fmt.Println("Price not matched by earlier expressions")
	}
	// else if kayakPrice > 200 && kayakPrice < 300 {
	// 	fmt.Println("Price is between 200 and 300")
	// }
}

// ifStatementScope - each scopedVar is in scope for the code block it's in.
func ifStatementScope() {
	kayakPrice := 275.00
	if kayakPrice > 500 {
		scopedVar := 500
		fmt.Println("Price is greater than", scopedVar)
	} else if kayakPrice < 100 {
		scopedVar := "Price is less than 100"
		fmt.Println(scopedVar)
	} else {
		scopedVar := false
		fmt.Println("Matched: ", scopedVar)
	}
}

func initWithIfStatement() {
	priceString := "275"
	if kayakPrice, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price:", kayakPrice)
	} else {
		fmt.Println("Error:", err)
	}

	/* That's a combination of following:
		kayakPrice, err := strconv.Atoi(priceString)
	 	if err == nil { blah... }
	*/
}
