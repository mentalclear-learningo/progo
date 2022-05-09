package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer{
		Name: "Alice", City: "New York",
	}
	payment := Payment{
		Currency: "USD", Amount: 100.50,
	}
	printDetails(product, customer, payment, 10, true)
}

func printDetails(values ...interface{}) {
	for _, elem := range values {
		fieldDetails := []string{}
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		if elemType.Kind() == reflect.Struct {
			for i := 0; i < elemType.NumField(); i++ {
				fieldName := elemType.Field(i).Name
				fieldVal := elemValue.Field(i)
				fieldDetails = append(fieldDetails,
					fmt.Sprintf("%v: %v", fieldName, fieldVal))
			}
			Printfln("%v: %v", elemType.Name(), strings.Join(fieldDetails, ", "))
		} else {
			Printfln("%v: %v", elemType.Name(), elemValue)
		}
	}
}

// This one will only proces types that are known in advance.
// SO if there is a new type it will require changes.
// func printDetails(values ...interface{}) {
// 	for _, elem := range values {
// 		// Type assertion here
// 		switch val := elem.(type) {
// 		case Product:
// 			Printfln("Product: Name: %v, Category: %v, Price: %v",
// 				val.Name, val.Category, val.Price)
// 		case Customer:
// 			Printfln("Customer: Name: %v, City: %v", val.Name, val.City)
// 		}
// 	}
// }

// Version 1 with no interfaces
// func printDetails(values ...Product) {
// 	for _, elem := range values {
// 		Printfln("Product: Name: %v, Category: %v, Price: %v",
// 			elem.Name, elem.Category, elem.Price)
// 	}
// }
