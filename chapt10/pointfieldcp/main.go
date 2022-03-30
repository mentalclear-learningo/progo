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

// Manual deep copy function:
func copyProduct(product *Product) Product {
	p := *product
	s := *product.Supplier
	p.Supplier = &s
	return p
}

func main() {
	acme := &Supplier{"Acme Co", "New York"}
	p1 := newProduct("Kayak", "Watersports", 275, acme)
	// This does a shallow copy for supplier, the pointer isn't followed.
	// p2 := *p1

	// Using manual deep copy fucntion
	p2 := copyProduct(p1)

	fmt.Println("Current supplier:", p1.Supplier.name)
	p1.name = "Original Kayak"
	p1.Supplier.name = "BoatCo"
	for _, p := range []Product{*p1, p2} {
		fmt.Println("Name:", p.name, "Supplier:",
			p.Supplier.name, p.Supplier.city)
	}
}
