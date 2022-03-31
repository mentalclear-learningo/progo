package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func (p Product) getName() string {
	return p.name
}

func (p Product) getCost(_ bool) float64 {
	return p.price
}

// Usuing interface in a function
func clacTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

// Using interface for struct fileds
type Account struct {
	accountNumber int
	expenses      []Expense
}

func main() {

	// Using an interface
	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
		Taxes{"Property Tax", 61.99},
	}

	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}

	fmt.Println("Total:", clacTotal(expenses))

	account := Account{
		accountNumber: 12345,
		expenses: []Expense{
			Product{"Kayak", "Watersports", 275},
			Service{"Boat Cover", 12, 89.50},
			Taxes{"Property Tax", 61.99},
		},
	}
	fmt.Println("\nSpecifically related to account:")
	for _, accExpense := range account.expenses {
		fmt.Println("Expense:", accExpense.getName(), "Cost:", accExpense.getCost(true))
	}
	fmt.Println("Total for the account:", clacTotal(account.expenses))

	// kayak := Product{"Kayak", "Watersports", 275}
	// insurance := Service{"Boat Cover", 12, 89.50}
	// fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	// fmt.Println("Service:", insurance.description, "Price:",
	// 	insurance.monthlyFee*float64(insurance.durationMonths))
}
