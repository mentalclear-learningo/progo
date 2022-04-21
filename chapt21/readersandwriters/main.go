package main

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	//"bufio"
	"fmt"
)

// func writeReplaced(writer io.Writer, str string, subs ...string) {
// 	replacer := strings.NewReplacer(subs...)
// 	replacer.WriteString(writer, str)
// }

func main() {
	encodingBasicTypes()
	fmt.Println("\nEncoding Arrays and Slices:")
	encodingArraysAndSlices()
	fmt.Println("\nEncoding Maps:")
	encodingMaps()
	fmt.Println("\nEncoding Structs:")
	encodingStructs()
	fmt.Println("\nUnderstanding the Effect of Promotion in JSON in Encoding:")
	promoEffect()
	fmt.Println("\nCustomizing the JSON Encoding of Structs:")
	customEncoding()
	fmt.Println("\nOmmiting")
	ommitingFields()
	fmt.Println("\nForcing Fields to be Encoded as Strings:")
	forcingFieldsAsStrings()
	fmt.Println("\nEncoding Interfaces:")
	encodeInterfaces()
}

func encodeInterfaces() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct{
		Product:  &Kayak,
		Discount: 10.50,
	}
	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
	encoder.Encode(namedItems)
	fmt.Print(writer.String())
}

func forcingFieldsAsStrings() {
	type DiscountedProduct struct {
		*Product `json:",omitempty"`
		Discount float64 `json:",string"`
	}

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct{
		Product:  &Kayak, //Pointer here. No copy created!
		Discount: 10.50,
	}
	encoder.Encode(&dp)
	fmt.Print(writer.String())
	// Results in Discount result to be "10.5" as string
}

func ommitingFields() {
	type DiscountedProductOmmitField struct {
		*Product `json:"product"` // Product is there
		Discount float64          `json:"-"` // But the Discount is ommited
	}

	type DiscountedProductOmmitEmpty struct {
		*Product `json:"product,omitempty"`
		Discount float64 `json:"-"`
	}

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProductOmmitField{
		Product:  &Kayak, //Pointer here. No copy created!
		Discount: 10.50,
	}
	dp2 := DiscountedProductOmmitField{
		Discount: 10.50} // Product itself isn't there!

	dp3 := DiscountedProductOmmitEmpty{
		Discount: 10.50} // Ommited empty.
	encoder.Encode(&dp)
	encoder.Encode(&dp2) // Returns: {"product":null}
	encoder.Encode(&dp3) // Returns: {}

	fmt.Print(writer.String())

}

func customEncoding() {
	// How a struct is encoded can be customized using struct tags,
	// which are string literals that follow fields.
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProductTag{
		Product:  &Kayak, //Pointer here. No copy created!
		Discount: 10.50,
	}
	encoder.Encode(&dp)
	fmt.Print(writer.String())
}

func promoEffect() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct{
		Product:  &Kayak, //Pointer here. No copy created!
		Discount: 10.50,
	}
	encoder.Encode(&dp)
	fmt.Print(writer.String())
	// Resulting:
	// {"Name":"Kayak","Category":"Watersports","Price":279,"Discount":10.5}
}

func encodingStructs() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(Kayak)
	fmt.Print(writer.String())
}

func encodingMaps() {
	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(m)
	fmt.Print(writer.String())
}

// Encoding Arrays and Slices
func encodingArraysAndSlices() {
	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice) // Produces: "S2F5YWs=" - base64 encoded string
	fmt.Print(writer.String())

	// Decode Base64 string encoded from byteSlice
	b, err := base64.StdEncoding.DecodeString("S2F5YWs=")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", b)
}

// Encoding some of the basic Go types
func encodingBasicTypes() {
	var b bool = true
	var str string = "Hello"
	var fval float64 = 99.99
	var ival int = 200
	var pointer *int = &ival
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	for _, val := range []interface{}{b, str, fval, ival, pointer} {
		encoder.Encode(val)
	}
	fmt.Print(writer.String())
}
