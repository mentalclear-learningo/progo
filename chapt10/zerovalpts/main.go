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

func main() {
	// var prod Product
	var prodPtr *Product
	var prod Product = Product{Supplier: &Supplier{}}
	// Attempt to access the name field of the embedded struct, this is resolved
	// by init the struct pointer fields see above line
	fmt.Println("Value:", prod.name, prod.category, prod.price, prod.Supplier.name)
	fmt.Println("Pointer:", prodPtr)
}
