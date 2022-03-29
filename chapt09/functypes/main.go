package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price * 1.2
}

func calcWithoutTax(price float64) float64 {
	return price
}

func functCompareZero(products map[string]float64) {
	for product, price := range products {
		var calcFunc func(float64) float64
		// Understanding Function Comparisons and the Zero Type
		fmt.Println("Function assigned:", calcFunc == nil)
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		fmt.Println("Function assigned:", calcFunc == nil)
		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}
}

// Using Functions as Arguments
func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func functionCompareZero(products map[string]float64) {
	for product, price := range products {
		if price > 100 {
			printPrice(product, price, calcWithTax)
		} else {
			printPrice(product, price, calcWithoutTax)
		}
	}
}

// Using Functions as Results
func selectCalculator(price float64) func(float64) float64 { // returning function
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
}

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	functCompareZero(products)
	fmt.Println("\nUsing Functions as Arguments:")
	functionCompareZero(products)

	fmt.Println("\nUsing Functions as Results:")
	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}
