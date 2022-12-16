package main

import (
	"fmt"
	"strings"

	"github.com/bbriggs/aoc-2022/util"
)

// Initialize the root directory of the file system
var root = File{
	Name:     "/",
	IsDir:    true,
	Children: []File{},
	Parent:   nil,
}

// Set the current directory to the root
var dir = &root

func main() {
	input := util.GetInput()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	f := parseInput(input)

	return f.totalSize()
}

func part2(input []string) int {
	return 0
}

// File
type File struct {
	Name     string
	Size     int
	IsDir    bool
	Children []File
	Parent   *File // nil if root
}

func (f File) totalSize() int {
	size := f.Size

	// DFS to get total size of all children
	for _, child := range f.Children {
		size += child.totalSize()
	}

	return size
}

func (f *File) cd(name string) *File {
	// handle .. and filename
	for _, child := range f.Children {
		if child.Name == name {
			return &child
		}
	}

	// handle ..
	if name == ".." && f.Parent != nil { // nil if root
		return f.Parent
	}

	// handle /
	if name == "/" {
		return &root
	}

	return f
}

func (f File) GetName() string {
	if f.Parent == nil {
		return "/"
	}

	return f.Parent.GetName() + f.Name + "/"
}

func parseInput(input []string) *File {
	dir := &root

	for i, line := range input {
		fmt.Println(dir.GetName())
		// Split the line into the path and the size
		split := strings.Split(line, " ")
		if len(split) < 2 {
			continue // skip empty lines
		}

		if split[1] == "cd" {
			// update the global dir variable
			dir = dir.cd(split[2])
		} else if split[1] == "ls" {
			for i++; i < len(input); i++ { // start on the next line, run until end of ls response
				if strings.Split(input[i], " ")[0] == "$" {
					break // end of ls response
				}

				if len(input[i]) == 0 {
					continue // skip empty lines
				}

				// refresh the splitLine variable becuase we incremented i
				splitLine := strings.Split(input[i], " ")

				// We can just slap things in there without needing to do an if statement
				dir.Children = append(dir.Children, File{
					Name:     splitLine[1],
					Size:     util.MustAtoi(splitLine[0]), // returns 0 if not a number
					IsDir:    splitLine[0] == "dir",       // true if dir, false if file
					Children: []File{},
					Parent:   dir,
				})
			}
		}
	}
	return &root
}
