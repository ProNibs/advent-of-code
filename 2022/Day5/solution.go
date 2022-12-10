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

func reverseArray(inputArray []string) []string {
	var output_array []string
	for i := len(inputArray) - 1; i != -1; i = i - 1 {
		output_array = append(output_array, inputArray[i])
	}
	return output_array
}

func moveStacks(movement int, source []string, target []string) ([]string, []string) {
	for i := 0; i < movement; i++ {
		var pop_value = source[len(source)-1]
		source = source[:len(source)-1]
		target = append(target, pop_value)
	}
	return source, target
}

func moveMoreStacks(movement int, source []string, target []string) ([]string, []string) {
	// Move multiple at once, actually easier?
	var pop_value = source[len(source)-movement:]
	//fmt.Println("Pop?", pop_value)
	source = source[:len(source)-movement]
	target = append(target, pop_value...)
	return source, target
}

func main() {
	lines, _ := readLines("data.txt")
	//var solution_one = 0
	//var solution_two = 0
	var stacks [][]string
	var movements [][3]int
	var movement_flag = false
	for _, line := range lines {
		// Check for the empty line
		// Once we hit it, we in the movement text area now
		if len(line) == 0 {
			movement_flag = true
			continue
		}
		if movement_flag {
			// Assume everything is "move X from Y to Z"
			var string_array = strings.Split(line, " ")
			var first, _ = strconv.Atoi(string_array[1])
			var second, _ = strconv.Atoi(string_array[3])
			var third, _ = strconv.Atoi(string_array[5])
			// Alter stack numbers down by 1 since we index by 0
			int_array := [3]int{first, second - 1, third - 1}
			movements = append(movements, int_array)
		} else {
			// Make assumptions because screw up/down parsing
			// Noticed it's [X] (with a space) so let's go by 4
			var number_of_stacks = (len(line) + 1) / 4
			// Make the top level array the right length if not already
			if len(stacks) != number_of_stacks {
				for i := 0; i < number_of_stacks; i++ {
					stacks = append(stacks, []string{})
				}
			}
			// Check to see if it's just the stack numbers
			if strings.Contains(line, "1") {
				continue
			}
			for stack := 0; stack < number_of_stacks; stack++ {
				var possible_crate = line[(stack * 4) : (stack*4)+2]
				// Verify a crate is there before dropping it in
				if strings.Contains(possible_crate, "[") {
					// Prune the [] now that I don't need them
					stacks[stack] = append(stacks[stack], possible_crate[1:2])
				}
			}
		}

	}
	// At the end here, all the stacks are currently reverse order
	for i, stack := range stacks {
		stacks[i] = reverseArray(stack)
	}

	// Deep copy for solution two to do its changes
	solution_two_stacks := make([][]string, len(stacks))
	copy(solution_two_stacks, stacks)

	// Comment out -- seems golang still keeps refs in array of arrays crap
	// for i, _ := range movements {
	// 	var source_stack = movements[i][1]
	// 	var target_stack = movements[i][2]
	// 	stacks[source_stack], stacks[target_stack] = moveStacks(movements[i][0], stacks[source_stack], stacks[target_stack])
	// }

	var solution_one = ""
	var solution_two = ""
	for _, stack := range stacks {
		solution_one += stack[len(stack)-1]
	}
	fmt.Println(solution_one)

	for i, _ := range movements {
		var source_stack = movements[i][1]
		var target_stack = movements[i][2]
		solution_two_stacks[source_stack], solution_two_stacks[target_stack] = moveMoreStacks(movements[i][0], solution_two_stacks[source_stack], solution_two_stacks[target_stack])
	}
	//fmt.Println(solution_two_stacks)
	for _, stack := range solution_two_stacks {
		solution_two += stack[len(stack)-1]
	}
	fmt.Println(solution_two)
}
