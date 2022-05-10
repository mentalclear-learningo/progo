package main

import (
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}
var mutex = sync.Mutex{}

func doSum(count int, val *int) {
	time.Sleep(time.Second)
	mutex.Lock() // To have fewer mutex ops, use before loop
	for i := 0; i < count; i++ {
		//mutex.Lock() // Lock on variable change
		*val++
		//mutex.Unlock() // Unock after variable change
	}
	mutex.Unlock()
	waitGroup.Done()
}

func main() {
	counter := 0
	numRoutines := 5
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go doSum(5000, &counter)
	}
	waitGroup.Wait()
	Printfln("Total: %v", counter)
}
