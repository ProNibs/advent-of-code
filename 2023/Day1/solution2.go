package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

type stringInt struct {
	name  string
	value string
}

var one = stringInt{name: "one", value: "1"}
var two = stringInt{name: "two", value: "2"}
var three = stringInt{name: "three", value: "3"}
var four = stringInt{name: "four", value: "4"}
var five = stringInt{name: "five", value: "5"}
var six = stringInt{name: "six", value: "6"}
var seven = stringInt{name: "seven", value: "7"}
var eight = stringInt{name: "eight", value: "8"}
var nine = stringInt{name: "nine", value: "9"}

func getFirstDigit(input string) string {
	input = rightAdjustedDetection(input)
	println("Left adjusted input:", input)
	for _, character := range input {
		if unicode.IsDigit(character) {
			return string(character)
		}
	}
	return "Nothing"
}

func getLastDigit(input string) string {
	input = leftAdjustedDetection(input)
	println("Right adjusted input:", input)
	for i := len(input) - 1; i > -1; i-- {
		// println(rune(input[i]))
		if unicode.IsDigit(rune(input[i])) {
			return string(input[i])
		}
	}
	return "Nothing"
}

func leftAdjustedDetection(input string) string {
	// Idea here is when the string "one" shows up, just replace with number 1
	input = strings.Replace(input, one.name, one.value+one.name, -1)
	input = strings.Replace(input, two.name, two.value+two.name, -1)
	input = strings.Replace(input, three.name, three.value+three.name, -1)
	input = strings.Replace(input, four.name, four.value+four.name, -1)
	input = strings.Replace(input, five.name, five.value+five.name, -1)
	input = strings.Replace(input, six.name, six.value+six.name, -1)
	input = strings.Replace(input, seven.name, seven.value+seven.name, -1)
	input = strings.Replace(input, eight.name, eight.value+eight.name, -1)
	input = strings.Replace(input, nine.name, nine.value+nine.name, -1)
	return input
}

func rightAdjustedDetection(input string) string {
	// Idea here is when the string "one" shows up, just replace with number 1
	input = strings.Replace(input, one.name, one.name+one.value, -1)
	input = strings.Replace(input, two.name, two.name+two.value, -1)
	input = strings.Replace(input, three.name, three.name+three.value, -1)
	input = strings.Replace(input, four.name, four.name+four.value, -1)
	input = strings.Replace(input, five.name, five.name+five.value, -1)
	input = strings.Replace(input, six.name, six.name+six.value, -1)
	input = strings.Replace(input, seven.name, seven.name+seven.value, -1)
	input = strings.Replace(input, eight.name, eight.name+eight.value, -1)
	input = strings.Replace(input, nine.name, nine.name+nine.value, -1)
	return input
}

func main() {
	lines, _ := readLines("data.txt")
	solution := 0
	for _, line := range lines {
		println("OG:", line)
		// line = leftAdjustedDetection(line)
		// println("Replaced version", line)
		current_number, _ := strconv.Atoi(getFirstDigit(line) + getLastDigit(line))
		println("Converted string:", current_number)
		solution += current_number
	}
	println(solution)
}
