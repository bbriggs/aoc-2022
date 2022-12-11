package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	day1()

	day2()
}

func day1() {

	// read input.txt
	input := getInput()

	var (
		commonChars []string // slice of common characters
		itemSum     int      // sum of the values of the common characters
	)

	// loop through each line in the input
	for _, line := range input {

		// split the line into two equal length strings
		half := getHalfString(line)

		// find the common character and append it to the slice
		commonChars = append(commonChars, getCommonCharacter(half[0], half[1]))
	}

	// convert the slice of common characters to a slice of runes and sum the values
	for _, v := range commonChars {
		itemSum += sumPrioities(stringToRuneSlice(v))
	}

	// print the sum of the values of the common characters
	println(itemSum)
}

func day2() {
	var (
		commonChars string
	)
	// Get the input
	input := getInput()

	// Group the input into groups of three
	groups := groupByThree(input)

	// Find the common chars in each group
	for _, group := range groups {
		c := getCommonCharacters(group[0], group[1], group[2])
		if c != "" {
			commonChars += c
		}
	}

	itemSum := sumPrioities(stringToRuneSlice(commonChars))

	fmt.Println(itemSum)

	fmt.Println("Day 2")
}

// getInput reads input.txt and returns the contents as a slice of strings split by newlines
func getInput() []string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(input), "\n")
}

// getHalfString returns a slice of two strings, each half the length of the input string
func getHalfString(s string) []string {
	half := len(s) / 2
	return []string{s[:half], s[half:]}
}

// getCommonCharacter returns the single common character between two strings
// I am making an assumption there is only one common character but multiple instances of it
func getCommonCharacter(s1, s2 string) string {

	for _, r := range s1 {

		for _, r2 := range s2 {

			if r == r2 {
				return string(r)
			}
		}
	}
	return "" // if no common character is found, return an empty string
}

// stringToRuneSlice converts a string to a slice of runes
// didn't realize I did't need a slice of runes until I was done but I'm leaving it in
func stringToRuneSlice(s string) []rune {
	return []rune(s)
}

// sumRunes sums the values of a slice of runes accounting for the offset such that
// a = 1, b = 2, c = 3, etc.
func sumPrioities(r []rune) int {
	var (
		sum    int // sum of the priorities
		offset int // offset to account for the difference between lowercase and uppercase letters
	)
	for _, v := range r {
		// lowercase letters start at 97 and end at 122
		// uppercase letters start at 65 and end at 90

		if v < 97 { // if the value is less than 97, it is a capital letter
			offset = 38 // Start of uppercase letters - offset needed to put A at 27
		} else { // otherwise it is a lowercase letter
			offset = 96 // Start of lowercase letters - offset needed to put a at 1
		}

		sum += int(v) - offset
	}
	return sum
}

// groupByThree takes the output of getInput and groups the lines into groups of three
func groupByThree(input []string) [][]string {
	var (
		allGroups [][]string // slice of slices of 3 lines from the input
		group     []string   // slice of 3 lines from the input
	)

	// loop through each line in the input
	for i, line := range input {

		// append the line to the group
		group = append(group, line)

		// if the group is full, append it to the slice of all groups and reset the group
		if (i+1)%3 == 0 {
			allGroups = append(allGroups, group)
			group = []string{}
		}
	}

	return allGroups
}

// getCommonCharacters returns the common characters between three strings
func getCommonCharacters(s1, s2, s3 string) string {

	// there is probably some fancy algorithm to do this but i'm a dropout with no ambition
	for _, r := range s1 {

		for _, r2 := range s2 {

			for _, r3 := range s3 {

				if r == r2 && r == r3 {
					return string(r)
				}
			}
		}
	}
	return "" // if no common character is found, return an empty string
}
