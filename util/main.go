package util

import (
	"io/ioutil"
	"strings"
)

type Elf struct {
	Calories int
}

type Elves []Elf

// SortElvesByCalories sorts the elves by calories
func SortElvesByCalories(elves Elves) {
	// Use literally any sorting algorithm you want because this is just AoC

	// I'm using a bubble sort because it's easy to implement and I'm lazy
	for i := 0; i < len(elves); i++ {
		for j := 0; j < len(elves)-1; j++ {
			if elves[j].Calories > elves[j+1].Calories {
				elves[j], elves[j+1] = elves[j+1], elves[j]
			}
		}
	}
}

// getInput reads input.txt and returns the contents as a slice of strings split by newlines
func GetInput() []string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(input), "\n")
}

// GroupByN groups the input into groups of n
func GroupByN(input []string, groupSize int) [][]string {
	var (
		allGroups [][]string // slice of slices of 3 lines from the input
		group     []string   // slice of 3 lines from the input
	)

	// loop through each line in the input
	for i, line := range input {

		// append the line to the group
		group = append(group, line)

		// if the group is full, append it to the slice of all groups and reset the group
		if (i+1)%groupSize == 0 {
			allGroups = append(allGroups, group)
			group = []string{}
		}
	}

	return allGroups
}
