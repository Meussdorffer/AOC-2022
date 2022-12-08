package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parses a 2D array of ints from the problem input.
func parseArray(input string) [][]int {
	var arr [][]int
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, char := range strings.Split(line, "") {
			num, _ := strconv.Atoi(char)
			row = append(row, num)
		}
		arr = append(arr, row)
	}

	return arr
}

// creates a transposed version of an array.
func transposeArray(arr [][]int) [][]int {
	nRows := len(arr)
	nCols := len(arr[0])

	var arrT [][]int
	for colIdx := 0; colIdx < nCols; colIdx++ {
		var row []int
		for rowIdx := 0; rowIdx < nRows; rowIdx++ {
			row = append(row, arr[rowIdx][colIdx])
		}
		arrT = append(arrT, row)
	}

	return arrT
}

// convenience function for finding the max value in a slice.
func max(slice []int) int {
	val := slice[0]
	for _, i := range slice {
		if i > val {
			val = i
		}
	}
	return val
}

// convenience function for finding the min value in a slice.
func min(slice []int) int {
	val := slice[0]
	for _, i := range slice {
		if i < val {
			val = i
		}
	}
	return val
}

// convenience function for reversing a slice.
func reverseSlice(slice []int) []int {
	var reversedSlice []int
	for i := len(slice) - 1; i >= 0; i-- {
		reversedSlice = append(reversedSlice, slice[i])
	}

	return reversedSlice
}

// determines the number of visible trees (part 1)
func calcVisibleTrees(arr [][]int, arrT [][]int) int {
	nRows := len(arr)
	nCols := len(arr[0])
	visibleTrees := (nRows * 2) + (nCols * 2) - 4
	for rowIdx := 1; rowIdx < (nRows - 1); rowIdx++ {
		for colIdx := 1; colIdx < (nCols - 1); colIdx++ {
			treeHeight := arr[rowIdx][colIdx]

			maxLeft := max(arr[rowIdx][0:colIdx])
			maxRight := max(arr[rowIdx][colIdx+1:])
			maxUp := max(arrT[colIdx][0:rowIdx])
			maxDown := max(arrT[colIdx][rowIdx+1:])

			if treeHeight > min([]int{maxLeft, maxRight, maxUp, maxDown}) {
				visibleTrees++
			}
		}
	}

	return visibleTrees
}

// calculates the max scenic score for every tree in the array.
func calcMaxScenicScore(arr [][]int, arrT [][]int) int {
	nRows := len(arr)
	nCols := len(arr[0])
	maxScenicScore := 0
	for rowIdx := 0; rowIdx < nRows; rowIdx++ {
		for colIdx := 0; colIdx < nCols; colIdx++ {
			scenicScore := 1
			treeHeight := arr[rowIdx][colIdx]

			// gather all trees in each direction; reverse slices where necessary for easier iteration.
			leftSlice := reverseSlice(arr[rowIdx][0:colIdx])
			rightSlice := arr[rowIdx][colIdx+1:]
			upSlice := reverseSlice(arrT[colIdx][0:rowIdx])
			downSlice := arrT[colIdx][rowIdx+1:]

			// iterate over each element in the directional slice until a tree that blocks
			// the current tree's view is encountered.
			for _, slice := range [][]int{leftSlice, rightSlice, upSlice, downSlice} {
				viewDistance := 0
				for _, height := range slice {
					viewDistance++
					if height >= treeHeight {
						break
					}
				}
				scenicScore *= viewDistance
			}

			// save only the max scenicScore found.
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	return maxScenicScore
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	arr := parseArray(string(file))
	arrT := transposeArray(arr)

	println(calcVisibleTrees(arr, arrT))
	println(calcMaxScenicScore(arr, arrT))
}
