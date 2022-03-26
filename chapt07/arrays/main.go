package main

import "fmt"

func main() {
	arrayExample()
	fmt.Println("Arrays literal syntax:")
	arraysLiteral()
	fmt.Println("Arrays types:")
	arrayTypes()
	fmt.Println("Array values:")
	arrayValues()
	fmt.Println("\nComparing arrays:")
	comparingArrays()
	fmt.Println("\nEnumerating arrays:")
	enumeratingArrays()
}

func arrayExample() {
	var names [3]string
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)

	for _, elem := range names {
		fmt.Println("Element:", elem)
	}
}

func arraysLiteral() {
	names := [4]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names)
	for _, elem := range names {
		fmt.Println("Element:", elem)
	}
}

func arrayTypes() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	// var otherArray [4]string = names // Creates compilation error
	fmt.Println(names)
}

func arrayValues() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	otherArray := names
	// Pointer can be used here too:
	otherArray2 := &names
	names[0] = "Canoe"
	fmt.Println("names:", names)
	fmt.Println("otherArray: ", otherArray)
	fmt.Println("otherArray2:", *otherArray2)
}

// Arrays are equal if they are the same type and contain equal
// elements in the same order.
func comparingArrays() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	moreNames := [3]string{"Kayak", "Lifejacket", "Paddle"}
	same := names == moreNames
	fmt.Println("comparison:", same)
}

func enumeratingArrays() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}

	for _, value := range names {
		fmt.Println("Value:", value)
	}
}
