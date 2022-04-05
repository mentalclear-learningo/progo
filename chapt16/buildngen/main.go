package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "It was a boat. A small boat."
	elements := strings.Fields(text)
	joined := strings.Join(elements, "--")
	fmt.Println("Joined:", joined)

	fmt.Println("\nBuilding strings:")
	var builder strings.Builder
	for _, sub := range strings.Fields(text) {
		fmt.Println("This is sub:", sub)
		if sub == "small" {
			builder.WriteString("very ")
		}
		builder.WriteString(sub)
		builder.WriteRune(' ')
	}
	fmt.Println("String:", builder.String())
}
