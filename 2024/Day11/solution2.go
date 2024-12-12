package day11

import (
	"bufio"
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

func reduce_left_side(input int) int {
	if input%2024 == 0 {
		input = input / 2024
		if input%2024 == 0 {
			// Still divisible, recursion plox
			return reduce_left_side(input)
		}
	}
	return input
}

func address_one_stone(input int) (int, int) {
	input_as_string := strconv.Itoa(input)
	if input == 0 {
		return 1, -1
	} else if len(input_as_string)%2 == 0 {
		left_side, _ := strconv.Atoi(input_as_string[:len(input_as_string)/2])
		right_side, _ := strconv.Atoi(input_as_string[(len(input_as_string))/2:])
		// Check if leftside is divisible by 2024 and do it
		left_side = reduce_left_side(left_side)
		return left_side, right_side
	} else {
		new_number := input * 2024
		// Let's try to keep these numbers small
		return new_number, -1
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
	number_of_loops := 42
	for i := 0; i < number_of_loops; i++ {
		var new_stone_array []int
		for _, one_int := range initial_stone_array {
			// fmt.Println(one_int, address_one_stone(one_int))
			temp1, temp2 := address_one_stone(one_int)
			new_stone_array = append(new_stone_array, temp1)
			if temp2 != -1 {
				new_stone_array = append(new_stone_array, temp2)
			}
		}
		// fmt.Println("New stone:", new_stone_array)
		// deep_copy := make([]int, len(new_stone_array))
		// Once a new stone array has been finalized, make the initial one equal it now
		// copy(deep_copy, new_stone_array)
		initial_stone_array = new_stone_array[:]
		println("On iteration:", i, "Array is", len(initial_stone_array), "long right now.")
	}
	// fmt.Println(initial_stone_array)
	println("Length for first array:", len(initial_stone_array))
	println("For data on 25 loops, solution should equal 186175")
}

// Workshopping how to reduce these big numbers
// 93339*2024 = 188918136 (len=9)
// Divide that by 2 = 94459068 (len=8, so not a good thing)
// Divide by 24 = 7871589 (len=7, sounds good)
// Divide by 20 made a decimal, so that's a no-go

// 10000*2024 = 20240000 (len is even)
// 20000*2024 = 40480000 (len is even)
// 30000*2024 = 60720000 (len is even)
// 40000*2024 = 80960000 (len is even)
// 50000*2024 = 101200000 (len is odd)
// ...
// 90000*2024 = 182160000 (len is odd)
