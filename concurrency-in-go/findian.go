package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var isFirstI bool = false
	var isContainsA bool = false
	var lastChar string

	fmt.Println("Please input a string:")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	for pos, char := range strings.TrimSuffix(line, "\n") {
		if pos == 0 && strings.EqualFold(string(char), "I") {
			isFirstI = true
		}
		if strings.EqualFold(string(char), "A") {
			isContainsA = true
		}
		lastChar = string(char)
	}
	if isFirstI == true && isContainsA == true && strings.EqualFold(lastChar, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
