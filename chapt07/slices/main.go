package main

import "fmt"

func main() {
	fmt.Println("\nArray: ")
	arrayExample()
	fmt.Println("\nSlice: ")
	sliceExample()
	fmt.Println("\nSlice literal: ")
	sliceLiteral()
	fmt.Println("\nAppending Elements to a Slice:")
	appendToSlice()
	fmt.Println("\nBoth items:")
	bothItems()
	fmt.Println("\nAllocating Additional Slice Capacity:")
	additionalSliceCapacity()
	fmt.Println("\nArray istn't replaced:")
	singleArrSlice()
	fmt.Println("\nAppending One Slice to Another:")
	appendSliceToAnother()
	fmt.Println("\nCreating Slices from Existing Arrays:")
	createSliceFromArray()
	fmt.Println("\nAppending Elements When Using Existing Arrays for Slices:")
	appendingElementsUsingExisting()
	fmt.Println("\nAppend another to slice:")
	appendAnotherToSlice()
	fmt.Println("\nAppending Another Element in the main.go File in the collections Folder:")
	appendAnotherToSliceOverCapacity()
	fmt.Println("\nSpecifying Capacity When Creating a Slice from an Array:")
	specifyCapacity()
	fmt.Println("\nCreating Slices from Other Slices:")
	createSliceFromSlice()
}

func arrayExample() {
	var names [3]string
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)
	fmt.Printf("Type - %T\n", names)
}

func sliceExample() {
	names := make([]string, 3)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)
	fmt.Printf("Type - %T\n", names)
}

func sliceLiteral() {
	names := []string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names)
}

func appendToSlice() {
	names := []string{"Kayak", "Lifejacket", "Paddle"}
	names = append(names, "Hat", "Gloves")
	fmt.Println(names)
}

func bothItems() {
	names := []string{"Kayak", "Lifejacket", "Paddle"}
	appendedNames := append(names, "Hat", "Gloves")
	names[0] = "Canoe"
	fmt.Println("names:", names)
	fmt.Println("appendedNames:", appendedNames)
}

func additionalSliceCapacity() {
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))
	for index, item := range names {
		fmt.Println("Index:", index, "Item:", item)
	}
}

// The underlying array isn't replaced when the append function
// is called on a slice with enough capacity to accommodate
// the new elements
func singleArrSlice() {
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	// When there's enough room the same backing array will be used
	appendedNames := append(names, "Hat", "Gloves")

	// When there's NOT enough room - the new array will be created:
	// appendedNames := append(names, "Hat", "Gloves", "Water Bottle", "First Aid Kit")

	names[0] = "Canoe"
	fmt.Println("names:", names)
	fmt.Println("appendedNames:", appendedNames)
}

func appendSliceToAnother() {
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	moreNames := []string{"Hat", "Gloves"}
	appendedNames := append(names, moreNames...)
	fmt.Println("appendedNames:", appendedNames)
}

func createSliceFromArray() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	fmt.Printf("%T\n", products)
	someNames := products[1:3]
	allNames := products[:]
	fmt.Println("someNames:", someNames)
	fmt.Printf("%T\n", someNames)
	fmt.Println("allNames", allNames)
	fmt.Printf("%T\n", allNames)

	var allItems [3][]string
	allItems[0] = products[:]
	allItems[1] = someNames[:]
	allItems[2] = allNames[:]
	fmt.Println("All in:", allItems)

	fmt.Println("Changes to the slices - changes single backing array:")
	someNames[0] = "Canoe-1"
	fmt.Println("Results:", someNames, "and", allNames)
}

func appendingElementsUsingExisting() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
}

func appendAnotherToSlice() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	someNames = append(someNames, "Gloves")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
}

func appendAnotherToSliceOverCapacity() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	someNames = append(someNames, "Gloves")
	someNames = append(someNames, "Boots")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))

	someNames[0] = "Canoe-1"
	fmt.Println("Results:", someNames, "and", allNames)
}

func specifyCapacity() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3:3]
	allNames := products[:]
	someNames = append(someNames, "Gloves")
	//someNames = append(someNames, "Boots")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
}

func createSliceFromSlice() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]
	fmt.Println("1. allNames now:", allNames)
	someNames := allNames[1:3]
	allNames = append(allNames, "Gloves")
	fmt.Println("2. allNames now:", allNames)
	allNames[1] = "Canoe"
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}
