package main

import (
	"composition/store"
	"fmt"
)

func main() {
	// kayak := store.NewProduct("Kayak", "Watersports", 275)
	// lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}
	// for _, p := range []*store.Product{kayak, lifejacket} {
	// 	fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.GetPrice(0.2))
	// }

	// Adding Boat
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	/* In this assignment values can be red and assigned buy:
	boat.Name = "Green Kayak"

	but if I'd create value using literal synatax(not using constructor):
	boat := store.Boat {
		Name: "Kayak",
		Category: "Watersports",
		Capacity: 1,
		Motorized: false,
	}
	I'd get "unknow field" error from the compiler, so I would have to assign
	values like below:
	*/
	boat := store.Boat{
		Product: &store.Product{
			Name:     "Literal Kayak",
			Category: "Watersports"},
		Capacity:  1,
		Motorized: false,
	}

	for _, b := range boats {
		// b.Name can be used due to "field promotion"
		fmt.Println("Conventional:", b.Product.Name, "| Direct:", b.Name)
		// Same thing happens with methods
		fmt.Println("Boat:", b.Name, "| Price:", b.GetPrice(0.2))
	}

	fmt.Println("Boat assigned with literal:", boat.Name)

	// Changed after Multiple Nested Types in the Same Struct
	// fmt.Println("\nChain of Nested Types results:")
	// rentals := []*store.RentalBoat{
	// 	store.NewRentalBoat("Rubber Ring", 10, 1, false, false),
	// 	store.NewRentalBoat("Yacht", 50000, 5, true, true),
	// 	store.NewRentalBoat("Super Yacht", 100000, 15, true, true),
	// }
	// for _, r := range rentals {
	// 	fmt.Println("Rental Boat:", r.Name, "| Rental Price:", r.GetPrice(0.2))
	// }

	fmt.Println("\nUsing Multiple Nested Types in the Same Struct:")
	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true,
			"Dora", "Charlie"),
	}
	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.GetPrice(0.2),
			"Captain:", r.Captain)
	}

	fmt.Println("\nUnderstanding When Promotion Cannot Be Performed")
	product := store.NewProduct("Kayak", "Watersports", 279)
	deal := store.NewSpecialDeal("Weekend Special", product, 50)
	Name, price, Price := deal.GetDetails()
	fmt.Println("Name:", Name)
	fmt.Println("Price field:", price)
	fmt.Println("Price method:", Price)

	fmt.Println("\nUnderstanding Promotion Ambiguity:")
	kayak := store.NewProduct("Kayak", "Watersports", 279)
	type OfferBundle struct {
		*store.SpecialDeal
		*store.Product
	}
	bundle := OfferBundle{
		store.NewSpecialDeal("Weekend Special", kayak, 50),
		store.NewProduct("Lifrejacket", "Watersports", 48.95),
	}
	//fmt.Println("Price:", bundle.GetPrice(0)) // Go cannot decide which method to use!
	fmt.Println("One special:", bundle.SpecialDeal.Name) // bundle.Name is ambigous too!
}
