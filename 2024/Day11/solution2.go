package main

import (
	"bufio"
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

func straight_up_diagonal_searches(input [][]string) [][]int {
	var output [][]int
	for i, one_row := range input {
		for j, one_character := range one_row {
			// Looking for an "M" to start this off but don't want index errors
			if one_character == "M" && i-2 >= 0 {
				// println("Found an M here:", i, j)
				// Scenario: diagonal up and going left
				if j-2 >= 0 {
					if input[i-1][j-1] == "A" {
						if input[i-2][j-2] == "S" {
							output = append(output, []int{i - 1, j - 1})
						}
					}
				}
				if j+2 < len(one_row) {
					// Scenario: diagonal up and going right
					if input[i-1][j+1] == "A" {
						if input[i-2][j+2] == "S" {
							output = append(output, []int{i - 1, j + 1})
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

func straight_down_diagonal_searches(input [][]string) [][]int {
	var output [][]int
	for i, one_row := range input {
		for j, one_character := range one_row {
			// Looking for an "M" to start this off but don't want index errors
			if one_character == "M" && i+2 < len(one_row) {
				// println("Found an X here:", i, j)
				// Scenario: diagonal down and going left
				if j-2 >= 0 {
					if input[i+1][j-1] == "A" {
						if input[i+2][j-2] == "S" {
							output = append(output, []int{i + 1, j - 1})
						}
					}
				}
				if j+2 < len(one_row) {
					// Scenario: diagonal down and going right
					if input[i+1][j+1] == "A" {
						if input[i+2][j+2] == "S" {
							output = append(output, []int{i + 1, j + 1})
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
	up_diagonal := straight_up_diagonal_searches(word_search_arrays)
	down_diagonal := straight_down_diagonal_searches(word_search_arrays)
	diagonals_array := append(up_diagonal, down_diagonal...)

	// fmt.Println(up_diagonal)
	// fmt.Println(down_diagonal)
	// fmt.Println(diagonals_array)
	// At this point, has the location of "A" for each up diagonal and down diagonal
	// Now just need to check if there is a pair in either array
	for i, one_coordinate_pair := range diagonals_array {
		// Check for if our pair is further to the right here
		// Safe assumption: there can only be two matches, never 3+ sharing the same "A"
		// Lazily for-looping again and just checking if we are i+1 essentially
		for j, latest_coordinate_pair := range diagonals_array {
			if j < i+1 {
				continue
			}
			// fmt.Println(one_coordinate_pair, latest_coordinate_pair)
			if slices.Equal(one_coordinate_pair, latest_coordinate_pair) {
				solution += 1
			}
		}
	}
	// fmt.Println(word_search_arrays)
	println("Solution:", solution)
}
