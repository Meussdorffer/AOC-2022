package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func max(slice []int) int {
	val := slice[0]
	for _, i := range slice {
		if i > val {
			val = i
		}
	}
	return val
}

func min(slice []int) int {
	val := slice[0]
	for _, i := range slice {
		if i < val {
			val = i
		}
	}
	return val
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	arr := parseArray(string(file))
	arrT := transposeArray(arr)

	// for _, line := range arr {
	// 	fmt.Println(line)
	// }
	// println("\n")
	// for _, line := range arrT {
	// 	fmt.Println(line)
	// }

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
	println(visibleTrees)

}
