package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func comparingStrings() {
	product := "Kayak"
	fmt.Println("Contains:", strings.Contains(product, "yak"))
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))

	price := "€100"
	fmt.Println("Strings Prefix:", strings.HasPrefix(price, "€"))
	fmt.Println("Bytes Prefix:", bytes.HasPrefix([]byte(price),
		[]byte{226, 130}))

}

func convertingStringCase() {
	description := "A boat for sailing"
	fmt.Println("Original:", description)
	fmt.Println("Title:", strings.Title(description))

	specialChar := "\u01c9"
	fmt.Println("Original:", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper:", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title:", titleChar, []byte(titleChar))
}

func workingWithCharacterCase() {
	product := "Kayak"
	for _, char := range product {
		fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
	}
}

func inspectingString() {
	description := "A boat for one person"
	fmt.Println("Count:", strings.Count(description, "o"))
	fmt.Println("Index:", strings.Index(description, "o"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("IndexAny:", strings.IndexAny(description, "abcd"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "abcd"))
}

func inspectingWithCustom() {
	description := "A boat for one person"
	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))
}

func main() {
	comparingStrings()
	fmt.Println("\nConverting String Case")
	convertingStringCase()
	fmt.Println("\nWorking with Character Case")
	workingWithCharacterCase()
	fmt.Println("\nInspecting Strings")
	inspectingString()
	fmt.Println("\nInspecting Strings with Custom Functions:")
	inspectingWithCustom()
}
