package main

import (
	"bufio"
	"os"
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

type Direction struct {
	left  string
	right string
}

func followTheMap(direction_input []string, route_table map[string]Direction) int {
	steps_taken := 0
	// Always start with "AAA"
	current_location := "AAA"
	for true {
		for _, direction := range direction_input {
			if direction == "L" {
				current_location = route_table[current_location].left
			}
			if direction == "R" {
				current_location = route_table[current_location].right
			}
			steps_taken += 1
			if current_location == "ZZZ" {
				return steps_taken
			}
		}
	}
	return steps_taken
}

func main() {
	lines, _ := readLines("testdata2.txt")
	solution := 0
	var directions []string
	route_table := make(map[string]Direction)
	for i, line := range lines {
		// println("Raw:", line)
		// First line is to say left or right stuff
		if i == 0 {
			for _, character := range line {
				directions = append(directions, string(character))
			}
		}
		if i > 1 {
			temp_array := strings.Split(line, " ")
			for j, one_string := range temp_array {
				if one_string == "=" {
					continue
				}
				temp_array[j] = strings.Trim(one_string, "(),")
			}
			route_table[temp_array[0]] = Direction{temp_array[2], temp_array[3]}
		}
	}
	// fmt.Println(directions)
	// fmt.Println(route_table)
	solution = followTheMap(directions, route_table)
	println("Solution:", solution)
}
