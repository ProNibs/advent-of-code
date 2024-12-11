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

func address_one_stone(input int) []int {
	input_as_string := strconv.Itoa(input)
	if input == 0 {
		return []int{1}
	} else if len(input_as_string)%2 == 0 {
		left_side, _ := strconv.Atoi(input_as_string[:len(input_as_string)/2])
		right_side, _ := strconv.Atoi(input_as_string[(len(input_as_string))/2:])
		return []int{left_side, right_side}
	} else {
		return []int{input * 2024}
	}
}

func main() {
	lines, _ := readLines("data.txt")
	// solution := 0
	var initial_stone_array []int
	for _, line := range lines {
		println("Raw:", line)
		for _, one_int_as_string := range strings.Split(line, " ") {
			one_int, _ := strconv.Atoi(one_int_as_string)
			initial_stone_array = append(initial_stone_array, one_int)
		}
	}
	number_of_loops := 25
	for i := 0; i < number_of_loops; i++ {
		var new_stone_array []int
		for _, one_int := range initial_stone_array {
			// fmt.Println(one_int, address_one_stone(one_int))
			new_stone_array = append(new_stone_array, address_one_stone(one_int)...)
		}
		// fmt.Println("New stone:", new_stone_array)
		deep_copy := make([]int, len(new_stone_array))
		// Once a new stone array has been finalized, make the initial one equal it now
		copy(deep_copy, new_stone_array)
		initial_stone_array = deep_copy
	}
	fmt.Println(initial_stone_array)
	println("Solution:", len(initial_stone_array))
}
