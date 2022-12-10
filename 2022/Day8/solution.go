package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, _ := readLines("data.txt")

	var visible_trees = 0
	for idx, line := range lines {
		// Top and bottom rows always visible
		if idx == 0 || idx == len(lines)-1 {
			// Add the WHOLE row
			visible_trees += len(line)
		} else {
			// Not top/bottom, so time to loop through each one
			for i, j := range line {
				// Check for edges still :)
				if i == 0 || i == len(line)-1 {
					visible_trees += 1
					continue
				}
				var current_number, _ = strconv.Atoi(string(j))
				var visible_from_left = true
				var visible_from_right = true
				fmt.Println("Current number", current_number, "Row", idx+1, "Column", i+1)
				// Definitely not an edge, let's check left
				for x := 0; x < i; x++ {
					var left_number, _ = strconv.Atoi(string(line[x]))
					// If a number to the left is too big, not visible this way
					if left_number >= current_number {
						fmt.Println("Not visible from left")
						visible_from_left = false
						break
					}
				}
				if visible_from_left {
					visible_trees += 1
					continue
				}
				// Check right now
				for x := len(line) - 1; x > i; x-- {
					var right_number, _ = strconv.Atoi(string(line[x]))
					// If a number to the right is too big, not visible that way
					if right_number >= current_number {
						fmt.Println("Not visible from right")
						visible_from_right = false
						break
					}
				}
				if visible_from_right {
					visible_trees += 1
					continue
				}
				// Now to check for visibility up and down, which is more annoying
				// Let's start with up
				var visible_from_north = true
				for x := 0; x < idx; x++ {
					//fmt.Println("Row", x)
					var above_number, _ = strconv.Atoi(string(lines[x][i]))
					//fmt.Println("Above number is", above_number)
					if above_number >= current_number {
						fmt.Println("Not visible from north")
						visible_from_north = false
						break
					}
				}
				if visible_from_north {
					visible_trees += 1
					continue
				}
				var visible_from_south = true
				for x := len(lines) - 1; x > idx; x-- {
					//fmt.Println("Row", x)
					var below_number, _ = strconv.Atoi(string(lines[x][i]))
					//fmt.Println("Below number is", below_number)
					if below_number >= current_number {
						fmt.Println("Not visible from south")
						visible_from_south = false
						break
					}
				}
				if visible_from_south {
					visible_trees += 1
					continue
				}
			}
		}
	}
	fmt.Println(visible_trees)
}
