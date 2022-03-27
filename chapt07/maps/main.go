package main

import (
	"fmt"
	"sort"
)

func main() {
	createMap()
	fmt.Println("\nUsing the Map Literal Syntax")
	mapLiteralSyntax()
	fmt.Println("\nChecking for Items in a Map")
	checkingItemsInMaps()
	fmt.Println("\nRemoving Items from a Map:")
	removingItemsFromMap()
	fmt.Println("\nEnumerating the Contents of a Map:")
	enumerMapContents()
	fmt.Println("\nEnumerating a Map in Order:")
	enumerMapInOrder()
}

func createMap() {
	products := make(map[string]float64, 10)
	products["Kayak"] = 279
	products["Lifejacket"] = 48.95
	fmt.Println("Map size:", len(products))
	fmt.Println("Price:", products["Kayak"])
	fmt.Println("Price:", products["Hat"])
}

func mapLiteralSyntax() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        11.32,
	}
	fmt.Println("Map size:", len(products))
	fmt.Println("Price:", products["Kayak"])
	fmt.Println("Price:", products["Hat"])
}

func checkingItemsInMaps() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	// value, ok := products["Hat"]
	// if ok {
	// 	fmt.Println("Stored value:", value)
	// } else {
	// 	fmt.Println("No stored value")
	// }
	if value, ok := products["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
}

func removingItemsFromMap() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	delete(products, "Hat")
	if value, ok := products["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
}

func enumerMapContents() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	for key, value := range products {
		fmt.Println("Key:", key, "| Value:", value)
	}
}

func enumerMapInOrder() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	keys := make([]string, 0, len(products))
	for key := range products {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", products[key])
	}
}
