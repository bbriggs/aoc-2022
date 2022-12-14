package main

import (
	"fmt"

	"github.com/bbriggs/aoc-2022/util"
)

func main() {
	fmt.Println(part1(util.GetInput()))

	fmt.Println(part2(util.GetInput()))
}

func part1(input []string) int {
	return getStartofMessage(input[0])
}

func part2(input []string) string {
	return ""
}

// getStartofMessage returns the index position of the beginning of the message in a given string
// The beginning of the message is defined as the first sequence of four characters that are all the different from each other
// Returns -1 if no message is found
func getStartofMessage(input string) int {
	for i := 0; i < len(input)-3; i++ { // -3 because we need 4 characters remaining to check
		if !hasDuplicates([]rune(input[i : i+4])) { // check if the next 4 characters are all different
			return i + 4 // return the index of the first character after the message
		}
	}
	return -1 // no message found
}

// hasDuplicates returns true if the given string contains duplicate characters
func hasDuplicates(input []rune) bool {
	for i, v := range input {
		for j, w := range input {
			if i != j && v == w {
				return true
			}
		}
	}
	return false
}

// getMessages returns the subslice of the input that contains the message
// It takes the original message and the index position of the start of the message
// It returns the slice of the original message that contains the message
func getMessages(input string, start int) string {
	return input[start:]
}
