package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
	var solution_one = 0
	var solution_two = 0
	for _, line := range lines {
		// Split in half via the comma
		var comma_split = strings.Split(line, ",")
		// Split based on dash now
		var left_array = strings.Split(comma_split[0], "-")
		var right_array = strings.Split(comma_split[1], "-")

		// Gotta convert to int for proper compares
		var left_min, _ = strconv.Atoi(left_array[0])
		var left_max, _ = strconv.Atoi(left_array[1])
		var right_min, _ = strconv.Atoi(right_array[0])
		var right_max, _ = strconv.Atoi(right_array[1])

		// If one side is the same, auto complete
		if left_min == right_min {
			solution_one += 1
			solution_two += 1
		} else if left_max == right_max {
			solution_one += 1
			solution_two += 1
		} else if left_min > right_min {
			// For solution one
			if left_max <= right_max {
				solution_one += 1
				solution_two += 1
				// Check if max of right falls within left
			} else if left_min <= right_max && right_max <= left_max {
				solution_two += 1
			}
		} else if right_min > left_min {
			if right_max <= left_max {
				solution_one += 1
				solution_two += 1
				// Check if max of left falls within right
			} else if right_min <= left_max && left_max <= right_max {
				solution_two += 1
			}
		}
	}
	println(solution_one)
	println(solution_two)
}
