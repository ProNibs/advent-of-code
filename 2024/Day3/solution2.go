package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func multiplyStuff(input string) int {
	// Assume we just get a mul(X,X) string
	string_split := strings.Split(input, ",")
	re := regexp.MustCompile(`\d+`)
	left_number, _ := strconv.Atoi(re.FindStringSubmatch(string_split[0])[0])
	right_number, _ := strconv.Atoi(re.FindStringSubmatch(string_split[1])[0])
	// fmt.Println(left_number, right_number)
	return left_number * right_number
}

func main() {
	lines, _ := readLines("testdata.txt")
	solution := 0
	regex_matcher := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	for _, line := range lines {
		println("Raw:", line)
		fmt.Println(regex_matcher.FindAllStringSubmatch(line, -1))
		for _, one_string := range regex_matcher.FindAllStringSubmatch(line, -1) {
			solution += multiplyStuff(one_string[0])
		}
	}
	println("Solution:", solution)
}
