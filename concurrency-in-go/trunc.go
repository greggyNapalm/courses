package main

import "fmt"

func main2() {
	var userInputFloat float64
	fmt.Println("Please input a number with floating point:")
	fmt.Scanf("%g", &userInputFloat)
	fmt.Println(int64(userInputFloat))

}
