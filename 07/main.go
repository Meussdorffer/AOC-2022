package main

import (
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	fmt.Println(string(file))
}
