package main

import (
	"fmt"
	"os"
	"sort"
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
	contentGroups := strings.Split(fileContent, "\n\n")

	// part 1
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

	// part 2
	elfCalories := make([]int, len(contentGroups))
	for i, line := range contentGroups {
		for _, val := range strings.Split(line, "\n") {
			calories, _ := strconv.Atoi(val)
			elfCalories[i] += calories
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))

	top3Calories := 0
	for _, elfCalorie := range elfCalories[:3] {
		top3Calories += elfCalorie
	}

	fmt.Println(top3Calories)
}
