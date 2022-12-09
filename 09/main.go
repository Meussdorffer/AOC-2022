package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	x                 int
	y                 int
	distinctPositions map[[2]int]bool
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

// issues commands to move the head of a rope with N number of knots.
// returns the state of each knot in the rope.
func moveRope(nKnots int, commands []string) []Knot {
	var rope []Knot
	for i := 0; i < nKnots; i++ {
		rope = append(rope, Knot{0, 0, map[[2]int]bool{{0, 0}: true}})
	}

	var ropeHead *Knot = &rope[0]

	for _, command := range commands {
		dir := string(command[0])
		steps, _ := strconv.Atoi(string(command[2:]))

		for i := 0; i < steps; i++ {
			// apply head movement.
			switch dir {
			case "R":
				ropeHead.x++
			case "L":
				ropeHead.x--
			case "U":
				ropeHead.y++
			case "D":
				ropeHead.y--
			}

			// apply movement for all other knots.
			for knotIdx := 1; knotIdx < nKnots; knotIdx++ {
				var head *Knot = &rope[knotIdx-1]
				var tail *Knot = &rope[knotIdx]

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

				// move tail in direction of head if head is not a neighbor.
				if !tailNeighbors[[2]int{head.x, head.y}] {
					if tail.x != head.x {
						tail.x += (head.x - tail.x) / Abs(head.x-tail.x)
					}

					if tail.y != head.y {
						tail.y += (head.y - tail.y) / Abs(head.y-tail.y)
					}
				}

				tail.distinctPositions[[2]int{tail.x, tail.y}] = true

			}
		}
	}

	return rope
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	commands := strings.Split(string(file), "\n")

	// part 1
	movedRope2Knots := moveRope(2, commands)
	fmt.Println(len(movedRope2Knots[len(movedRope2Knots)-1].distinctPositions))

	// part 2
	movedRope10Knots := moveRope(10, commands)
	fmt.Println(len(movedRope10Knots[len(movedRope10Knots)-1].distinctPositions))
}
