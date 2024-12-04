package main

import (
	"bufio"
	"os"
	"regexp"
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

func left_to_right_search(input []string) int {
	// Just recombine the input so we can do regex for "XMAS"
	one_line := strings.Join(input, "")
	re := regexp.MustCompile(`XMAS`)
	// fmt.Println(one_line, re.FindAllStringSubmatch(one_line, -1))
	return len(re.FindAllStringSubmatch(one_line, -1))
}

func right_to_left_search(input []string) int {
	// Just recombine the input so we can do regex for reverse of "XMAS"
	one_line := strings.Join(input, "")
	re := regexp.MustCompile(`SAMX`)
	// fmt.Println(one_line, re.FindAllStringSubmatch(one_line, -1))
	return len(re.FindAllStringSubmatch(one_line, -1))
}

func straight_down_search(input [][]string) int {
	output := 0
	for i, one_row := range input {
		for j, one_character := range one_row {
			// Looking for an "X" to start this off but don't want index errors
			if one_character == "X" && i+3 < len(input) {
				if input[i+1][j] == "M" {
					if input[i+2][j] == "A" {
						if input[i+3][j] == "S" {
							// XMAS straight down, yay
							output += 1
						}
					}
				}
			}
		}
	}
	// println("Straight down:", output, "Expected number: 1")
	return output
}

func straight_up_search(input [][]string) int {
	output := 0
	for i, one_row := range input {
		for j, one_character := range one_row {
			// Looking for an "X" to start this off but don't want index errors
			if one_character == "X" && i-3 >= 0 {
				if input[i-1][j] == "M" {
					if input[i-2][j] == "A" {
						if input[i-3][j] == "S" {
							// XMAS straight up, yay
							output += 1
						}
					}
				}
			}
		}
	}
	// println("Straight up:", output, "Expected number: 2")
	return output
}

func straight_up_diagonal_searches(input [][]string) int {
	output := 0
	for i, one_row := range input {
		for j, one_character := range one_row {
			// Looking for an "X" to start this off but don't want index errors
			if one_character == "X" && i-3 >= 0 {
				// println("Found an X here:", i, j)
				// Scenario: diagonal up and going left
				if j-3 >= 0 {
					if input[i-1][j-1] == "M" {
						if input[i-2][j-2] == "A" {
							if input[i-3][j-3] == "S" {
								// XMAS straight up, yay
								output += 1
							}
						}
					}
				}
				if j+3 < len(one_row) {
					// Scenario: diagonal up and going right
					if input[i-1][j+1] == "M" {
						if input[i-2][j+2] == "A" {
							if input[i-3][j+3] == "S" {
								// XMAS straight up, yay
								output += 1
							}
						}
					}
				}
				// println("Output ended up being", output)
			}
		}
	}
	// println("Diagonal and up:", output, "Expected: 8")
	return output
}

func straight_down_diagonal_searches(input [][]string) int {
	output := 0
	for i, one_row := range input {
		for j, one_character := range one_row {
			// Looking for an "X" to start this off but don't want index errors
			if one_character == "X" && i+3 < len(one_row) {
				// println("Found an X here:", i, j)
				// Scenario: diagonal down and going left
				if j-3 >= 0 {
					if input[i+1][j-1] == "M" {
						if input[i+2][j-2] == "A" {
							if input[i+3][j-3] == "S" {
								output += 1
							}
						}
					}
				}
				if j+3 < len(one_row) {
					// Scenario: diagonal down and going right
					if input[i+1][j+1] == "M" {
						if input[i+2][j+2] == "A" {
							if input[i+3][j+3] == "S" {
								output += 1
							}
						}
					}
				}
				// println("Output ended up being", output)
			}
		}
	}
	// println("Diagonal and down:", output, "Expected: 2")
	return output
}

func main() {
	lines, _ := readLines("data.txt")
	solution := 0
	var word_search_arrays [][]string
	for _, line := range lines {
		// println("Raw:", line)
		new_array := make([]string, len(line))
		for i, one_character := range line {
			new_array[i] = string(one_character)
		}
		word_search_arrays = append(word_search_arrays, new_array)
	}

	for _, one_line := range word_search_arrays {
		solution += left_to_right_search(one_line)
		// println("Left to right:", solution, "Expected: 3")
		solution += right_to_left_search(one_line)
	}
	// println("Right to left added:", solution, "Expected: 3+2=5")
	solution += straight_down_search(word_search_arrays)
	solution += straight_up_search(word_search_arrays)
	solution += straight_up_diagonal_searches(word_search_arrays)
	solution += straight_down_diagonal_searches(word_search_arrays)
	// fmt.Println(word_search_arrays)
	println("Solution:", solution)
}
