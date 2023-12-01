package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
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

func getFirstDigit(input string) string {
	for _, character := range input {
		if unicode.IsDigit(character) {
			return string(character)
		}
	}
	return "Nothing"
}

func getLastDigit(input string) string {
	for i := len(input) - 1; i > -1; i-- {
		// println(rune(input[i]))
		if unicode.IsDigit(rune(input[i])) {
			return string(input[i])
		}
	}
	return "Nothing"
}

func main() {
	lines, _ := readLines("data.txt")
	solution_one := 0
	for _, line := range lines {
		println(line)
		current_number, _ := strconv.Atoi(getFirstDigit(line) + getLastDigit(line))
		println("Converted string:", current_number)
		solution_one += current_number
	}
	println(solution_one)
}
