package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cycles = map[int]bool{20: true, 60: true, 100: true, 140: true, 180: true, 220: true}
var newlineCycles = map[int]bool{40: true, 80: true, 120: true, 160: true, 200: true}

func parseInstructions(input string) []int {
	var instructions []int
	for _, line := range strings.Split(string(input), "\n") {
		if line == "noop" {
			instructions = append(instructions, 0)
		} else {
			i, _ := strconv.Atoi(strings.Split(line, " ")[1])
			instructions = append(instructions, i)
		}
	}

	return instructions
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	crt := ""
	cycle := 0
	register := 1
	signalStrength := 0
	for _, instruction := range parseInstructions(string(file)) {
		var nCycles int
		if instruction == 0 {
			nCycles = 1
		} else {
			nCycles = 2
		}

		for i := 0; i < nCycles; i++ {
			var ok bool

			_, ok = newlineCycles[cycle]
			if ok {
				crt += "\n"
			}

			if (register-1 <= cycle%40) && (cycle%40 <= register+1) {
				crt += "#"
			} else {
				crt += "."
			}

			_, ok = cycles[cycle]
			if ok {
				signalStrength += cycle * register
			}

			cycle++
		}
		register += instruction
	}

	fmt.Printf("Signal strength: %d\n", signalStrength)
	fmt.Println(crt)
}
