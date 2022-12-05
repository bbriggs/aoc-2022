package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Get input from AOC website for day 1
// https://adventofcode.com/2020/day/1/input
func GetInput(day int) string {

	// get the website
	resp, err := http.Get(fmt.Sprintf("https://adventofcode.com/2020/day/%d/input", day))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// read the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// return the body as a string
	return string(body)
}
