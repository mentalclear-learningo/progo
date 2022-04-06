package main

import "fmt"

func getProductName(index int) (name string, err error) {
	if len(Products) > index {
		name = fmt.Sprintf("Name of product: %v", Products[index].Name)
	} else {
		err = fmt.Errorf("Error for index %v", index)
	}
	return
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func formattingStrings() {
	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)

	fmt.Print("Product:", Kayak.Name, "Price:", Kayak.Price, "\n")

	// Formatting strings
	fmt.Printf("Product: %v, Price: $%4.2f", Kayak.Name, Kayak.Price)

	fmt.Println("Sprintf: ")
	name, _ := getProductName(1)
	fmt.Println(name)
	_, err := getProductName(10)
	fmt.Println(err.Error())

	fmt.Println("\nUsing the General-Purpose Formatting Verbs:")
	Printfln("Value: %v", Kayak)
	Printfln("Go syntax: %#v", Kayak)
	Printfln("Type: %T", Kayak)

	fmt.Println("\nControlling Struct Formatting:")

	// Value with fields: {Name:Lifejacket Category:Watersports Price:49.95}
	// before adding the func to Products
	Printfln("Value with fields: %+v", Products[1])

	fmt.Println("\nUsing the Integer Formatting Verbs:")

	number := 250
	Printfln("Binary: %b", number)
	Printfln("Decimal: %d", number)
	Printfln("Octal: %o, %O", number, number)
	Printfln("Hexadecimal: %x, %X", number, number)

	fmt.Println("\nUsing the Floating-Point Formatting Verbs:")
	number1 := 279.00
	Printfln("Decimalless with exponent: %b", number1)
	Printfln("Decimal with exponent: %e", number1)
	Printfln("Decimal without exponent: %f", number1)
	Printfln("Hexadecimal: %x, %X", number1, number1)
	Printfln("Decimal without exponent: >>%8.2f<<", number1)
	Printfln("Decimal without exponent: >>%.2f<<", number1)
	Printfln("Sign: >>%+.2f<<", number1)
	Printfln("Zeros for Padding: >>%010.2f<<", number1)
	Printfln("Right Padding: >>%-8.2f<<", number1)

	fmt.Println("\nUsing the String and Character Formatting Verbs:")
	name1 := "Kayak"
	Printfln("String: %s", name1)
	Printfln("Character: %c", []rune(name1)[0])
	Printfln("Unicode: %U", []rune(name1)[0])

	fmt.Println("\nUsing the Boolean Formatting Verb:")
	Printfln("Bool: %t", len(name1) > 1)
	Printfln("Bool: %t", len(name1) > 100)

	fmt.Println("\nUsing the Pointer Formatting Verb:")
	Printfln("Pointer: %p", &name1)
}

func scanningStrings() {
	var name string
	var category string
	var price float64
	fmt.Print("Enter text to scan: ")
	//n, err := fmt.Scanln(&name, &category, &price)
	source := "Lifejacket Watersports 48.95"
	n, err := fmt.Sscan(source, &name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func scanningTemplate() {
	var name string
	var category string
	var price float64
	source := "Product Lifejacket Watersports 48.95"
	template := "Product %s %s %f"
	n, err := fmt.Sscanf(source, template, &name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func main() {
	formattingStrings()

	fmt.Println("\nScanning strings")
	scanningStrings()
	fmt.Println("\nUsing scanning template:")
	scanningTemplate()
}
