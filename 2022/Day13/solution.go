package main

import (
	"bufio"
	"fmt"
	"os"
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

// Need to have dynamic nesting of arrays for this to work
type Pairs struct {
	pairOne
}

type NestedArrays struct {
}

func main() {
	lines, _ := readLines("data.txt")
	// Pairs are defined by their index
	var master_list [2][][]string
	master_list[0] = []string{"hello"}
	master_list[1] = append(master_list[1], []string{"world"})
	//master_list[1] = []string{"hello"}
	fmt.Println(master_list)
	current_list_index := 0
	for _, line := range lines {
		// Line break, on to next pair
		if !strings.HasPrefix(line, "[") {
			fmt.Println("Next pair")
			current_list_index += 1
		} else {
			// Figure out what this array looks like
			// Remove [ and ] ends to clean up logic
			// Let's try split via comma first and go from there
			split := strings.Split(line[1:len(line)-1], ",")
			// fmt.Println(split)
			var output_array []string
			for _, item := range split {
				// fmt.Println(item)
				output_array = append(output_array, item)
			}
			// Dumbly stuck everything into array, check if it's valid
			// Is an array present?
			array_present := false
			for _, item := range output_array {
				if strings.Contains(item, "[") || strings.Contains(item, "[") {
					fmt.Println("Still an embedded array here")
					array_present = true
					break
				}
			}
			if !array_present {
				fmt.Println(output_array)
			}
			// master_list[current_pair_index] = append(output_array
			// current_pair_index += 1
		}
	}
	fmt.Println("Master List:", master_list)
}
