package main

import (
	"bufio"
	"os"
	"sort"
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
	solution := 0
	var left_array []int
	var right_array []int
	for _, line := range lines {
		// println("Raw:", line)
		array := strings.Split(line, "   ")
		left_number, _ := strconv.Atoi(array[0])
		right_number, _ := strconv.Atoi(array[1])
		left_array = append(left_array, left_number)
		right_array = append(right_array, right_number)
		// fmt.Println(left_array, right_array)
	}
	sort.Ints(left_array)
	sort.Ints(right_array)
	for _, left_number := range left_array {
		matches := 0
		for _, right_number := range right_array {
			if left_number == right_number {
				matches += 1
			}
		}
		// println("Idea:", left_number, "*", matches)
		solution += left_number * matches
	}
	println("Solution:", int(solution))
}
