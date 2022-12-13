package main

import (
	"fmt"
	"os"
	"strings"
)

func getNeighbors(x int, y int, grid [][]string) [][2]int {
	var neighbors [][2]int
	height := grid[x][y]

	rows := len(grid)
	cols := len(grid[0])
	if x+1 < rows && rune(grid[x+1][y])-height {
		neighbors = append(neighbors, [2]int{x + 1, y})
	}

	if x-1 >= 0 {
		neighbors = append(neighbors, [2]int{x - 1, y})
	}

	if y+1 < cols {
		neighbors = append(neighbors, [2]int{x, y + 1})
	}

	if y-1 >= 0 {
		neighbors = append(neighbors, [2]int{x, y - 1})
	}

	return neighbors
}

func parseInput(input string) ([][]string, int, int) {
	var grid [][]string
	x, y := -1, -1
	for rowIdx, line := range strings.Split(input, "\n") {
		row := strings.Split(line, "")
		grid = append(grid, row)

		if x == -1 {
			for colIdx, cell := range row {
				if cell == "S" {
					x = rowIdx
					y = colIdx
					break
				}
			}
		}
	}

	return grid, x, y
}

// 1  procedure BFS(G, root) is
// 2      let Q be a queue
// 3      label root as explored
// 4      Q.enqueue(root)
// 5      while Q is not empty do
// 6          v := Q.dequeue()
// 7          if v is the goal then
// 8              return v
// 9          for all edges from v to w in G.adjacentEdges(v) do
// 10              if w is not labeled as explored then
// 11                  label w as explored
// 12                  w.parent := v
// 13                  Q.enqueue(w)

func pathBfs(grid [][]string, x int, y int) int {
	q := Queue{}
	explored := map[[2]int]bool{{x, y}: true}
	parents := make(map[[2]int][2]int)

	q.Enqueue([2]int{x, y})

	var goalNode [2]int
	for !q.isEmpty() {
		node, _ := q.Dequeue()
		nX, nY := node[0], node[1]

		fmt.Printf("Exploring node %s", grid[nX][nY])

		if grid[nX][nY] == "E" {
			goalNode = [2]int{nX, nY}
			break
		}

		for _, neighbor := range getNeighbors(nX, nY, grid) {
			_, ok := explored[neighbor]
			if !ok {
				explored[neighbor] = true
				parents[neighbor] = [2]int{nX, nY}
				q.Enqueue(neighbor)
			}
		}
	}

	pathLen := 1
	curNode := goalNode
	for curNode != [2]int{x, y} {
		curNode = parents[curNode]
		pathLen++
	}

	return pathLen
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	grid, startX, startY := parseInput(string(file))
	fmt.Println(pathBfs(grid, startX, startY))

}
