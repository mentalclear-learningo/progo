package main

import "fmt"

//func printPrice(product string, price float64, taxRate float64) {
func printPrice(product string, price, taxRate float64) { // Ommit parameter data type
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

// Ommiting parameter name
func printPriceOmmitName(product string, price, _ float64) {
	taxAmount := price * 0.25
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

// Ommiting all parameter names
func printPriceOmmitAll(string, float64, float64) {
	//taxAmount := price * 0.25
	fmt.Println("No Parameters")
}

func printSuppliers(product string, suppliers []string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

// Variadic parameter must be the last parameter defined
func printSuppliersVariadic(product string, suppliers ...string) {
	// for _, supplier := range suppliers {
	// 	fmt.Println("Product:", product, "Supplier:", supplier)
	// }

	// Guarding against no value provided for suppliers
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

// func swapValues(first, second int) {
// 	fmt.Println("Before swap:", first, second)
// 	temp := first
// 	first = second
// 	second = temp
// 	fmt.Println("After swap:", first, second)
// }

func swapValues(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)
}

func calcTax(price float64) float64 {
	//return price * 1.2
	return price + (price * 0.2)
}

func swapValuesMulti(first, second int) (int, int) {
	return second, first
}

// Giving multiple meanings to single result
// func calcTax2(price float64) float64 {
// 	if price > 100 {
// 		return price * 0.2
// 	}
// 	return -1
// }
func calcTax2(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

func calcTotalPrice(products map[string]float64,
	minSpend float64) (total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcTax2(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	return
}

func calcTotalPriceDiscard(products map[string]float64) (count int, total float64) {
	count = len(products)
	for _, price := range products {
		total += price
	}
	return
}

func calcTotalPriceDefer(products map[string]float64) (count int, total float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
}

func main() {
	printPrice("Kayak", 275, 0.2)
	printPrice("Lifejacket", 48.95, 0.2)
	printPrice("Soccer Ball", 19.50, 0.15)

	fmt.Println("\nOmmiting Parameter Name:")
	printPriceOmmitName("Kayak", 275, 0.2)
	printPriceOmmitName("Lifejacket", 48.95, 0.2)
	printPriceOmmitName("Soccer Ball", 19.50, 0.15)

	fmt.Println("\nOmmiting All Parameter Names:")
	printPriceOmmitAll("Kayak", 275, 0.2)
	printPriceOmmitAll("Lifejacket", 48.95, 0.2)
	printPriceOmmitAll("Soccer Ball", 19.50, 0.15)

	fmt.Println("\nVariadic Parameters aren't used:")
	printSuppliers("Kayak", []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
	printSuppliers("Lifejacket", []string{"Sail Safe Co"})

	fmt.Println("\nDefining Variadic Parameters:")
	printSuppliersVariadic("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliersVariadic("Lifejacket", "Sail Safe Co")

	fmt.Println("\nNo Arguments for a Variadic Parameter:")
	printSuppliersVariadic("Soccer Ball") // Will print nothing since the slice is nil

	fmt.Println("\nUsing slices with variadic parameters:")
	names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliersVariadic("Kayak", names...)

	fmt.Println("\nUsing Pointers as Function Parameters:")
	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	swapValues(&val1, &val2)
	fmt.Println("After calling function", val1, val2)

	fmt.Println("\nDefining and Using Function Results:")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		// priceWithTax := calcTax(price)
		// fmt.Println("Product: ", product, "Price:", priceWithTax)

		// This works as well
		fmt.Println("Product: ", product, "Price:", calcTax(price))
	}

	fmt.Println("\nReturning Multiple Function Results:")
	val3, val4 := 10, 20
	fmt.Println("Before calling function", val3, val4)
	val3, val4 = swapValuesMulti(val3, val4)
	fmt.Println("After calling function", val3, val4)

	fmt.Println("\nUsing Multiple Results Instead of Multiple Meanings:")
	products2 := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products2 {
		// Giving multiple meanings to single result
		// tax := calcTax2(price)
		// if tax != -1 {
		// 	fmt.Println("Product: ", product, "Tax:", tax)
		// } else {
		// 	fmt.Println("Product: ", product, "No tax due")
		// }

		// Work with multiple outputs
		// taxAmount, taxDue := calcTax2(price)
		// if taxDue {
		// 	fmt.Println("Product: ", product, "Tax:", taxAmount)
		// } else {
		// 	fmt.Println("Product: ", product, "No tax due")
		// }

		// Another optimisation, initilization statement:
		if taxAmount, taxDue := calcTax2(price); taxDue {
			fmt.Println("Product: ", product, "Tax:", taxAmount)
		} else {
			fmt.Println("Product: ", product, "No tax due")
		}
	}

	fmt.Println("\nUsing Named Results:")
	total1, tax1 := calcTotalPrice(products, 10)
	fmt.Println("Total 1:", total1, "Tax 1:", tax1)
	total2, tax2 := calcTotalPrice(nil, 10)
	fmt.Println("Total 2:", total2, "Tax 2:", tax2)

	fmt.Println("\nUsing the Blank Identifier to Discard Results:")
	_, total := calcTotalPriceDiscard(products2)
	fmt.Println("Total:", total)

	fmt.Println("\nUsing the defer Keyword:")
	// The defer keyword is used to schedule a function call that will be
	// performed immediately before the current function returns
	products3 := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	_, total3 := calcTotalPriceDefer(products3)
	fmt.Println("Total:", total3)
}
