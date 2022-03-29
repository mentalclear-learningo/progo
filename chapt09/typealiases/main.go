package main

import "fmt"

// This all based on the previous examples, see functypes here in the folder

type calcFunc func(float64) float64

func calcWithTax(price float64) float64 {
	return price * 1.2
}
func calcWithoutTax(price float64) float64 {
	return price
}
func printPrice(product string, price float64, calculator calcFunc) { // type as func arg
	fmt.Println("Product:", product, "Price:", calculator(price))
}
func selectCalculator(price float64) calcFunc { // type used here, as function return
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
	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}
