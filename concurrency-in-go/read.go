package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	type FileRow struct {
		FirstName string
		Lastname  string
	}
	var slice []FileRow
	var fpath string
	fmt.Println("Please input a target file path:")
	fmt.Scanf("%s", &fpath)

	file, err := os.Open(fpath)
	if err != nil {
		fmt.Printf(" > Failed with error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(line) > 0 {
			fileds := strings.Fields(line)
			slice = append(slice, FileRow{fileds[0], fileds[1]})
		}

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		os.Exit(1)
	}
	for _, row := range slice {
		fmt.Printf("%s:%s\n", row.FirstName, row.Lastname)
	}
}
