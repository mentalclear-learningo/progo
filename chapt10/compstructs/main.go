package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	type Product struct {
		name, category string
		price          float64
		// Structs cannot be compared when they have incomparable types:
		//otherNames []string
	}
	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}
	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("p1 == p3:", p1 == p3)

	fmt.Println("\nConverting between struct types:")
	type Item struct {
		name     string
		category string
		price    float64
	}

	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item := Item{name: "Kayak", category: "Watersports", price: 275.00}

	fmt.Println("prod == item:", prod == Product(item))

	fmt.Println("\nDefining Anonymous Struct Types")
	newItem := Item{name: "Stadium", category: "Soccer", price: 75000}
	writeName(prod)
	writeName(newItem)

	fmt.Println("\nAssigining value to anonymous struct:")
	var builder strings.Builder
	// Here's an example of defining and assigning a value to an anonimous struct
	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod.name,
		ProductPrice: prod.price,
	})
	fmt.Println(builder.String())
}

// Defining Anonymous Struct Types
func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name:", val.name)
}
