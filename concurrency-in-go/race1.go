package main

import (
	"fmt"
	"time"
)

func execute(some string) {

	// initializing a infinite for loop
	for i := 1; true; i++ {

		// prints string and number
		fmt.Println(some, i)

		// this makes the program sleep for
		// 100 milliseconds, wiz 10 seconds
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	// Simple go program with concurrency
	go execute("First")

	// Placing the go command infront of the
	// func call simply creates a goroutine
	execute("Second")

	// The goroutine ensures that both functions
	// execute simultaneously & successfully
	fmt.Println("program ends successfully")

	// This statement still won't execute because
	// the func call is stuck in an infinite loop
	// check the output
}
