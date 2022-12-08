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
	// Score based on what you pick
	if X == "X" {
		score += 1
	} else if X == "Y" {
		score += 2
		mychoice = "B"
	} else if X == "Z" {
		score += 3
		mychoice = "C"
	}

	// Check if we tied or won; losses don't score
	// Check if a tie
	if A == mychoice {
		score += 3
		// Rock vs Paper is a win
	} else if A == "A" && mychoice == "B" {
		score += 6
		// Paper vs Scissors is a win
	} else if A == "B" && mychoice == "C" {
		score += 6
		// Lastly, Scissors vs Rock is a win
	} else if A == "C" && mychoice == "A" {
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
