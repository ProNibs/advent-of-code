package main

import (
	"bufio"
	"math"
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

func convertToIntArray(input string) []int {
	array := strings.Split(input, " ")
	var int_array []int
	for _, string_as_int := range array {
		cast_to_int, _ := strconv.Atoi(string_as_int)
		int_array = append(int_array, cast_to_int)
	}
	return int_array
}

// The int here is if bool=false, int is the index of the array item to pop
func checkLineSafety(input []int) (bool, int) {
	isIncreasing := false
	first_level := input[0]
	second_level := input[1]
	if first_level < second_level {
		isIncreasing = true
	}
	for i, j := range input {
		// Skip if on the last item of the array
		if i+1 >= len(input) {
			continue
		}
		current_level := j
		next_level := input[i+1]
		level_difference := current_level - next_level
		// println(level_difference, current_level, next_level, !isIncreasing)
		// Check if even going in the right direction
		if level_difference > 0 && isIncreasing {
			// println("Supposed to be increasing but it's not")
			return false, i
		}
		if level_difference < 0 && !isIncreasing {
			// println("Supposed to be decreasing")
			return false, i
		}
		float64_level_diff := float64(level_difference)
		// Dumb math.Abs only deals with float64, gross
		if math.Abs(float64_level_diff) < 1.0 || math.Abs(float64_level_diff) > 3.0 {
			// println("Too far apart")
			return false, i
		}
	}
	// Only if the whole level for loops correctly is it safe
	return true, -1
}

func main() {
	lines, _ := readLines("data.txt")
	solution := 0
	for _, line := range lines {
		//println("Raw:", line)
		initial_int_array := convertToIntArray(line)
		// fmt.Println("Initial array", initial_int_array)
		state, pop_me := checkLineSafety(initial_int_array)
		if state {
			solution += 1
		} else {
			// Try 1 where we assume the left number is the one to be removed
			// Have to init the array for deep copy, but don't want it too long
			new_array := make([]int, pop_me)
			copy(new_array, initial_int_array[:pop_me])
			// fmt.Println("Deep copy array try 1", new_array)
			other_side := initial_int_array[pop_me+1:]
			new_array = append(new_array, other_side...)
			// fmt.Println("Fixed array try 1", new_array)
			new_state, _ := checkLineSafety(new_array)
			if new_state {
				solution += 1
			} else {
				// Try 2 where we assume the right number is the one to be removed
				new_array_two := make([]int, pop_me+1)
				copy(new_array_two, initial_int_array[:pop_me+1])
				other_side := initial_int_array[pop_me+2:]
				new_array_two = append(new_array_two, other_side...)
				// fmt.Println("Fixed array try 2", new_array_two)
				new_state, _ := checkLineSafety(new_array_two)
				if new_state {
					solution += 1
				} else {
					// I don't have logic to check if the first number is the bad guy, so eh, lazy
					// new_array_three := initial_int_array[1:]
					// fmt.Println("Fixed array try 3", new_array_three)
					new_state, _ := checkLineSafety(initial_int_array[1:])
					if new_state {
						solution += 1
					}
				}
			}
		}
	}
	println("Solution:", solution)
}
