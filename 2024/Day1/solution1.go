package main

import (
	"bufio"
	"math"
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
	solution := 0.0
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
	// fmt.Println("Left sorted:", left_array)
	// fmt.Println("Right sorted:", right_array)
	for i := range left_array {
		solution += math.Abs(float64(left_array[i] - right_array[i]))
	}
	println("Solution:", int(solution))
}
