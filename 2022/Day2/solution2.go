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

func getScore(A string, X string) int {
	var score = 0
	mychoice := "A"
	// Figure out what mychoice should be
	// Pick same for ties
	if X == "Y" {
		mychoice = A
	} else if X == "Z" {
		// Supposed to win
		if A == "A" {
			mychoice = "B"
		} else if A == "B" {
			mychoice = "C"
		} else if A == "C" {
			mychoice = "A"
		}
	} else if X == "X" {
		// Supposed to lose
		if A == "A" {
			mychoice = "C"
		} else if A == "B" {
			mychoice = "A"
		} else if A == "C" {
			mychoice = "B"
		}
	}

	// Score based on what I picked
	if mychoice == "A" {
		score += 1
	} else if mychoice == "B" {
		score += 2
	} else if mychoice == "C" {
		score += 3
	}

	// Score based on if tie or win
	if X == "Y" {
		score += 3
	} else if X == "Z" {
		score += 6
	}

	return score
}

func main() {
	lines, _ := readLines("data.txt")
	var totalScore = 0
	for _, line := range lines {
		var choices = strings.Split(line, " ")
		//fmt.Println(choices[0], choices[1])
		totalScore += getScore(choices[0], choices[1])
	}
	fmt.Println(totalScore)
}
