package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	big "github.com/ncw/gmp"
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
	items            []*big.Int
	operation_first  string
	operator         string
	operation_second string
	// Tests are always divisible by X
	test int
	// Pass/fail is always pass to X monkey
	test_pass     int
	test_fail     int
	inspect_count *big.Int
}

func stressRelief(input int) (output int) {
	output = int(math.Floor(float64(input) / 3))
	return output
}

func monkeyOperation(first *big.Int, second *big.Int, operator string) *big.Int {
	if operator == "+" {
		return first.Add(first, second)
	} else if operator == "*" {
		return first.Mul(first, second)
	}
	// Should only be one of the above
	fmt.Println("Error during monkey operation")
	os.Exit(0)
	return big.NewInt(0)
}

func oneMonkeyInspection(input Monkey) (output Monkey, inspectItem *big.Int, monkeyThrow int) {
	// Assume there is an item to inspect before this function is called
	inspectItem = input.items[0]
	output = input
	if len(input.items) == 1 {
		// This is the last item the monkey inspects
		output.items = make([]*big.Int, 0)
	} else { // Otherwise, pop it out
		output.items = output.items[1:]
	}
	first := big.NewInt(0)
	second := big.NewInt(0)
	if input.operation_first == "old" {
		first = inspectItem
	} else {
		// If not old, it's a number
		first.SetString(input.operation_first, 10)
	}
	if input.operation_second == "old" {
		second = inspectItem
	} else {
		// If not old, it's a number
		second.SetString(input.operation_second, 10)
	}
	// Worry level increased due to specific monkey and then relief!
	inspectItem = monkeyOperation(first, second, input.operator)
	// Got super slow running stuff, RIP this check
	// if inspectItem.Cmp(big.NewInt(0)) == -1 {
	// 	fmt.Println("Bad math???")
	// 	fmt.Println(first, second, input.operator, inspectItem)
	// 	os.Exit(0)
	// }
	// Comment below for solution 2
	// inspectItem = stressRelief(inspectItem)
	// Time to test it
	monkeyinspect := big.NewInt(0)
	monkeyinspect.Mod(inspectItem, big.NewInt(int64(input.test)))
	if monkeyinspect.BitLen() == 0 {
		monkeyThrow = input.test_pass
	} else {
		monkeyThrow = input.test_fail
	}
	output.inspect_count.Add(output.inspect_count, big.NewInt(1))
	return output, inspectItem, monkeyThrow
}

func reduceItemsThrown(input *big.Int, monkeyPrime *big.Int) *big.Int {
	if input.Cmp(monkeyPrime) == 1 {
		for i := 0; input.Cmp(monkeyPrime) == 1; i++ {
			input.Sub(input, monkeyPrime)
		}
	}
	// If input isn't greater than monkeyPrime, do nothing!
	return input
}

func main() {
	lines, _ := readLines("data.txt")
	//var current_monkey_count = 0
	var current_monkey_struct = Monkey{}
	var monkey_list []Monkey
	monkeyPrime := big.NewInt(1)
	for _, line := range lines {
		var line_array = strings.Split(line, " ")
		if strings.HasPrefix(line, "Monkey") {
			// Have to ignore the colon at the end to parse to int properly
			current_monkey_struct = Monkey{}
		} else if strings.Contains(line, "Starting") {
			for idx, x := range line_array {
				// Starting at index 4 are real stuff
				if idx < 4 {
					continue
				}
				var current_item, _ = strconv.ParseInt(strings.Trim(x, ","), 10, 0)
				current_monkey_struct.items = append(current_monkey_struct.items, big.NewInt(current_item))
				current_monkey_struct.inspect_count = big.NewInt(0)
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
			monkeyPrime.Mul(monkeyPrime, big.NewInt(int64(test_number)))
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
	//fmt.Println("Initial List:", monkey_list, len(monkey_list))
	// For Solution One
	//var rounds = 20
	// For Solution Two
	var rounds = 10000
	for i := 0; i < rounds; i++ {
		for k := range monkey_list {
			// Begin one round of inspections
			for range monkey_list[k].items {
				item_thrown := big.NewInt(0)
				var monkey_receiver = 0
				monkey_list[k], item_thrown, monkey_receiver = oneMonkeyInspection(monkey_list[k])
				item_thrown = reduceItemsThrown(item_thrown, monkeyPrime)
				monkey_list[monkey_receiver].items = append(monkey_list[monkey_receiver].items, item_thrown)
				//fmt.Println(item_thrown, monkey_receiver)
			}
		}
		if i%100 == 0 {
			fmt.Println("Hit", i)
		}
	}
	// Solution one
	// var max_inspects int = 0
	// var second_most_inspects int = 0
	// Solution Two crap
	max_inspects := big.NewInt(0)
	second_most_inspects := big.NewInt(0)
	for _, v := range monkey_list {
		if v.inspect_count.Cmp(max_inspects) == 1 {
			second_most_inspects = max_inspects
			max_inspects = v.inspect_count
		} else if v.inspect_count.Cmp(second_most_inspects) == 1 {
			second_most_inspects = v.inspect_count
		}
	}
	// Solution One
	// fmt.Println("Solution One:", max_inspects*second_most_inspects, max_inspects, second_most_inspects)
	// Solution Two
	fmt.Println("Solution Two:", big.NewInt(0).Mul(max_inspects, second_most_inspects), max_inspects, second_most_inspects)
}
