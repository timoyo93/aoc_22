package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	Action string
	Score  int
	Result string
}

const (
	Rock     string = "rock"
	Paper           = "paper"
	Scissors        = "scissors"
)

const (
	Loss string = "loss"
	Draw        = "draw"
	Win         = "win"
)

const (
	RockPoints    int = 1
	PaperPoints       = 2
	ScissorPoints     = 3
)

func main() {
	data := getInputData("input")
	result := solution(data)
	secondResult := solutionPartTwo(data)
	fmt.Printf("Total score: %d\n", result)
	fmt.Println("-------------------------")
	fmt.Printf("Total score (part 2): %d", secondResult)
}

func getInputData(filename string) []string {
	readFile, err := os.Open(fmt.Sprintf("%s.txt", filename))
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	return fileLines
}

func solution(data []string) int {
	totalScore := 0
	for _, line := range data {
		words := strings.Fields(line)
		s := getScore(words[0], words[1])
		totalScore = totalScore + s
	}
	return totalScore
}

func solutionPartTwo(data []string) int {
	totalScore := 0
	for _, line := range data {
		words := strings.Fields(line)
		s := getScoreForSecondPart(words[0], words[1])
		totalScore = totalScore + s
	}
	return totalScore
}

func getScore(opponentString, playerString string) int {
	var score int
	opponent := mapLetterToAction(opponentString)
	player := mapLetterToAction(playerString)

	// check if draw
	if opponent.Action == player.Action {
		score = player.Score + 3
	}

	// check if loose
	if opponent.Action == Rock && player.Action == Scissors {
		score = player.Score
	}
	if opponent.Action == Paper && player.Action == Rock {
		score = player.Score
	}
	if opponent.Action == Scissors && player.Action == Paper {
		score = player.Score
	}

	// check if won
	if opponent.Action == Rock && player.Action == Paper {
		score = player.Score + 6
	}
	if opponent.Action == Paper && player.Action == Scissors {
		score = player.Score + 6
	}
	if opponent.Action == Scissors && player.Action == Rock {
		score = player.Score + 6
	}
	return score
}

func getScoreForSecondPart(opponentString, playerSting string) int {
	var score int
	o := mapLetterToAction(opponentString)
	p := mapLetterToAction(playerSting)

	if p.Result == Draw {
		score = o.Score + 3
	}
	if p.Result == Win {
		winBonus := 6
		switch o.Action {
		case Rock:
			score = PaperPoints + winBonus
		case Paper:
			score = ScissorPoints + winBonus
		case Scissors:
			score = RockPoints + winBonus
		}
	}
	if p.Result == Loss {
		switch o.Action {
		case Rock:
			score = ScissorPoints
		case Paper:
			score = RockPoints
		case Scissors:
			score = PaperPoints
		}
	}

	return score
}

func mapLetterToAction(s string) Pair {
	var action Pair
	switch s {
	case "A", "X":
		action = Pair{Rock, RockPoints, Loss}
	case "B", "Y":
		action = Pair{Paper, PaperPoints, Draw}
	case "C", "Z":
		action = Pair{Scissors, ScissorPoints, Win}
	}
	return action
}
