package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func main() {

	fmt.Println("\nDecoding Basic Data Types:")
	decodingBasicDataTypes()
	fmt.Println("\nDecoding numbers:")
	decodingNumbers()
	fmt.Println("\nSpecifying Types for Decoding:")
	specifyingTypesForDecoding()
	fmt.Println("\nDecoding Arrays:")
	decodingArrays()
	fmt.Println("\nDecoding Arrays - Known structure:")
	decodingKnownArrays()
}

func decodingKnownArrays() {
	reader := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	ints := []int{}
	mixed := []interface{}{}
	vals := []interface{}{&ints, &mixed}
	decoder := json.NewDecoder(reader)
	for i := 0; i < len(vals); i++ {
		err := decoder.Decode(vals[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", ints, ints)
	Printfln("Decoded (%T): %v", mixed, mixed)
}

func decodingArrays() {
	reader := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	vals := []interface{}{}
	decoder := json.NewDecoder(reader)
	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}
	/*
		       The Decoder doesn't try to figure out if a JSON array can be represented using
			   a single Go type and decodes every array into an empty interface slice
	*/
}

func specifyingTypesForDecoding() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	var bval bool
	var sval string
	var fpval float64
	var ival int
	// Now the vlaues are known. So the interface can be non-empty
	vals := []interface{}{&bval, &sval, &fpval, &ival}
	decoder := json.NewDecoder(reader)
	// for i := 0; i < len(vals); i++ {
	// 	err := decoder.Decode(vals[i])
	// 	if err != nil {
	// 		Printfln("Error: %v", err.Error())
	// 		break
	// 	}
	// }

	// for - range loop can be used here too:
	for i := range vals {
		err := decoder.Decode(vals[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}

	Printfln("Decoded (%T): %v", bval, bval)
	Printfln("Decoded (%T): %v", sval, sval)
	Printfln("Decoded (%T): %v", fpval, fpval)
	Printfln("Decoded (%T): %v", ival, ival)
}

func decodingNumbers() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	vals := []interface{}{}
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		if num, ok := val.(json.Number); ok {
			if ival, err := num.Int64(); err == nil {
				Printfln("Decoded Integer: %v", ival)
			} else if fpval, err := num.Float64(); err == nil {
				Printfln("Decoded Floating Point: %v", fpval)
			} else {
				Printfln("Decoded String: %v", num.String())
			}
		} else {
			Printfln("Decoded (%T): %v", val, val)
		}
	}
}

func decodingBasicDataTypes() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	vals := []interface{}{}
	decoder := json.NewDecoder(reader)
	for {
		/*
		  The Decoder is able to select the appropriate Go data type
		  for JSON values, and this is achieved by providing
		  a pointer to an empty interface as the argument
		  to the Decode method, like this:
		*/
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)

		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}
}
