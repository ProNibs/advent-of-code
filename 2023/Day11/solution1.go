package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func expandColumns(original_map [][]int) [][]int {
	// Figure out which columns need to be expanded
	var columns_to_expand []int
	for j := range original_map[0] {
		expand_column := true
		for i := range original_map {
			if original_map[i][j] != 0 {
				expand_column = false
				break
			}
		}
		if expand_column {
			columns_to_expand = append(columns_to_expand, j)
		}
	}
	fmt.Println(columns_to_expand)
	// Actually add them left->right now
	for i, item := range columns_to_expand {
		original_map := slices.Insert(original_map, columns_to_expand[0], 0)
	}
	return original_map
}

func main() {
	lines, _ := readLines("testdata.txt")
	solution := 0
	galaxy_number := 1
	var original_map [][]int
	for _, line := range lines {
		var current_row []int
		for j, character := range line {
			if string(character) == "#" {
				current_row = append(current_row, galaxy_number)
				galaxy_number += 1
			} else {
				current_row = append(current_row, 0)
			}

			if j == len(lines)-1 {
				original_map = append(original_map, current_row)
				current_row = make([]int, len(lines[0]))
			}
		}
	}
	// fmt.Println(original_map)
	expandColumns(original_map)
	println("Solution:", solution)
}
