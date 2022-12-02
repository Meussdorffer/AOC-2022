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

const (
	Undefined Result = iota
	Lose
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

	m := make(map[key]MediaItem)
	m[key{oppMove: opponentMove, result: playerResult}] = MediaItem{}
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

	playerStrategy := map[string]Move{
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}

	totalPoints := 0
	for _, game := range strings.Split(fileContent, "\n") {
		moves := strings.Split(game, " ")
		opponentMove := opponentStrategy[moves[0]]
		playerResult := playerStrategy[moves[1]]
		playerMove := Rock
		totalPoints += gameResult(opponentMove, playerMove)
	}

	fmt.Println(totalPoints)
}
