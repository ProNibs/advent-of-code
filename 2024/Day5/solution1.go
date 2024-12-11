package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

// Thanks StackOverflow
func PopIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func UpdateOrderArray(ordered_array []int, left_int int, right_int int) []int {
	// If left int NOT inside array, this is -1
	// println("Left|Right to place", left_int, right_int)
	index_in_ordered_array := slices.Index(ordered_array, left_int)
	if index_in_ordered_array == -1 {
		// Place new stuff in the front
		ordered_array = append([]int{right_int}, ordered_array...)
		ordered_array = append([]int{left_int}, ordered_array...)
	} else {
		ordered_array = slices.Insert(ordered_array, index_in_ordered_array+1, right_int)
	}
	// It's possible the left or right were already present, need to pop those out
	// Have to for loop to check this, gross
	for i, one_int := range ordered_array {
		for j, later_int := range ordered_array {
			// Lazy looping
			if j <= i {
				continue
			}
			// Found a dupe, the one further to the right is more accurate
			if one_int == later_int {
				ordered_array = PopIndex(ordered_array, i)
			}
		}
	}
	return ordered_array
}

func main() {
	order_lines, _ := readLines("testdata-order.txt")
	update_lines, _ := readLines("testdata-updates.txt")
	var ordered_array []int
	solution := 0
	fmt.Println("Before anything:", ordered_array)
	for _, line := range order_lines {
		// fmt.Println("Raw:")
		// if line == "61|53" {
		// 	// Stop here for debugging
		// 	println("Expected result as this point")
		// 	println("[75 97 61 47 53 29 13]")
		// 	break
		// }
		left_int, _ := strconv.Atoi(strings.Split(line, "|")[0])
		right_int, _ := strconv.Atoi(strings.Split(line, "|")[1])
		ordered_array = UpdateOrderArray(ordered_array, left_int, right_int)
		// fmt.Println("After action:", ordered_array)
	}
	// fmt.Println("Ordered Array:", ordered_array)
	// Do it again, because a 2nd run should resolve any ordering issues now that everything is seeded
	for _, line := range order_lines {
		left_int, _ := strconv.Atoi(strings.Split(line, "|")[0])
		right_int, _ := strconv.Atoi(strings.Split(line, "|")[1])
		ordered_array = UpdateOrderArray(ordered_array, left_int, right_int)
	}
	fmt.Println("Ordered Array:", ordered_array)
	// Now that we've gotten an array of the required left-to-right, we can begin
	for _, line := range update_lines {
		check_me := strings.Split(line, ",")
		for i, number := range check_me {
			number_to_check, _ := strconv.Atoi(number)
			fmt.Println(i, number_to_check)
			index := slices.Index(ordered_array, num)
			// Number isn't even in our ordered array, so doesn't matter
			if index == -1 {
				continue
			} else {
				println(index)
			}
		}
	}
	println("Solution:", solution)
}
