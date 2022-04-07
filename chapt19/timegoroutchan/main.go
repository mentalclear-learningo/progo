package main

import (
	"fmt"
	"time"
)

func writeToChannel(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 1) // Adding 1 sec sleep
	}
	close(channel)
}

func putGoRutineToSleep() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}

func writeToChannelDef(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		//time.Sleep(time.Second * 1)
	}
	close(channel)
}

func deferExecOfFunc() {
	nameChannel := make(chan string)
	time.AfterFunc(time.Second*5, func() {
		writeToChannelDef(nameChannel)
	})
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}

func writeToChannelTimed(channel chan<- string) {
	Printfln("Waiting for initial duration...")
	_ = <-time.After(time.Second * 2)
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 1)
	}
	close(channel)
}

func receiveTimedNotif() {
	nameChannel := make(chan string)
	go writeToChannelTimed(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}

func main() {
	fmt.Println("\nPutting a Goroutine to Sleep:")
	putGoRutineToSleep()

	fmt.Println("\nDeferring Execution of a Function:")
	deferExecOfFunc()

	fmt.Println("\nReceiving Timed Notifications:")
	receiveTimedNotif()
}
