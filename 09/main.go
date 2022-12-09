package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rope struct {
	headX int
	headY int
	tailX int
	tailY int
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	commands := strings.Split(string(file), "\n")
	rope := Rope{0, 0, 0, 0}
	var distinctTailPositions = map[[2]int]bool{{0, 0}: true}

	for _, command := range commands {
		dir := string(command[0])
		steps, _ := strconv.Atoi(string(command[2:]))

		for i := 0; i < steps; i++ {
			// apply head movement.
			switch dir {
			case "R":
				rope.headX++
			case "L":
				rope.headX--
			case "U":
				rope.headY++
			case "D":
				rope.headY--
			}

			tailNeighbors := map[[2]int]bool{
				{rope.tailX, rope.tailY}:         true, // same position
				{rope.tailX + 1, rope.tailY}:     true, // right
				{rope.tailX - 1, rope.tailY}:     true, // left
				{rope.tailX, rope.tailY + 1}:     true, // up
				{rope.tailX, rope.tailY - 1}:     true, // down
				{rope.tailX + 1, rope.tailY + 1}: true, // upright
				{rope.tailX - 1, rope.tailY + 1}: true, // upleft
				{rope.tailX + 1, rope.tailY - 1}: true, // downright
				{rope.tailX - 1, rope.tailY - 1}: true, // downleft
			}

			// move tail if head is no longer its neighbor (or in same position).
			if !tailNeighbors[[2]int{rope.headX, rope.headY}] {
				// if head and tail don't share a row or column, we need the tail to move diagonally.
				sameRow := rope.headY == rope.tailY
				sameCol := rope.headX == rope.tailX

				switch dir {
				case "R":
					rope.tailX++
					if !(sameCol || sameRow) {
						rope.tailY = rope.headY
					}
				case "L":
					rope.tailX--
					if !(sameCol || sameRow) {
						rope.tailY = rope.headY
					}
				case "U":
					rope.tailY++
					if !(sameCol || sameRow) {
						rope.tailX = rope.headX
					}
				case "D":
					rope.tailY--
					if !(sameCol || sameRow) {
						rope.tailX = rope.headX
					}
				}
			}

			distinctTailPositions[[2]int{rope.tailX, rope.tailY}] = true

			fmt.Println(rope)
		}
	}

	// part 1
	fmt.Println(len(distinctTailPositions))
}
