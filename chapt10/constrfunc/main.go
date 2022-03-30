package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

// Constructor function:
func newProduct(name, category string, price float64) *Product {
	// Changed to apply discount of 10
	return &Product{name, category, price - 10}
}
func main() {
	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Hat", "Skiing", 42.50),
	}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price", p.price)
	}
}
