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

func checkLineSafety(input string) bool {
	array := strings.Split(input, " ")
	isIncreasing := false
	first_level, _ := strconv.Atoi(array[0])
	second_level, _ := strconv.Atoi(array[1])
	if first_level < second_level {
		isIncreasing = true
	}
	for i, j := range array {
		// Skip if on the last item of the array
		if i+1 >= len(array) {
			continue
		}
		current_level, _ := strconv.Atoi(j)
		next_level, _ := strconv.Atoi(array[i+1])
		level_difference := current_level - next_level
		println(level_difference, current_level, next_level, !isIncreasing)
		// Check if even going in the right direction
		if level_difference > 0 && isIncreasing {
			println("Supposed to be increasing but it's not")
			return false
		}
		if level_difference < 0 && !isIncreasing {
			println("Supposed to be decreasing")
			return false
		}
		float64_level_diff := float64(level_difference)
		// Dumb math.Abs only deals with float64, gross
		if math.Abs(float64_level_diff) < 1.0 || math.Abs(float64_level_diff) > 3.0 {
			println("Too far apart")
			return false
		}
	}
	// Only if the whole level for loops correctly is it safe
	return true
}

func main() {
	lines, _ := readLines("data.txt")
	solution_one := 0
	for _, line := range lines {
		println("Raw:", line)
		if checkLineSafety(line) {
			solution_one += 1
		}
	}
	println("Solution:", solution_one)
}
