package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadStdInLine() string {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}

func inputStrToSlice(inS string) []int {
	slieceOfStr := strings.Fields(inS)
	ints := make([]int, len(slieceOfStr))
	for idx, el := range slieceOfStr {
		ints[idx], _ = strconv.Atoi(el)
	}
	return ints
}

func Swap(InS []int, ElIdx int) {
	InS[ElIdx], InS[ElIdx-1] = InS[ElIdx-1], InS[ElIdx]
}

func BubbleSort(InS []int) {
	IsSwapped := true
	for IsSwapped {
		IsSwapped = false
		for i := 1; i < len(InS); i++ {
			if InS[i] < InS[i-1] {
				Swap(InS, i)
				IsSwapped = true
			}
		}
	}
}

func main() {
	fmt.Println("please enter ten integers:")
	UserInputNumbers := ReadStdInLine()
	//	var UserInputNumbers = "5 2 8 1 3"
	unsortedNums := inputStrToSlice(UserInputNumbers)
	fmt.Printf("unsorted:%v\n", unsortedNums)
	BubbleSort(unsortedNums)
	fmt.Printf("sorted:%v\n", unsortedNums)
}
