package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}
func selectCalculator(price float64) calcFunc {
	if price > 100 {
		return func(price float64) float64 { // returning anonymous function here
			return price + (price * 0.2)
		}
	}
	return func(price float64) float64 { // and here too
		return price
	}
}
func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}

	// example of anonymous func used as argument
	for product, price := range products {
		printPrice(product, price, func(price float64) float64 {
			return price + (price * 0.2)
		})
	}
}
