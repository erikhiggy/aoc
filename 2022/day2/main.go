package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	WIN = 6
	LOSE = 0
	DRAW = 3
)

func main() {
	in := readInput()
	processed := processInput(in)
	part1Score := playGame(processed)
	part2Score := playGame2(processed)
	fmt.Print("Part 1: ")
	fmt.Println(part1Score)
	fmt.Print("Part 2: ")
	fmt.Println(part2Score)
}

func playGame2(rounds [][]string) int {
	totalScore := 0

	for _, round := range rounds {
		totalScore += rps2(round[0], round[1]) + convertResultToScore(round[1])
	}
	return totalScore
}

func playGame(rounds [][]string) int {
	totalScore := 0

	for _, round := range rounds {
		totalScore += rps(round[0], round[1]) + convertLetterToScore(round[1])
	}
	return totalScore
}

func rps2(player1Hand string, result string) int {
	// draw
	if player1Hand == "A" && result == "Y" {
		return 1 // X
	}
	if player1Hand == "B" && result == "Y" {
		return 2 // Y
	}
	if player1Hand == "C" && result == "Y" {
		return 3 // Z
	}

	// win
	if player1Hand == "A" && result == "Z" {
		return 2 // Y
	}
	if player1Hand == "B" && result == "Z" {
		return 3 // Z
	}
	if player1Hand == "C" && result == "Z" {
		return 1 // X
	}

	// lose
	if player1Hand == "A" && result == "X" {
		return 3 // Z
	}
	if player1Hand == "B" && result == "X" {
		return 1 // X
	}
	if player1Hand == "C" && result == "X" {
		return 2 // Y
	}

	// unrecognized combo
	return -1
}

func rps(player1Hand string, player2Hand string) int  {
	// draw
	if player1Hand == "A" && player2Hand == "X" {
		return DRAW
	}
	if player1Hand == "B" && player2Hand == "Y" {
		return DRAW
	}
	if player1Hand == "C" && player2Hand == "Z" {
		return DRAW
	}

	// lose
	if player1Hand == "A" && player2Hand == "Z" {
		return LOSE
	}
	if player1Hand == "B" && player2Hand == "X" {
		return LOSE
	}
	if player1Hand == "C" && player2Hand == "Y" {
		return LOSE
	}

	// win
	if player1Hand == "A" && player2Hand == "Y" {
		return WIN
	}
	if player1Hand == "B" && player2Hand == "Z" {
		return WIN
	}
	if player1Hand == "C" && player2Hand == "X" {
		return WIN
	}

	// unrecognized combo
	return -1
}

func processInput(input []string) [][]string {
	processed := make([][]string, len(input))

	for i, in := range input {
		processed[i] = strings.Split(in, " ")
	}

	return processed
}

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func convertLetterToScore(hand string) int {
	switch hand {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}

	// unrecognized hand
	return -1
}

func convertResultToScore(result string) int {
	switch result {
	case "X":
		return LOSE
	case "Y":
		return DRAW
	case "Z":
		return WIN
	}

	// unrecognized result
	return -1
}