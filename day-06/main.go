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
	return getStartOfSequence(input[0], 4)
}

func part2(input []string) int {
	return getStartOfSequence(input[0], 14)
}

// getStartofSequence returns the index of the first character after a given sequence of unique characters
// it accepts a string and the length of the unique sequence as arguments and returns an the index of the first character after the sequence as an int
func getStartOfSequence(input string, seqLen int) int {
	for i := 0; i < len(input)-(seqLen-1); i++ { // Length of the input minus the length of the sequence plus 1
		if !hasDuplicates([]rune(input[i : i+seqLen])) { // check if the next N characters are all different
			return i + seqLen // return the index of the first character after the message
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
