package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "It was a boat. A small boat."
	replace := strings.Replace(text, "boat", "canoe", 1)
	replaceAll := strings.ReplaceAll(text, "boat", "truck")
	fmt.Println("Replace:", replace)
	fmt.Println("Replace All:", replaceAll)

	fmt.Println("\nAltering strings with map func:")
	mapper := func(r rune) rune {
		if r == 'b' {
			return 'c'
		}
		return r
	}
	mapped := strings.Map(mapper, text)
	fmt.Println("Mapped:", mapped)

	fmt.Println("\nString replacer:")
	replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
	replaced := replacer.Replace(text)
	fmt.Println("Replaced:", replaced)
}
