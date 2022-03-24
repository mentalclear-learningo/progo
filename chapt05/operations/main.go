package main

import (
	"fmt"
	"math"
)

func main() {
	price, tax := 275.00, 27.40
	sum := price + tax
	difference := price - tax
	product := price * tax
	quotient := price / tax
	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)

	fmt.Println("arithmeticOverflow:")
	arithmeticOverflow()

	fmt.Println("Using the Remainder Operator:")
	remainderOperator()

	fmt.Println("Using the Increment and Decrement Operators:")
	incrementAndDecrementOps()

	fmt.Println("Concatenating Strings:")
	concatStrings()

	fmt.Println("Comparison Operators:")
	compareOperators()

	fmt.Println("Comparing Pointers:")
	comparingPointers()

	fmt.Println("Understanding the Logical Operators:")
	logicalOps()

	fmt.Println("Explicit Type Conversions:")
	explicitTypeConvert()
}

func arithmeticOverflow() {
	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64
	fmt.Printf("%d \n%f \n", intVal, floatVal)
	fmt.Println(intVal * 2)
	fmt.Println(floatVal * 2)
	fmt.Println(math.IsInf((floatVal * 2), 0))
}

func remainderOperator() {
	posResult := 3 % 2
	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))
	fmt.Println(posResult)
	fmt.Println(negResult)
	fmt.Println(absResult)
}

func incrementAndDecrementOps() {
	value := 10.2
	value++
	fmt.Println(value)
	value += 2
	fmt.Println(value)
	value -= 2
	fmt.Println(value)
	value--
	fmt.Println(value)
}

func concatStrings() {
	greeting := "Hello"
	language := "Go"
	combinedString := greeting + ", " + language
	fmt.Println(combinedString)
}

func compareOperators() {
	first := 100
	const second = 200.00
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greaterThan := first > second
	greaterThanOrEqual := first >= second
	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lessThan)
	fmt.Println(lessThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)
}

func comparingPointers() {
	first := 100
	second := &first
	third := &first
	alpha := 100
	beta := &alpha
	// Memory location are compared here, not values!
	fmt.Println(second == third)
	fmt.Println(second == beta)

	// Compare values:
	fmt.Println(*second == *third)
	fmt.Println(*second == *beta)
}

func logicalOps() {
	maxMph := 50
	passengerCapacity := 4
	airbags := true
	familyCar := passengerCapacity > 2 && airbags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportsCar
	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)
}

func explicitTypeConvert() {
	kayak := 275
	soccerBall := 19.50
	total := float64(kayak) + soccerBall
	fmt.Println(total)
}
