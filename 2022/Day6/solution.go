package main

import (
	"bufio"
	"fmt"
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

func main() {
	lines, _ := readLines("data.txt")
	var input_string = lines[0]
	var solution_one_answered = false
	// Look at 4 characters at a time for repeats
	for i := 0; i < len(input_string); i++ {
		if i == len(input_string)-3 {
			fmt.Printf("Too far.")
			break
		}
		var four_char = input_string[i : i+4]
		var fourteen_char = input_string[i : i+14]

		var sol_one_repeated_letter = false
		var sol_two_repeated_letter = false
		// Check for repeats for solution one
		for j, one_char := range four_char {
			var one_char = string(one_char)
			if strings.Contains(four_char[j+1:], one_char) {
				sol_one_repeated_letter = true
				break
			}
		}
		if !sol_one_repeated_letter && !solution_one_answered {
			fmt.Println("Solution One:", i+4)
			solution_one_answered = true
		}
		// Check for repeats for solution two
		for j, one_char := range fourteen_char {
			var one_char = string(one_char)
			if strings.Contains(fourteen_char[j+1:], one_char) {
				sol_two_repeated_letter = true
				break
			}
		}
		if !sol_two_repeated_letter {
			fmt.Println("Solution Two:", i+14)
			break
		}
	}
}
