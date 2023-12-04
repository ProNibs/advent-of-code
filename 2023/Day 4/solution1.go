package main

import (
	"bufio"
	"fmt"
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
	solution_one := 0
	for _, line := range lines {
		// Trash the Card #, who cares
		splitBySpace := strings.Split(line, " ")[2:]
		println("Raw:", line)
		var winningNumbers []int
		var myNumbers []int
		swapToMyNumbers := false
		for _, number := range splitBySpace {
			if number == "|" {
				swapToMyNumbers = true
				continue
			}
			convertedNumber, _ := strconv.Atoi(number)
			if swapToMyNumbers {
				myNumbers = append(myNumbers, convertedNumber)
			} else {
				winningNumbers = append(winningNumbers, convertedNumber)
			}
		}
		fmt.Println(winningNumbers)
		fmt.Println(myNumbers)
	}
	println("Solution:", solution_one)
}
