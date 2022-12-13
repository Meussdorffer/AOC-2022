package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type assignment struct {
	low  int
	high int
}

// Parse assignment pairs into structs that can be compared easily.
func parseAssignmentPair(pair string) (assignment, assignment) {
	split := strings.Split(pair, ",")
	s1 := strings.Split(split[0], "-")
	s2 := strings.Split(split[1], "-")

	a11, _ := strconv.Atoi(s1[0])
	a12, _ := strconv.Atoi(s1[1])
	a21, _ := strconv.Atoi(s2[0])
	a22, _ := strconv.Atoi(s2[1])

	a1 := assignment{a11, a12}
	a2 := assignment{a21, a22}

	return a1, a2
}

// Compare assignment pairs for overlapping assignments.
// Returns 1 if pairs overlap completely, 0 otherwise.
func checkAssignmentsOverlapPart1(a1 assignment, a2 assignment) int {
	overlap := 0

	if a1.low <= a2.low && a1.high >= a2.high ||
		a2.low <= a1.low && a2.high >= a1.high {
		overlap = 1
	}

	return overlap
}

// Compare assignment pairs for overlapping assignments.
// Returns 1 if pairs partially overlap, 0 otherwise.
func checkAssignmentsOverlapPart2(a1 assignment, a2 assignment) int {
	overlap := 0

	if a1.high >= a2.low && a1.low <= a2.low ||
		a2.high >= a1.low && a2.low <= a1.low {
		overlap = 1
	}

	return overlap
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	fileContent := string(file)

	// part 1
	overlappingAssignmentPairs := 0
	for _, pair := range strings.Split(fileContent, "\n") {
		a1, a2 := parseAssignmentPair(pair)
		overlappingAssignmentPairs += checkAssignmentsOverlapPart1(a1, a2)
	}
	fmt.Println(overlappingAssignmentPairs)

	// part 2
	overlappingAssignmentPairs = 0
	for _, pair := range strings.Split(fileContent, "\n") {
		a1, a2 := parseAssignmentPair(pair)
		overlappingAssignmentPairs += checkAssignmentsOverlapPart2(a1, a2)
	}
	fmt.Println(overlappingAssignmentPairs)
}
