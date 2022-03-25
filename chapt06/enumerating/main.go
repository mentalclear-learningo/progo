package main

import "fmt"

func main() {
	enumeratingSeqs()
	fmt.Println("My experiment: ")
	myExerim()
	fmt.Println("Getting only Index")
	indexesOnly()
	fmt.Println("Enumerating Built-in Data Structures:")
	enumBuiltIn()
}

func enumeratingSeqs() {
	product := "Kayak"
	for index, character := range product {
		fmt.Println("Index:", index, "Character:", string(character))
	}

}

func myExerim() {
	product := "Kayak"
	for _, character := range product {
		fmt.Println("Character:", string(character))
	}
}

func indexesOnly() {
	product := "Kayak"
	for index := range product {
		fmt.Println("Index:", index)
	}
}

func enumBuiltIn() {
	products := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	for index, element := range products {
		fmt.Println("Index:", index, "Element:", element)
	}
}
