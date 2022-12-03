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

// TODO: fix
func getPlayerMove(opponentMove Move, playerResult Result) Move {
	type key struct {
		oppMove Move
		result  Result
	}

	m := map[key]Move{
		key{Rock, Lose}: Scissors,
		key{Rock, Draw}: Rock,
		key{Rock, Win}:  Paper,

		key{Paper, Lose}: Rock,
		key{Paper, Draw}: Paper,
		key{Paper, Win}:  Scissors,

		key{Scissors, Lose}: Paper,
		key{Scissors, Draw}: Scissors,
		key{Scissors, Win}:  Rock,
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

	totalPoints := 0
	for _, game := range strings.Split(fileContent, "\n") {
		moves := strings.Split(game, " ")
		opponentMove := opponentStrategy[moves[0]]
		playerResult := playerStrategy[moves[1]]
		playerMove := getPlayerMove(opponentMove, playerResult)
		totalPoints += gameResult(opponentMove, playerMove)
	}

	fmt.Println(totalPoints)
}
