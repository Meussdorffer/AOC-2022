package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	fileContent := string(file)

	maxCalories := 0
	for _, line := range strings.Split(fileContent, "\n\n") {
		elfCalories := 0
		for _, val := range strings.Split(line, "\n") {
			calories, _ := strconv.Atoi(val)
			elfCalories += calories
		}

		if elfCalories >= maxCalories {
			maxCalories = elfCalories
		}
	}
	fmt.Println(maxCalories)
}
