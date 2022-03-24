package main

import (
	"fmt"
)

//"math/rand"

func main() {
	// consts
	// const price, tax float32 = 275, 27.50
	// const quantity, inStock = 2, true
	// fmt.Println("Total:", 2*quantity*(price+tax))
	// fmt.Println("In stock: ", inStock)

	// vars
	// var price float32 = 275.00
	// var tax float32 = 27.50
	// fmt.Println(price + tax)
	// price = 300
	// fmt.Println(price + tax)

	// mixing types - creates error
	// var price = 275.00
	// var tax float32 = 27.50
	// fmt.Println(price + tax)

	// default values
	// var price float32
	// fmt.Println(price)
	// price = 275.00
	// fmt.Println(price)

	// Multiplevariables in one line
	// var price, tax = 275.00, 27.50
	// fmt.Println(price + tax)

	// var price, tax float64
	// price = 275.00
	// tax = 27.50
	// fmt.Println(price + tax)

	// short declaration:
	// price := 275.00
	// fmt.Println(price)

	// Multiple variables:
	// price, tax, inStock := 275.00, 27.50, true
	// fmt.Println("Total:", price+tax)
	// fmt.Println("In stock:", inStock)

	// redefine, with problem
	// price, tax, inStock := 275.00, 27.50, true
	// fmt.Println("Total:", price+tax)
	// fmt.Println("In stock:", inStock)
	// var price2, tax = 200.00, 25.00
	// fmt.Println("Total 2:", price2+tax)

	// This works:
	// price, tax, inStock := 275.00, 27.50, true
	// fmt.Println("Total:", price+tax)
	// fmt.Println("In stock:", inStock)
	// price2, tax := 200.00, 25.00
	// fmt.Println("Total 2:", price2+tax)

	price, tax, inStock, _ := 275.00, 27.50, true, true
	var _ = "Alice"
	fmt.Println("Total:", price+tax)
	fmt.Println("In stock:", inStock)

	behaviorWOPointer()
	fmt.Println("*****")
	withPointer()
	fmt.Println("*****")
	followingPointer()
	fmt.Println("*****")
	secondUseOfPointer()
	fmt.Println("***** Pointer Zero Value:")
	pointerZeroValues()
	fmt.Println("***** Pointing at Pointers")
	pointingAtPointers()
}

func behaviorWOPointer() {
	first := 100
	second := first
	first++
	fmt.Println("First:", first)
	fmt.Println("Second:", second)
}

func withPointer() {
	first := 100
	var second *int = &first
	first++
	fmt.Println("First:", first)
	//fmt.Println("Second:", second)

	// Following the pointer, dereferencing
	fmt.Println("Second:", *second)
}

func followingPointer() {
	first := 100
	var second *int = &first
	first++
	*second++
	fmt.Println("First:", first)

	// Following the pointer, dereferencing
	fmt.Println("Second:", *second)
}

func secondUseOfPointer() {
	first := 100
	second := &first
	first++
	*second++
	var myNewPointer *int
	myNewPointer = second
	*myNewPointer++
	fmt.Println("First:", first)
	fmt.Println("Second:", *second)
}

func pointerZeroValues() {
	first := 100
	var second *int
	fmt.Println(second)
	//fmt.Println(*second) //Following nil pointer produces an error
	second = &first
	fmt.Println(second)
}

func pointingAtPointers() {
	first := 100
	second := &first
	third := &second
	fmt.Println(first)
	fmt.Println(*second)
	fmt.Println(**third)
}
