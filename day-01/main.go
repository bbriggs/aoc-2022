package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bbriggs/aoc-2022/util"
)

func main() {
	// get the input from input.txt
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	elves := parseInput(string(input))

	// find the elf with the most calories
	elfIndex, maxCalories := findMaxCalories(elves)

	// print the result
	fmt.Println("The elf with the most calories is elf", elfIndex+1, "with", maxCalories, "calories.")

	// find the sum of the top 3 elves' calories
	top3 := top3calories(elves)
	fmt.Println("The sum of the top 3 elves' calories is", top3)
}

// findMaxCalories finds the elf with the most calories and returns their index and the number of calories
func findMaxCalories(elves util.Elves) (int, int) {
	var maxCalories int
	var elfIndex int
	for i, elf := range elves {
		if elf.Calories > maxCalories {
			maxCalories = elf.Calories
			elfIndex = i
		}
	}
	return elfIndex, maxCalories
}

// top3calories finds the 3 elves with the most calories and returns the sum of their calories
func top3calories(elves util.Elves) int {
	// sort the elves by calories
	util.SortElvesByCalories(elves)

	// Start at the end of the slice and work backwards to get the top 3 elves
	var sum int
	for i := len(elves) - 1; i >= len(elves)-3; i-- {
		sum += elves[i].Calories
	}

	return sum
}

// parse the input into a slice of []Elf
func parseInput(input string) util.Elves {

	// keep track of the current elf
	elfIndex := 0

	// Initialize a slice of Elves
	elves := util.Elves{}

	// append the first elf to the slice
	elves = append(elves, util.Elf{})

	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			// get the calories as an int
			cal, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}

			// Increment the elf's calories by the amount in the line
			elves[elfIndex].Calories += cal
		} else {
			// if the line is empty, we've reached the end of the elf's input
			// so we increment the elfIndex and append a new Elf to the slice
			elfIndex++
			elves = append(elves, util.Elf{})
		}
	}

	return elves
}
