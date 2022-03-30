package main

import "fmt"

type Product struct {
	name, category string
	price          float64
	*Supplier
}
type Supplier struct {
	name, city string
}

func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name, category, price - 10, supplier}
}
func main() {
	acme := &Supplier{"Acme Co", "New York"}
	rand := &Supplier{"Rand Co", "Richmond"}
	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275, acme),
		newProduct("Hat", "Skiing", 42.50, rand),
	}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Supplier:",
			p.Supplier.name, p.Supplier.city)
	}
}
