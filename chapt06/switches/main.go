package main

import "fmt"

func main() {
	switchExamp()
	fmt.Println("Matching Multiple Values:")
	matchMutlipleValues()
	fmt.Println("Terminate case Statement Execution:")
	terminateCaseExecution()
	fmt.Println("Forcing Falling Through:")
	forceFallThrough()
	fmt.Println("Providing a default Clause:")
	providingDefaultClause()
	fmt.Println("Using an Initialization Statement:")
	usingInitStatement()
	fmt.Println("Good - Using an Initialization Statement:")
	usingInitStatementGood()
	fmt.Println("\nOmitting a Comparison Value:")
	ommitingCompValue()
}

func switchExamp() {
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K':
			fmt.Println("K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}

func matchMutlipleValues() {
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K', 'k':
			fmt.Println("K or k at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}

func terminateCaseExecution() {
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}

func forceFallThrough() {
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K':
			fmt.Println("Uppercase character")
			fallthrough
		case 'k':
			fmt.Println("k at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}

func providingDefaultClause() {
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		default:
			fmt.Println("Character", string(character), "at position", index)
		}
	}
}

// Troubled example
func usingInitStatement() {
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", counter/2)
		default:
			fmt.Println("Non-prime value:", counter/2)
		}
	}
}

func usingInitStatementGood() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", val)
		default:
			fmt.Println("Non-prime value:", val)
		}
	}
}

func ommitingCompValue() {
	for counter := 0; counter < 10; counter++ {
		switch {
		case counter == 0:
			fmt.Println("Zero value")
		case counter < 3:
			fmt.Println(counter, "is < 3")
		case counter >= 3 && counter < 7:
			fmt.Println(counter, "is >= 3 && < 7")
		default:
			fmt.Println(counter, "is >= 7")
		}
	}
}
