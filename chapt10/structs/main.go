package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}

	type StockLevel struct {
		Product           // Embedded Field
		Alternate Product // If more embedded fields are needed, add name
		count     int
	}

	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
		// Partially assign values - commenting the below line:
		//price:    275,
	}
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)

	// Zero types assigned:
	var lifejacket Product
	fmt.Println("Name is zero value:", lifejacket.name == "")
	fmt.Println("Category is zero value:", lifejacket.category == "")
	fmt.Println("Price is zero value:", lifejacket.price == 0)

	// Using field positions to create a new value:

	var motorboat = Product{"Motorboat", "Watersports", 1859.95}
	fmt.Println("Name:", motorboat.name, "\nCategory:", motorboat.category, "\nPrice:", motorboat.price)

	fmt.Println("\nDefining Embedded Fields")
	stockItem := StockLevel{
		Product: Product{"Inflatable Boat", "Watersports", 255.00},
		// Example for more than one embedded field:
		Alternate: Product{"Paddle", "Watersports", 25.99},
		count:     100,
	}
	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Count:", stockItem.count)
	fmt.Println("Alt Name:", stockItem.Alternate.name)
}
