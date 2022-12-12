package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bbriggs/aoc-2022/util"
)

func main() {
	input := util.GetInput()

	println("Part 1:", part1(input))
	println("Part 2:", part2(input))

}

func part1(input []string) int {
	count := 0

	// Loop through each group and compare the ranges to see if one is a superset
	for _, group := range input {
		// split the group into a tuple of ranges
		ranges := strings.Split(group, ",")

		if len(ranges) == 1 { // skip groups with only one range
			continue
		}

		if isSuperSet(ranges[0], ranges[1]) {
			fmt.Println("Contains a superset: ", ranges[0], ranges[1])
			count++
		}
	}

	return count
}

func part2(input []string) string {
	return ""
}

// isSuperSet returns true if one integer range is a superset of another
// input ranges must be formatted as "a-b"
// Example: isSuperSet("1-5", "2-4") returns true because 2-4 is a subset of 1-5
func isSuperSet(a, b string) bool {
	// parse the ranges
	s1, e1 := parseRange(a)
	s2, e2 := parseRange(b)

	switch {
	case s1 <= s2 && e1 >= e2: // a is a superset of b
		return true
	case s1 >= s2 && e1 <= e2: // b is a superset of a
		return true
	}
	return false
}

// parseRange parses a range string and returns the start and end values
// input ranges must be formatted as "a-b"
// Example: parseRange("1-5") returns 1, 5
func parseRange(a string) (int, int) {
	start, end := strings.Split(a, "-")[0], strings.Split(a, "-")[1]

	// convert the strings to ints
	s, _ := strconv.Atoi(start)
	e, _ := strconv.Atoi(end)

	return s, e
}
