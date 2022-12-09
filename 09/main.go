package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

// issues commands to move the head of a rope with N number of knots.
// returns the positions visited by the final knot in the rope.
func moveRope(nKnots int, commands []string) int {
	var rope []Knot
	for i := 0; i < nKnots; i++ {
		rope = append(rope, Knot{0, 0})
	}

	ropeHead := &rope[0]
	ropeTail := &rope[nKnots-1]

	tailPositions := map[[2]int]bool{{0, 0}: true}
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
				head := &rope[knotIdx-1]
				tail := &rope[knotIdx]

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

				if tailNeighbors[[2]int{head.x, head.y}] {
					// if the head is a neighbor, this knot doesn't need to move.
					// if this knot doesn't need to move, neither do the knots further down the rope.
					break
				} else {
					// move tail in direction of head if head is not a neighbor.
					if tail.x != head.x {
						tail.x += (head.x - tail.x) / Abs(head.x-tail.x)
					}

					if tail.y != head.y {
						tail.y += (head.y - tail.y) / Abs(head.y-tail.y)
					}
				}
			}

			tailPositions[[2]int{ropeTail.x, ropeTail.y}] = true

		}
	}

	return len(tailPositions)
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	commands := strings.Split(string(file), "\n")

	// part 1
	fmt.Println(moveRope(2, commands))

	// part 2
	fmt.Println(moveRope(10, commands))
}
