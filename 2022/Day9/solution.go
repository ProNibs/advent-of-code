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

func deduplicateSlice(input [][2]int) [][2]int {
	var new_slice [][2]int
	new_slice = append(new_slice, input[0])
	for _, one_item := range input {
		fmt.Println(one_item)
		is_new := true
		for _, k := range new_slice {
			if one_item == k {
				is_new = false
				break
			}
		}
		if is_new {
			new_slice = append(new_slice, one_item)
		}
	}
	return new_slice
}

func moveOnce(head [2]int, tail [2]int, direction string) ([2]int, [2]int) {
	if direction == "L" {
		head[0] -= 1
		//fmt.Println("Head now at", head)
		if tail[0]-head[0] > 1 {
			//fmt.Println("Need to move tail LEFT from", tail)
			tail[0] = head[0] + 1
			tail[1] = head[1]
			//fmt.Println("Tail moved to", tail, "to match", head)
		}
	} else if direction == "R" {
		head[0] += 1
		// fmt.Println("Head now at", head)
		if head[0]-tail[0] > 1 {
			// fmt.Println("Need to move tail RIGHT from", tail)
			tail[0] = head[0] - 1
			tail[1] = head[1]
			// fmt.Println("Tail moved to", tail, "to match", head)
		}
	} else if direction == "U" {
		head[1] += 1
		// fmt.Println("Head now at", head)
		if head[1]-tail[1] > 1 {
			// fmt.Println("Need to move tail UP from", tail)
			tail[1] = head[1] - 1
			tail[0] = head[0]
			// fmt.Println("Tail moved to", tail, "to match", head)
		}
	} else if direction == "D" {
		head[1] -= 1
		// fmt.Println("Head now at", head)
		if tail[1]-head[1] > 1 {
			// fmt.Println("Need to move tail DOWN from", tail)
			tail[1] = head[1] + 1
			tail[0] = head[0]
			// fmt.Println("Tail moved to", tail, "to match", head)
		}
	}
	return head, tail
}

func main() {
	lines, _ := readLines("data.txt")
	// Assume 0,0 is starting point
	head_position := [2]int{0, 0}
	tail_position := [2]int{0, 0}
	var tail_positions [][2]int
	tail_positions = append(tail_positions, tail_position)
	for _, line := range lines {
		var line_array = strings.Split(line, " ")
		direction := line_array[0]
		amount, _ := strconv.Atoi(line_array[1])
		for i := 0; i < amount; i++ {
			//fmt.Println(i)
			head_position, tail_position = moveOnce(head_position, tail_position, direction)
			tail_positions = append(tail_positions, tail_position)
		}
	}
	//fmt.Println(tail_positions)
	fmt.Println("Solution One:", len(deduplicateSlice(tail_positions)))

}
