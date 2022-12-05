package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bbriggs/aoc-2022/util"
)

type Elf struct {
	Calories int
}

type Elves []Elf

func main() {
	input := util.GetInput(1)

	elves := parseInput(input)

	// get the elf with the most calories
	var maxCalories int
	var elfIndex int
	for i, elf := range elves {
		if elf.Calories > maxCalories {
			maxCalories = elf.Calories
			elfIndex = i
		}
	}

	// print the result
	fmt.Println("The elf with the most calories is elf ", elfIndex+1, " with ", maxCalories, " calories.")

	println("Hello, World!")
}

// parse the input into a slice of []Elf
func parseInput(input string) Elves {

	// keep track of the current elf
	elfIndex := 0

	// Initialize a slice of Elves
	elves := Elves{}

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
			elves = append(elves, Elf{})
		}
	}

	return elves
}
