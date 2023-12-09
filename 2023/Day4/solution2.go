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

func getWinningNumbers(myNumbers []int, winningNumbers []int) []int {
	var myWinningNumbers []int
	for _, one_my_number := range myNumbers {
		// Take one of my numbers and try it against winning numbers
		for _, one_winning_number := range winningNumbers {
			if one_my_number == one_winning_number {
				myWinningNumbers = append(myWinningNumbers, one_my_number)
				break
			}
		}
	}
	return myWinningNumbers
}

func main() {
	lines, _ := readLines("testdata.txt")
	solution := 0
	// Assume index = Card #, will need to chop off to match the lengths here
	cardCounts := make([]int, len(lines))
	for cardNumber, line := range lines {
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
			// Will =0 if a blank space lolz
			if convertedNumber == 0 {
				continue
			}
			if swapToMyNumbers {
				myNumbers = append(myNumbers, convertedNumber)
			} else {
				winningNumbers = append(winningNumbers, convertedNumber)
			}
		}
		matchingNumbers := getWinningNumbers(myNumbers, winningNumbers)
		amountOfWins := len(matchingNumbers)
		fmt.Println(matchingNumbers)
		fmt.Println("Bigger than 0:", amountOfWins, cardNumber)
		// Init all the cards at 0, so +1 by default
		cardCounts[cardNumber] += 1
		// Got matching cards, time to get some dupes going
		for j := 0; j < amountOfWins; j++ {
			// 1 win = cardNumber+1 gets duped, and then so on
			cardCounts[cardNumber+1+j] += cardCounts[cardNumber]
		}
	}
	fmt.Println(cardCounts)

	for _, item := range cardCounts {
		solution += item
	}
	println("Solution:", int(solution))
}
