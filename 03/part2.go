package main

import (
	"fmt"
	"os"
	"strings"
)

// identifies the common item appearing across three rucksacks.
func compareRucksacks(r1 string, r2 string, r3 string) rune {

	// make maps of each rune in string.
	m1 := make(map[rune]bool)
	for _, char := range r1 {
		m1[char] = true
	}

	m2 := make(map[rune]bool)
	for _, char := range r2 {
		m2[char] = true
	}

	m3 := make(map[rune]bool)
	for _, char := range r3 {
		m3[char] = true
	}

	// find common rune in each map.
	var commonRune rune
	for char, _ := range m1 {
		_, ok2 := m2[char]
		_, ok3 := m3[char]
		if ok2 && ok3 {
			commonRune = char
			break
		}
	}

	return commonRune
}

func convertRuneToPriority(char rune) int {
	priority := int(char)
	if priority >= 97 {
		priority -= 96 // sets "a" to 1
	} else if priority < 97 {
		priority -= 38 // sets "A" to 27
	}
	return priority
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	fileContent := string(file)
	fileLines := strings.Split(fileContent, "\n")

	prioritySum := 0
	for i := 0; i < len(fileLines); i += 3 {
		r1 := fileLines[i]
		r2 := fileLines[i+1]
		r3 := fileLines[i+2]
		prioritySum += convertRuneToPriority(compareRucksacks(r1, r2, r3))
	}

	fmt.Println(prioritySum)
}
