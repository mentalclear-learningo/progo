package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

// func calcTax(product *Product) {
// 	if (*product).price > 100 {
// 		(*product).price += (*product).price * 0.2
// 	}
// }

// Simplified:
// func calcTax(product *Product) {
// 	if product.price > 100 {
// 		product.price += product.price * 0.2
// 	}
// }

// Another level:
func calcTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

func main() {

	p1 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	// p2 := p1 // creates a copy
	p2 := &p1
	p1.name = "Original Kayak"
	fmt.Println("P1:", p1.name)
	// fmt.Println("P2:", p2.name)
	fmt.Println("P2:", (*p2).name)

	fmt.Println("\nStruct pointer convinience syntax")
	// kayak := Product{
	// 	name:     "Kayak",
	// 	category: "Watersports",
	// 	price:    275,
	// }
	// calcTax(&kayak)

	// Creating a poiner directly:
	// kayak := &Product{
	// 	name:     "Kayak",
	// 	category: "Watersports",
	// 	price:    275,
	// }
	// calcTax(kayak)

	// Another level
	kayak := calcTax(&Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	})

	fmt.Println("Name:", kayak.name, "Category:",
		kayak.category, "Price:", kayak.price)
}
