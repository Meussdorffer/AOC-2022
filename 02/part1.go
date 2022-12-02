package main

import (
	"fmt"
	"os"
	"strings"
)

type Move int64

const (
	Undefined Move = iota
	Rock
	Paper
	Scissors
)

func game_result(opponentMove Move, playerMove Move) int {

	resultPoints := 0 // base case: loss
	playerWin := (playerMove == Rock && opponentMove == Scissors) ||
		(playerMove == Paper && opponentMove == Rock) ||
		(playerMove == Scissors && opponentMove == Paper)

	if playerMove == opponentMove {
		// draw
		resultPoints = 3
	} else if playerWin {
		// win
		resultPoints = 6
	}

	return resultPoints + int(playerMove)
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	fileContent := string(file)

	strategy := map[string]Move{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	totalPoints := 0
	for _, game := range strings.Split(fileContent, "\n") {
		moves := strings.Split(game, " ")
		totalPoints += game_result(strategy[moves[0]], strategy[moves[1]])
	}

	fmt.Println(totalPoints)
}
