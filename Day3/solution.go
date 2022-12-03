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

func getRucksackLetters(rucksack string) string {
	// Split into compartments
	var compartment_one = rucksack[:len(rucksack)/2]
	var compartment_two = rucksack[len(rucksack)/2:]
	// Verify I split correctly
	if len(compartment_one) == len(compartment_two) {
		// Loop over individual letters from 1 to check if in 2
		for _, letter := range compartment_one {
			if strings.Contains(compartment_two, string(letter)) {
				return string(letter)
			}
		}
		return "0"
	} else {
		fmt.Println(len(compartment_one), len(compartment_two))
		return "Failed to split the rucksack in half properly."
	}
}

func getRucksackBadges(rucksack_one string, rucksack_two string, rucksack_three string) string {
	// Loop over each letter in one to check if in two
	for _, letter := range rucksack_one {
		if strings.Contains(rucksack_two, string(letter)) {
			if strings.Contains(rucksack_three, string(letter)) {
				return string(letter)
			}
		}
	}
	return "Didn't find the letter."
}

func scoreLetters(letter string) int {
	var ascii_decimal = int(letter[0])
	// Check if lowercase letter
	if ascii_decimal > 96 {
		return ascii_decimal - 96
	} else { // Otherwise, assume capitol letter
		return ascii_decimal - 38
	}
}

func main() {
	rucksack_list, _ := readLines("data.txt")
	// Solution One
	var totalScoreSolOne = 0
	for _, rucksack := range rucksack_list {
		totalScoreSolOne += scoreLetters(getRucksackLetters(rucksack))
	}
	println("Solution One:", totalScoreSolOne)
	// Solution Two
	var first_rucksack string
	var second_rucksack string
	var third_rucksack string
	var totalScoreSolTwo = 0
	for i, rucksack := range rucksack_list {
		if i%3 == 0 {
			first_rucksack = rucksack
		} else if i%3 == 1 {
			second_rucksack = rucksack
		} else if i%3 == 2 {
			third_rucksack = rucksack
			totalScoreSolTwo += scoreLetters(getRucksackBadges(first_rucksack, second_rucksack, third_rucksack))
		}
	}
	println("Solution Two:", totalScoreSolTwo)

}
