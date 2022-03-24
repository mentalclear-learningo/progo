package main

import (
	"fmt"
	"sort"
)

func main() {
	noPointer()
	fmt.Println("With pointer:")
	withPointer()
}

func noPointer() {
	names := [3]string{"Alice", "Charlie", "Bob"}
	secondName := names[1]
	fmt.Println(secondName)
	fmt.Println(names)
	sort.Strings(names[:])
	fmt.Println(secondName)
	fmt.Println(names)
}

func withPointer() {
	names := [3]string{"Alice", "Charlie", "Bob"}
	secondName := &names[1]
	fmt.Println(*secondName)
	fmt.Println(names)
	sort.Strings(names[:])
	fmt.Println(*secondName)
	fmt.Println(names)
}
