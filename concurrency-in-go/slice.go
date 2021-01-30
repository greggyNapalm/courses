package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var slice = make([]int, 3, 3)
	for {
		var userInput string
		fmt.Println("Please input a number or x to exit:")
		fmt.Scanf("%s", &userInput)
		if strings.EqualFold(userInput, "x") {
			break
		}
		userInputAsInt, _ := strconv.Atoi(userInput)
		slice = append(slice, userInputAsInt)
		sort.Ints(slice)
		fmt.Println(slice)
	}
}
