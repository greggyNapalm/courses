package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ReadStdInLine() string {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}

func makeRandSlice(lenth, maxVal int) []int {
	rv := make([]int, lenth, lenth)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < lenth; i++ {
		rv[i] = rand.Intn(maxVal)
	}
	return rv
}

func splitSlice(inSlice []int, chunksNum int) [][]int {
	var chunkLenth = int(math.Ceil(float64(len(inSlice)) / float64(chunksNum))) // works for positive int's only
	endPossition := 0
	rv := make([][]int, 0)
	for i := 0; i < len(inSlice); i += chunkLenth {
		endPossition = i + chunkLenth
		if endPossition > len(inSlice) {
			endPossition = len(inSlice)
		}
		rv = append(rv, inSlice[i:endPossition])
	}
	return rv
}

func sortChunk(name string, inSlice []int, c chan []int) {
	fmt.Println("routain #", name, ": input slice:\t", inSlice)
	// There is no requirement that each chunk has to be sorted via merge sort.
	//So I decided to cut the corner and to use built-in.
	sort.Ints(inSlice)
	fmt.Println("routain #", name, ": sorted slice:\t", inSlice)
	c <- inSlice
}

func mergeChunks(chunk1, chunk2 []int) []int {
	rv := make([]int, len(chunk1)+len(chunk2))
	i := 0
	for len(chunk1) > 0 && len(chunk2) > 0 {
		if chunk1[0] < chunk2[0] {
			rv[i] = chunk1[0]
			chunk1 = chunk1[1:]
		} else {
			rv[i] = chunk2[0]
			chunk2 = chunk2[1:]
		}
		i++
	}

	for j := 0; j < len(chunk1); j++ {
		rv[i] = chunk1[j]
		i++
	}
	for j := 0; j < len(chunk2); j++ {
		rv[i] = chunk2[j]
		i++
	}
	return rv
}

func main() {
	var slice []int
	chunksNum := 4
	chunksSorted := make([][]int, 0)

	fmt.Println("Please enter space separated array of int numbers or press enter to generate random.")
	usrInput := ReadStdInLine()
	if usrInput == "" {
		slice = makeRandSlice(19, 199)
	} else {
		for _, digit := range strings.Fields(usrInput) {
			i, _ := strconv.Atoi(digit)
			slice = append(slice, i)
		}
	}
	fmt.Println("Original unsorted array:", slice, "\n")
	chunksUnsorted := splitSlice(slice, chunksNum)

	c := make(chan []int)
	for idx, chunk := range chunksUnsorted {
		routain_name := strconv.Itoa(idx)
		go sortChunk(routain_name, chunk, c)
	}
	for i := 0; i < chunksNum; i++ {
		chunksSorted = append(chunksSorted, <-c)
	}
	// Better to implement recursive merge here to make it work with any number of chunks.
	//I'm quite busy, so it has to be good enough for the defined task.
	fmt.Println("Result sorted array:", mergeChunks(mergeChunks(chunksSorted[0], chunksSorted[1]), mergeChunks(chunksSorted[2], chunksSorted[3])), "\n")
}
