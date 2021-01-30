package main

import (
	"fmt"
	"time"
)

var count int

func race(routainName string) {
	count++
	fmt.Printf("in %s. count == %d\n", routainName, count)
}

func main() {
	/*
		Run the program a few times. You will get an unpredictable combination of `routainName`  and `count` variable
		values. It depends on the scheduler's decision. You can't guarantee program results.
	*/
	go race("1st routain")
	go race("2nd routain")
	time.Sleep(2 * time.Second)
}
