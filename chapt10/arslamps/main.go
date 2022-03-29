package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
		//otherNames []string
	}
	type StockLevel struct {
		Product
		Alternate Product
		count     int
	}
	array := [1]StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Array:", array[0].Product.name)
	slice := []StockLevel{
		{
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Slice:", slice[0].Product.name)
	kvp := map[string]StockLevel{
		"kayak": {
			Product:   Product{"Kayak", "Watersports", 275.00},
			Alternate: Product{"Lifejacket", "Watersports", 48.95},
			count:     100,
		},
	}
	fmt.Println("Map:", kvp["kayak"].Product.name)
}
