package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bbriggs/aoc-2022/util"
)

type Move struct {
	from  int // the stack to move from
	to    int // the stack to move to
	count int // the number of elements to move
}

type State [][]string

// move executes a move on the state
func moveOne(m Move, s State) State {
	for i := 0; i < m.count; i++ {
		// pop the top element off the from stack
		top, from := pop(s[m.from])

		// push the element onto the to stack
		to := push(s[m.to], top)

		// update the state
		s[m.from] = from
		s[m.to] = to

	}

	return s // return the new state so the caller can write it back to the original
}

// tracks the board state
// imagine the input stacks rotated 90 degrees clockwise
var state = State{
	{"F", "T", "C", "L", "R", "P", "G", "Q"},
	{"N", "Q", "H", "W", "R", "F", "S", "J"},
	{"F", "B", "H", "W", "P", "M", "Q"},
	{"V", "S", "T", "D", "F"},
	{"Q", "L", "D", "W", "V", "F", "Z"},
	{"Z", "C", "L", "S"},
	{"Z", "B", "M", "V", "D", "F"},
	{"T", "J", "B"},
	{"Q", "N", "B", "G", "L", "S", "P", "H"}}

func main() {

	input := util.GetInput()

	// fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(input []string) string {
	for i, line := range input {
		move := parseMove(line)
		state = moveOne(move, state)
		log.Println("Index:", i, "Move:", move)
	}

	// get the top element of the each stack in order and join them
	var result []string
	for _, stack := range state {
		result = append(result, peek(stack))
	}

	return strings.Join(result, "")
}

func part2(input []string) string {

	for _, line := range input {
		move := parseMove(line)

		// pop the top N elements off the from stack
		chunk, from := popN(state[move.from], move.count)

		// update the state
		state[move.from] = from

		// push the elements onto the to stack
		state[move.to] = append(state[move.to], chunk...)
	}

	// get the top element of the each stack in order and join them
	var result []string
	for _, stack := range state {
		result = append(result, peek(stack))
	}

	return strings.Join(result, "")
}

// popN pops the last N elements from a slice and returns them
func popN(s []string, n int) ([]string, []string) {
	return s[len(s)-n:], s[:len(s)-n]
}

// pop removes the last element from a slice and returns it
func pop(s []string) (string, []string) {
	return s[len(s)-1], s[:len(s)-1]
}

// push adds an element to the end of a slice
func push(s []string, e string) []string {
	return append(s, e)
}

// peek returns the last element of a slice without removing it
func peek(s []string) string {
	return s[len(s)-1]
}

// parseMove parses a move from the input
func parseMove(s string) Move {
	splitMove := strings.Split(s, " ")

	count, _ := strconv.Atoi(splitMove[1])

	from, _ := strconv.Atoi(splitMove[3])
	from -= 1 // convert to 0-indexed

	to, _ := strconv.Atoi(splitMove[5])
	to -= 1 // convert to 0-indexed

	return Move{from, to, count}
}
