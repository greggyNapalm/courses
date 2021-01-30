package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getLine() string {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}

func main() {
	fmt.Println("Please input a Name:")
	NameLine := getLine()
	fmt.Println("Please input a Address:")
	AddrLine := getLine()

	userInputMap := map[string]string{
		"name":    string(NameLine),
		"address": string(AddrLine),
	}
	jsonStr, _ := json.Marshal(userInputMap)
	fmt.Println(string(jsonStr))
}
