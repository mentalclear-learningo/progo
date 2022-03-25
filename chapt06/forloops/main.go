package main

import "fmt"

func main() {
	basicForLoop()
	fmt.Println("Incorporating the Condition into the Loop:")
	conditionForLoop()
	fmt.Println("Using Initialization and Completion Statements:")
	intiAndConditions()
	fmt.Println("Continuing Loops:")
	continuingLoops()

}

func basicForLoop() {
	counter := 0
	for {
		fmt.Println("Counter: ", counter)
		counter++
		if counter > 3 {
			break
		}
	}
}

func conditionForLoop() {
	counter := 0
	for counter <= 3 {
		fmt.Println("Counter:", counter)
		counter++
		// if (counter > 3) {
		//     break
		// }
	}
}

func intiAndConditions() {
	for counter := 0; counter <= 3; counter++ {
		fmt.Println("Counter:", counter)
		// counter++
	}
}

func continuingLoops() {
	for counter := 0; counter <= 3; counter++ {
		if counter == 1 {
			continue
		}
		fmt.Println("Counter:", counter)
	}
}
