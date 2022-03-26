package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	fmt.Println("\nUsing the copy Function to Ensure Slice Array Separation:")
	usingCopy()
	fmt.Println("\nUninitialized Slice Pitfall:")
	uninitSlicePitfall()
	fmt.Println("\nSpecifying Ranges When Copying Slices:")
	specifyRanges()
	fmt.Println("\nCopying Slices with Different Sizes:")
	copySliceDifSize()
	fmt.Println("\nCopying a Larger Source Slice:")
	copySliceSmallerSize()
	fmt.Println("\nDeleting Slice Elements:")
	deletingSliceElements()
	fmt.Println("\nEnumerating Slices:")
	enumerSlices()
	fmt.Println("\nSorting Slices:")
	sortingSlices()
	fmt.Println("\nComparing Slices:")
	compareSlices()
	fmt.Println("\nGetting the Array Underlying a Slice:")
	arrayUnderlyungSlice()
}

func usingCopy() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]
	fmt.Println("allNames 1:", allNames)
	someNames := make([]string, 2)
	fmt.Println("someNames 1:", someNames)
	fmt.Printf("someNames length %d, capacity %d\n", len(someNames), cap(someNames))
	copy(someNames, allNames)
	fmt.Println("someNames 2:", someNames)
	fmt.Println("allNames 2:", allNames)
}

func uninitSlicePitfall() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]
	// Nothing is going to be copied because the slice hasn't been init.
	var someNames []string
	fmt.Printf("someNames length %d, capacity %d\n", len(someNames), cap(someNames))
	copy(someNames, allNames)
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}

func specifyRanges() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]
	someNames := []string{"Boots", "Canoe"}
	copy(someNames[1:], allNames[2:3])
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}

func copySliceDifSize() {
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacementProducts := []string{"Canoe", "Boots"}
	copy(products, replacementProducts)
	fmt.Println("products:", products)
}

func copySliceSmallerSize() {
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacementProducts := []string{"Canoe", "Boots"}
	copy(products[0:1], replacementProducts)
	fmt.Println("products:", products)
}

func deletingSliceElements() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	deleted := append(products[:2], products[3:]...)
	fmt.Println("Deleted:", deleted)
}

func enumerSlices() {
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	for index, value := range products[2:] {
		fmt.Println("Index:", index, "Value:", value)
	}
}

func sortingSlices() {
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	sort.Strings(products)
	for index, value := range products {
		fmt.Println("Index:", index, "Value:", value)
	}
}

func compareSlices() {
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	p2 := p1
	// Invalid operation: compared (slice can only be compared to nil)
	// fmt.Println("Equal:", p1 == p2)
	fmt.Println("Equal:", reflect.DeepEqual(p1, p2))
}

func arrayUnderlyungSlice() {
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	arrayPtr := (*[3]string)(p1)
	array := *arrayPtr
	fmt.Println(array)
}
