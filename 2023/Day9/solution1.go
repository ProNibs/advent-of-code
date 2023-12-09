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

func convertToIntArray(input []string) []int {
	output := make([]int, len(input))
	for i, item := range input {
		output[i], _ = strconv.Atoi(item)
	}
	return output
}

func createDiffArray(input []int) []int {
	output := make([]int, len(input)-1)
	// Remember that len(output) should be -1 of len(input)
	for i := 0; i < len(input)-1; i++ {
		// Almost made this more complicated by doing absolute values...
		raw_answer := input[i+1] - input[i]
		output[i] = raw_answer
	}
	return output
}

// Return true if all zero
func checkIfAllZero(input []int) bool {
	output := true
	for _, one_number := range input {
		if one_number != 0 {
			output = false
			break
		}
	}
	return output
}

func getLastArrayItem(input []int) int {
	return input[len(input)-1]
}

func main() {
	lines, _ := readLines("data.txt")
	solution := 0
	for l, line := range lines {
		// Make array that holds the other arrays
		var arrayOfArrays [][]int
		// Create initial array from input and make ints
		initialArray := convertToIntArray(strings.Split(line, " "))
		arrayOfArrays = append(arrayOfArrays, initialArray)
		// Create first diff array and stash into arrayOfArrays
		secondArray := createDiffArray(initialArray)
		arrayOfArrays = append(arrayOfArrays, secondArray)
		currentArray := secondArray
		// DiffArray is not all zeros, time to loop a bunch
		if !checkIfAllZero(currentArray) {
			for i := 0; !checkIfAllZero(currentArray); i++ {
				new_array := createDiffArray(currentArray)
				arrayOfArrays = append(arrayOfArrays, new_array)
				currentArray = new_array
			}
		}
		fmt.Println("Currently on loop:", l)
		for k, one_array := range arrayOfArrays {
			fmt.Println("Index:", k, "Array:", one_array, len(one_array))
		}
		// fmt.Println(arrayOfArrays)
		// Now reverse loop arrayOfArrays to get prediction right
		// Safely assume arrayOfArrays[-1] is the 0s, so not important
		for i := len(arrayOfArrays) - 2; i >= 0; i-- {
			fmt.Println(arrayOfArrays[i], i)
			predictedNumber := getLastArrayItem(arrayOfArrays[i]) + getLastArrayItem(arrayOfArrays[i+1])
			arrayOfArrays[i] = append(arrayOfArrays[i], predictedNumber)
		}
		//fmt.Println(arrayOfArrays)
		solution += getLastArrayItem(arrayOfArrays[0])
	}
	println("Solution:", int(solution))
}
