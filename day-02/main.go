package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var scoreMap = map[string]int{
	"A":    1, // rock
	"B":    2, // paper
	"C":    3, // scissors
	"X":    1, // rock
	"Y":    2, // paper
	"Z":    3, // scissors
	"lose": 0,
	"win":  6,
	"draw": 3,
}

func main() {
	// read input from input.txt
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// split the input into a slice of strings
	input := strings.Split(string(b), "\n")

	// keep track of the total score
	var totalScore int

	// loop through the input and calculate the score for each round
	// this is only for day 2 part 1
	for _, round := range input {
		// split the round into a slice of strings
		round := strings.Split(round, " ")

		// add the score to the total score
		totalScore += calculateScore(string(round[0]), string(round[1]))
	}

	// day 2 part 2
	// calculate the move based on the desired outcome
	// then calculate the score for that move and add it to the total score
	var day2total int

	for _, round := range input {
		// split the round into a slice of strings
		round := strings.Split(round, " ")

		// calculate the move to play to meet the desired outcome
		move := calculateMove(string(round[0]), string(round[1]))

		// calculate the score for the round and add it to the total score
		day2total += calculateScore(string(round[0]), move)

	}

	fmt.Println("Part 1: The total score is", totalScore)
	fmt.Println("Part 2: The total score is", day2total)
}

func calculateScore(player1, player2 string) int {
	switch {
	// handle ties first because they are the easiest
	case scoreMap[player1] == scoreMap[player2]:
		return scoreMap["draw"] + scoreMap[player2]

	// handle player 2 winning with rock
	case scoreMap[player2] == 1 && scoreMap[player1] == 3:
		return scoreMap["win"] + scoreMap[player2]

	// handle player 2 losing with rock
	case scoreMap[player2] == 1 && scoreMap[player1] == 2:
		return scoreMap["lose"] + scoreMap[player2]

	// handle player 1 winning with rock
	case scoreMap[player1] == 1 && scoreMap[player2] == 3:
		return scoreMap["lose"] + scoreMap[player2]

	// handle player 2 winning with anything else
	case scoreMap[player2] > scoreMap[player1]:
		return scoreMap["win"] + scoreMap[player2]

	// anything else means player 1 wins
	default:
		return scoreMap["lose"] + scoreMap[player2]
	}
}

// calculateMove takes two inputs and returns the move to play to meet the desired outcome
// Player 1's move is the first input, the desired outcome is the second input
// The desired outcome can be "X", "Y", or "Z" for "lose", "draw", or "win" respectively
func calculateMove(player1, desiredOutcome string) string {

	// map score to move to allow for arithmetic to determine the move to play
	scoreToMove := map[int]string{
		0: "Z", // handle case where player 2 wins against scissors or loses against rock
		1: "X",
		2: "Y",
		3: "Z",
	}

	// capture the score of player 1's move for readability
	p1 := scoreMap[player1]

	switch desiredOutcome {
	case "X": // player 2 wants to lose
		move := (p1 + 2) % 3
		return scoreToMove[move]
	case "Y": // player 2 wants to draw
		return scoreToMove[p1] // no arithmetic needed
	case "Z": // player 2 wants to win
		move := (p1 + 1) % 3
		return scoreToMove[move]
	default:
		return "invalid input"
	}
}
