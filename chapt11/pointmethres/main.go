package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

// func (p Product) getName() string {
// 	return p.name
// }
// func (p Product) getCost(_ bool) float64 {
// 	return p.price
// }

// Change #2
func (p *Product) getName() string {
	return p.name
}
func (p *Product) getCost(_ bool) float64 {
	return p.price
}

func main() {
	product := Product{"Kayak", "Watersports", 275}
	fmt.Printf("product type: %T\n", product)
	//var expense Expense = product // Copied here
	// var expense Expense = &product // Reference to the same, Change #2

	// Change #3
	// var expense Expense = product
	/* Results:
	Product does not implement Expense (getCost method has pointer receiver)

	This means that now method recievers are enforcing the way the product can be assigned
	to the a var of Expense type
	*/

	// Added this to be able to complie:
	var expense Expense = &product

	fmt.Printf("expense type: %T\n", &expense)
	product.price = 100
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(false))

	// Comparing Interface Values
	// fmt.Println("\nComparing Interface Values:")
	// var e1 Expense = &Product{name: "Kayak"} // Dynamic types are pointers so not equal
	// var e2 Expense = &Product{name: "Kayak"}
	// fmt.Printf("e1 %T, e2 %T\n", &e1, &e2)
	// var e3 Expense = Service{description: "Boat Cover"} // These are regular structs
	// var e4 Expense = Service{description: "Boat Cover"} // with the same fields
	// fmt.Printf("e3 %T, e4 %T\n", &e3, &e4)
	// fmt.Println("e1 == e2", e1 == e2)
	// fmt.Println("e3 == e4", e3 == e4)

	// Performing Type Assertions - !see chagnes to the Service struct!
	fmt.Println("\nPerforming Type Assertions")
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		// Change, produces an error: interface conversion: main.Expense is *main.Product, not main.Service
		&Product{"Kayak", "Watersports", 275},
	}
	for _, expense := range expenses {
		//s := expense.(Service) // type assertion is performed by applying a period after a value
		// It is used to access the dynamic Service value from a slices of Expenses interface types
		// Now all the fields and methods are directly accesible:
		// fmt.Println("Service:", s.description, "Price:",
		//	s.monthlyFee*float64(s.durationMonths))

		// If assertion is Ok then access directly, if not - use methods
		// if s, ok := expense.(Service); ok {
		// 	fmt.Println("Service:", s.description, "Price:",
		// 		s.monthlyFee*float64(s.durationMonths))
		// } else {
		// 	fmt.Println("Expense:", expense.getName(),
		// 		"Cost:", expense.getCost(true))
		// }

		// Switching on Dynamic Types:
		switch value := expense.(type) { // special type assertion, with type keyword
		case Service:
			fmt.Printf("value's type: %T\n", value)
			fmt.Println("Service:", value.description, "Price:",
				value.monthlyFee*float64(value.durationMonths))
		case *Product:
			fmt.Printf("value's type: %T\n", value)
			fmt.Println("Product:", value.name, "Price:", value.price)
		default:
			fmt.Printf("value's type: %T\n", value)
			fmt.Println("Expense:", expense.getName(),
				"Cost:", expense.getCost(true))
		}

	}

}
