package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}
type ProductList []Product

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

// Added this one
func getProducts() []Product {
	return []Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
}

func main() {
	// Moved to the function above:
	// products := ProductList{
	// 	{"Kayak", "Watersports", 275},
	// 	{"Lifejacket", "Watersports", 48.95},
	// 	{"Soccer Ball", "Soccer", 19.50},
	// }
	products := ProductList(getProducts())
	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total:", total)
	}
}
