package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	crateLen    = 3
	drawingStop = " 1"
)

type Command struct {
	quantity  int
	fromStack int
	toStack   int
}

// splits problem input into two parts to be parsed further.
func readInputs(inputFile string) ([]string, []string) {
	fileLines := strings.Split(inputFile, "\n")

	// parse drawing input.
	lineNum := 0
	var drawingInput []string
	for _, line := range fileLines {
		if line[:2] == drawingStop {
			break
		}
		drawingInput = append(drawingInput, line)
		lineNum++
	}

	// drop two lines after drawing input to get command input.
	commandInput := fileLines[lineNum+2:]

	return drawingInput, commandInput
}

// parses the portion of the input specificng the current state of each crate stack.
func parseDrawing(drawingInput []string) []Stack {
	fileWidth := len(drawingInput[0])
	numStacks := fileWidth / int(crateLen)

	// Load crates into stacks by iterating over the drawing in reverse order to load from bottom to top.
	crateStacks := make([]Stack, numStacks)
	for lineNum := len(drawingInput) - 1; lineNum >= 0; lineNum-- {
		line := drawingInput[lineNum]
		if line[:2] == drawingStop {
			break
		}

		// increment i by size of crate + 1 to account for space delim.
		for i, stackNum := 0, 0; i < fileWidth; i, stackNum = i+crateLen+1, stackNum+1 {
			crate := string(line[i : i+crateLen][crateLen/2])

			if crate != " " {
				crateStacks[stackNum].Push(crate)
			}
		}
	}

	return crateStacks
}

// parses the portion of the input specificng the CrateMover commands.
func parseCommands(commandInput []string) []Command {
	var commands []Command

	for _, commandLine := range commandInput {
		// parse commands using regex.
		r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		commandStr := r.FindStringSubmatch(commandLine)

		// convert string commands to ints.
		q, _ := strconv.Atoi(commandStr[1])
		f, _ := strconv.Atoi(commandStr[2])
		t, _ := strconv.Atoi(commandStr[3])
		command := Command{
			quantity: q,

			// commands are 1-indexed, while our data structures are 0-indexed.
			fromStack: f - 1,
			toStack:   t - 1,
		}

		commands = append(commands, command)
	}

	return commands
}

// issues all commands to move crates between stacks using CrateMover9000.
func crateMover9000(crateStacks []Stack, commands []Command) []Stack {
	for _, command := range commands {
		for i := 0; i < command.quantity; i++ {
			crate, success := crateStacks[command.fromStack].Pop()
			if !success {
				fmt.Printf("Failed to pop from stack %d \n", command.fromStack)
				os.Exit(1)
			}

			crateStacks[command.toStack].Push(crate)
		}
	}

	return crateStacks
}

// issues all commands to move crates between stacks using CrateMover9001
// (i.e., multiple crates moved at once are done so in-order)
func crateMover9001(crateStacks []Stack, commands []Command) []Stack {
	for _, command := range commands {

		// we can maintain order when moving multiple crates in a stack by using a Stack.
		var bufferStack Stack

		for i := 0; i < command.quantity; i++ {
			crate, success := crateStacks[command.fromStack].Pop()
			if !success {
				fmt.Printf("Failed to pop from stack %d \n", command.fromStack)
				os.Exit(1)
			}

			bufferStack.Push(crate)
		}

		// unload the crates in the bufferStack into the destination stack in-order.
		for !bufferStack.isEmpty() {
			crate, _ := bufferStack.Pop()
			crateStacks[command.toStack].Push(crate)
		}
	}

	return crateStacks
}

// get topmost crate in each stack to form solution message.
func buildCrateMessage(stacks []Stack) string {
	var message []string
	for _, stack := range stacks {
		char, _ := stack.Pop()
		message = append(message, char)
	}
	return strings.Join(message, "")
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	// parse input.
	drawingInput, commandInput := readInputs(string(file))
	crateStacks1 := parseDrawing(drawingInput)
	crateStacks2 := parseDrawing(drawingInput)
	commands := parseCommands(commandInput)

	var sortedStacks []Stack

	// Part 1
	sortedStacks = crateMover9000(crateStacks1, commands)
	fmt.Println(buildCrateMessage(sortedStacks))

	// Part 2
	sortedStacks = crateMover9001(crateStacks2, commands)
	fmt.Println(buildCrateMessage(sortedStacks))

}
