package main

import (
	"fmt"
	"os"
	"strings"
)

func readInputs(inputFile string) [][]string {
	fileLines := strings.Split(inputFile, "\n")

	var inputs [][]string
	for _, line := range fileLines {
		inputs = append(inputs, strings.Split(line, ""))
	}

	return inputs
}

// detects the start marker of encoded message given N number of distinct
// characters to detect before returning.
func detectStartMarker(buffer []string, nDistinct int) int {
	startPos := nDistinct
	for i := 0; i < len(buffer)-(nDistinct+1); i += 1 {
		charMap := make(map[string]bool)
		for _, char := range buffer[i : i+nDistinct] {
			charMap[char] = true
		}

		if len(charMap) == nDistinct {
			break
		}

		startPos++
	}

	return startPos
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	inputs := readInputs(string(file))

	for i, str := range inputs {
		fmt.Printf("Input %d start-of-packet marker: %d, start-of-message marker: %d\n", i, detectStartMarker(str, 4), detectStartMarker(str, 14))
	}
}
