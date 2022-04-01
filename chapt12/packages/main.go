package main

import (
	"fmt"
	// currencyFmt "packages/fmt"
	// Dot import
	. "packages/fmt" // Use with caution
	"packages/store"

	// Importing nesting package
	"packages/store/cart"

	// Init effect from package data
	_ "packages/data"
	// External package
	"github.com/fatih/color"
)

func main() {
	// product := store.Product{
	// 	Name:     "Kayak",
	// 	Category: "Watersports",
	// 	price:    279,
	// }
	product := store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", product.GetPrice())

	// Aliased import
	// fmt.Println("Formatted Price:", currencyFmt.ToCurrency(product.GetPrice()))

	// Dotted import
	fmt.Println("Formatted Price:", ToCurrency(product.GetPrice()))

	// Nested Packages code:
	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}
	fmt.Println("Name:", cart.CustomerName)
	fmt.Println("Total:", ToCurrency(cart.GetTotal()))

	// Use extarnal package
	color.Green("Name: " + cart.CustomerName)
	color.Cyan("Total: " + ToCurrency(cart.GetTotal()))
}
