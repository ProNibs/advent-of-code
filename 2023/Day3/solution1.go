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

func printArrayOfArrays(input [][]string) {
	for _, one_array := range input {
		fmt.Println(one_array)
	}
}

// This creates an array of arrays with "valid spots" =1 in those location
func createValidSymbolArray(input [][]string) [][]int {
	var output [][]int
	for i := range input {
		// Gotta init each row so we can set them
		output = append(output, make([]int, len(input[i])))
		for j, one_character := range input[i] {
			// Check if it's just a dot
			if one_character == "." {
				continue
			}
			// Golang is annoying to just check a string is a digit
			_, err := strconv.Atoi(one_character)
			// No error means it's a number
			if err == nil {
				continue
			}
			// If not a digit and not a dot, it's a symbol to me
			// Do just left, rights, and up -- we'll do down later
			// Don't need to check if already set, everything was init to be 0
			output[i][j] = 1

			if j > 0 {
				if i > 0 {
					output[i-1][j-1] = 1
				}
				output[i][j-1] = 1
			}
			if j < len(output[i])-1 {
				if i > 0 {
					output[i-1][j+1] = 1
				}
				output[i][j+1] = 1
			}
			if i > 0 {
				output[i-1][j] = 1
			}
		}
	}
	// Do downs now that everything has been init
	for i := range input {
		for j, one_character := range input[i] {
			// Check if it's just a dot
			if one_character == "." {
				continue
			}
			// Golang is annoying to just check a string is a digit
			_, err := strconv.Atoi(one_character)
			// No error means it's a number
			if err == nil {
				continue
			}

			if i < len(input)-1 {
				output[i+1][j] = 1
				if j > 0 {
					output[i+1][j-1] = 1
				}
				if j < len(input[i])-1 {
					output[i+1][j+1] = 1
				}
			}
		}
	}
	return output
}

func getNumberForOneRow(input []string, validSymbolArray []int) int {
	fmt.Println("Line", input, validSymbolArray)
	final_number_for_row := 0
	current_number := ""
	valid_symbol := false
	for i, one_character := range input {
		// Golang is annoying to just check a string is a digit
		_, err := strconv.Atoi(one_character)
		// No error means it's a number
		if err == nil {
			current_number += one_character
			if validSymbolArray[i] == 1 {
				valid_symbol = true
			}
			if i == len(input)-1 && valid_symbol {
				final_number, _ := strconv.Atoi(current_number)
				final_number_for_row += final_number
			}
		} else {
			// Not a number, reset the state of everything
			// First, add whatever number we got for the row in
			if valid_symbol {
				final_number, _ := strconv.Atoi(current_number)
				final_number_for_row += final_number
			}
			current_number = ""
			valid_symbol = false
		}
	}
	// println("Added numbers:", final_number_for_row)
	return final_number_for_row
}

func main() {
	lines, _ := readLines("data.txt")
	solution := 0
	var arrayOfArrays [][]string
	for _, line := range lines {
		// println("Raw:", line)
		var currentArray []string
		for _, one_character := range line {
			currentArray = append(currentArray, string(one_character))
		}
		arrayOfArrays = append(arrayOfArrays, currentArray)
	}
	// printArrayOfArrays(arrayOfArrays)
	validSymbols := createValidSymbolArray(arrayOfArrays)
	// fmt.Println(createValidSymbolArray(arrayOfArrays))
	for i, one_row := range arrayOfArrays {
		// if i == 7 {
		solution += getNumberForOneRow(one_row, validSymbols[i])
		// }
	}

	println("Solution:", solution)
}
