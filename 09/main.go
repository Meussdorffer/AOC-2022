package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//	type Rope struct {
//		headX int
//		headY int
//		tailX int
//		tailY int
//	}
type Knot struct {
	x int
	y int
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

// func moveRope(knots int, commands []string) []{

// }

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	commands := strings.Split(string(file), "\n")
	head := Knot{0, 0}
	tail := Knot{0, 0}
	var distinctTailPositions = map[[2]int]bool{{0, 0}: true}

	for _, command := range commands {
		dir := string(command[0])
		steps, _ := strconv.Atoi(string(command[2:]))

		for i := 0; i < steps; i++ {
			// apply head movement.
			switch dir {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}

			tailNeighbors := map[[2]int]bool{
				{tail.x, tail.y}:         true, // same position
				{tail.x + 1, tail.y}:     true, // right
				{tail.x - 1, tail.y}:     true, // left
				{tail.x, tail.y + 1}:     true, // up
				{tail.x, tail.y - 1}:     true, // down
				{tail.x + 1, tail.y + 1}: true, // upright
				{tail.x - 1, tail.y + 1}: true, // upleft
				{tail.x + 1, tail.y - 1}: true, // downright
				{tail.x - 1, tail.y - 1}: true, // downleft
			}

			// move tail if head is no longer its neighbor (or in same position).
			if !tailNeighbors[[2]int{head.x, head.y}] {
				// if head and tail don't share a row or column, we need the tail to move diagonally.
				sameRow := head.y == tail.y
				sameCol := head.x == tail.x

				switch dir {
				case "R":
					tail.x++
					if !(sameCol || sameRow) {
						tail.y = head.y
					}
				case "L":
					tail.x--
					if !(sameCol || sameRow) {
						tail.y = head.y
					}
				case "U":
					tail.y++
					if !(sameCol || sameRow) {
						tail.x = head.x
					}
				case "D":
					tail.y--
					if !(sameCol || sameRow) {
						tail.x = head.x
					}
				}
			}

			distinctTailPositions[[2]int{tail.x, tail.y}] = true

			fmt.Println(head, tail)
		}
	}

	// part 1
	fmt.Println(len(distinctTailPositions))
}
