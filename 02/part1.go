package main

import (
	"fmt"
	"os"
	"strings"
)

type Move int64

const (
	UndefinedMove Move = iota
	Rock
	Paper
	Scissors
)

type Result int64

const (
	Lose Result = iota
	Draw
	Win
)

func gameResult(opponentMove Move, playerMove Move) int {

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

func getPlayerMove(opponentMove Move, playerResult Result) Move {
	type key struct {
		oppMove Move
		result  Result
	}

	m := map[key]Move{
		{Rock, Lose}: Scissors,
		{Rock, Draw}: Rock,
		{Rock, Win}:  Paper,

		{Paper, Lose}: Rock,
		{Paper, Draw}: Paper,
		{Paper, Win}:  Scissors,

		{Scissors, Lose}: Paper,
		{Scissors, Draw}: Scissors,
		{Scissors, Win}:  Rock,
	}

	return m[key{opponentMove, playerResult}]
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	fileContent := string(file)

	// part 1
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
		totalPoints += gameResult(strategy[moves[0]], strategy[moves[1]])
	}

	fmt.Println(totalPoints)

	// part 2
	opponentStrategy := map[string]Move{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}

	playerStrategy := map[string]Result{
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}

	totalPoints = 0
	for _, game := range strings.Split(fileContent, "\n") {
		moves := strings.Split(game, " ")
		opponentMove := opponentStrategy[moves[0]]
		playerResult := playerStrategy[moves[1]]
		playerMove := getPlayerMove(opponentMove, playerResult)
		totalPoints += gameResult(opponentMove, playerMove)
	}

	fmt.Println(totalPoints)
}
