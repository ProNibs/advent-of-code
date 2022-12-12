package main

import (
	"bufio"
	"fmt"
	"math"
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

type Monkey struct {
	items            []int
	operation_first  string
	operator         string
	operation_second string
	// Tests are always divisible by X
	test int
	// Pass/fail is always pass to X monkey
	test_pass     int
	test_fail     int
	inspect_count int
}

func stressRelief(input int) (output int) {
	output = int(math.Floor(float64(input) / 3))
	return output
}

func monkeyOperation(first int, second int, operator string) int {
	if operator == "+" {
		return first + second
	} else if operator == "*" {
		return first * second
	}
	// Should only be one of the above
	return -1
}

func oneMonkeyInspection(input Monkey) (output Monkey, inspectItem int, monkeyThrow int) {
	// Assume there is an item to inpsect before this function is called
	inspectItem = input.items[0]
	output = input
	if len(input.items) == 1 {
		// This is the last item the monkey inspects
		output.items = make([]int, 0)
	} else { // Otherwise, pop it out
		output.items = output.items[1:]
	}
	var first = 0
	var second = 0
	if input.operation_first == "old" {
		first = inspectItem
	} else {
		// If not old, it's a number
		first, _ = strconv.Atoi(input.operation_first)
	}
	if input.operation_second == "old" {
		second = inspectItem
	} else {
		// If not old, it's a number
		second, _ = strconv.Atoi(input.operation_second)
	}
	// Worry level increased due to specific monkey and then relief!
	inspectItem = monkeyOperation(first, second, input.operator)
	inspectItem = stressRelief(inspectItem)
	// Time to test it
	if inspectItem%input.test == 0 {
		monkeyThrow = input.test_pass
	} else {
		monkeyThrow = input.test_fail
	}
	output.inspect_count += 1
	return output, inspectItem, monkeyThrow

}

func main() {
	lines, _ := readLines("data.txt")
	//var current_monkey_count = 0
	var current_monkey_struct = Monkey{}
	var monkey_list []Monkey
	for _, line := range lines {
		var line_array = strings.Split(line, " ")
		if strings.HasPrefix(line, "Monkey") {
			// Have to ignore the colon at the end to parse to int properly
			//var monkey_number, _ = strconv.Atoi(strings.Trim(line_array[1], ":"))
			//current_monkey_count = monkey_number
			current_monkey_struct = Monkey{}
		} else if strings.Contains(line, "Starting") {
			for idx, x := range line_array {
				// Starting at index 4 are real stuff
				if idx < 4 {
					continue
				}
				var current_item, _ = strconv.Atoi(strings.Trim(x, ","))
				current_monkey_struct.items = append(current_monkey_struct.items, current_item)
				current_monkey_struct.inspect_count = 0
			}
		} else if strings.Contains(line, "Operation:") {
			for idx, current_item := range line_array {
				if idx == 5 {
					current_monkey_struct.operation_first = string(current_item)
				} else if idx == 6 {
					current_monkey_struct.operator = string(current_item)
				} else if idx == 7 {
					current_monkey_struct.operation_second = string(current_item)
				}
			}
		} else if strings.Contains(line, "Test: ") {
			var test_number, _ = strconv.Atoi(line_array[5])
			current_monkey_struct.test = test_number
		} else if strings.Contains(line, "If true:") {
			var throw_number, _ = strconv.Atoi(line_array[9])
			current_monkey_struct.test_pass = throw_number
		} else if strings.Contains(line, "If false:") {
			var throw_number, _ = strconv.Atoi(line_array[9])
			current_monkey_struct.test_fail = throw_number
		} else if line == "" {
			// It's the line break, put the monkey into the list
			monkey_list = append(monkey_list, current_monkey_struct)
		}
	}
	// Last monkey to be defined doesn't have a line break or anything
	monkey_list = append(monkey_list, current_monkey_struct)
	// Initial monkey list is created
	fmt.Println("Initial List:", monkey_list, len(monkey_list))
	var rounds = 20
	for i := 0; i < rounds; i++ {
		for k, _ := range monkey_list {
			// Begin one round of inspections
			for range monkey_list[k].items {
				var item_thrown = 0
				var monkey_receiver = 0
				monkey_list[k], item_thrown, monkey_receiver = oneMonkeyInspection(monkey_list[k])
				monkey_list[monkey_receiver].items = append(monkey_list[monkey_receiver].items, item_thrown)
				//fmt.Println(item_thrown, monkey_receiver)
			}
		}
		//fmt.Println("After round", i+1, monkey_list)
	}
	// Solution one
	var max_inspects = 0
	var second_most_inspects = 0
	for _, v := range monkey_list {
		if v.inspect_count > max_inspects {
			second_most_inspects = max_inspects
			max_inspects = v.inspect_count
		} else if v.inspect_count > second_most_inspects {
			second_most_inspects = v.inspect_count
		}
	}
	fmt.Println("Solution One:", max_inspects*second_most_inspects)

}
