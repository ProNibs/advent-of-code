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
	// Figure out all the starting spots
	var starting_locations []string
	for k := range route_table {
		if strings.HasSuffix(k, "A") {
			starting_locations = append(starting_locations, k)
		}
	}

	for true {
		// I know I can use LCM for this
		// But I assume first z + 2 = another z and can't use LCM that assumption
		// Reddit says it works, but that's just luck
		for _, direction := range direction_input {
			for i, current_location := range starting_locations {
				if direction == "L" {
					starting_locations[i] = route_table[current_location].left
				}
				if direction == "R" {
					starting_locations[i] = route_table[current_location].right
				}
			}
			steps_taken += 1
			all_end_with_z := true
			for _, one_location := range starting_locations {
				if !strings.HasSuffix(one_location, "Z") {
					all_end_with_z = false
					break
				}
			}
			if all_end_with_z {
				return steps_taken
			}
		}
	}
	return steps_taken
}

func main() {
	lines, _ := readLines("testdata3.txt")
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
