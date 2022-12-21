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

func deduplicateSlice(input [][2]int) [][2]int {
	var new_slice [][2]int
	new_slice = append(new_slice, input[0])
	for _, one_item := range input {
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

func moveHead(head [2]int, direction string) [2]int {
	if direction == "L" {
		head[0] -= 1
	} else if direction == "R" {
		head[0] += 1
	} else if direction == "U" {
		head[1] += 1
	} else if direction == "D" {
		head[1] -= 1
	}
	return head
}

func calculateFollowing(head [2]int, tail [2]int) [2]int {
	// Need to account for complete diagonals
	if math.Abs(float64(head[0]-tail[0])) > 1 && math.Abs(float64(head[1]-tail[1])) > 1 {
		// fmt.Println("Hit pure diagonal.")
		// fmt.Println("Comparing:", head, tail)
		// Figure out which way to go diagonally
		// Top right
		if head[0] > tail[0] && head[1] > tail[1] {
			tail[0] = head[0] - 1
			tail[1] = head[1] - 1
			//Bottom right
		} else if head[0] > tail[0] && head[1] < tail[1] {
			tail[0] = head[0] - 1
			tail[1] = head[1] + 1
			// Bottom left
		} else if head[0] < tail[0] && head[1] < tail[1] {
			tail[0] = head[0] + 1
			tail[1] = head[1] + 1
			// Top left
		} else if head[0] < tail[0] && head[1] > tail[1] {
			tail[0] = head[0] + 1
			tail[1] = head[1] - 1
		}
		return tail
	}
	// Back to normal off-by-only-one statements
	if tail[0]-head[0] > 1 && tail[0] > head[0] {
		//fmt.Println("Need to move tail LEFT from", tail)
		tail[0] = head[0] + 1
		tail[1] = head[1]
		//fmt.Println("Tail moved to", tail, "to match", head)
	} else if head[0]-tail[0] > 1 && head[0] > tail[0] {
		// fmt.Println("Need to move tail RIGHT from", tail)
		tail[0] = head[0] - 1
		tail[1] = head[1]
		// fmt.Println("Tail moved to", tail, "to match", head)
	} else if head[1]-tail[1] > 1 && head[1] > tail[1] {
		// fmt.Println("Need to move tail UP from", tail)
		tail[1] = head[1] - 1
		tail[0] = head[0]
		// fmt.Println("Tail moved to", tail, "to match", head)
	} else if tail[1]-head[1] > 1 && tail[1] > head[1] {
		// fmt.Println("Need to move tail DOWN from", tail)
		tail[1] = head[1] + 1
		tail[0] = head[0]
		// fmt.Println("Tail moved to", tail, "to match", head)
	}
	return tail
}

func main() {
	lines, _ := readLines("data.txt")
	// Assume 0,0 is starting point
	head_position := [2]int{0, 0}
	one_position := [2]int{0, 0}
	two_position := [2]int{0, 0}
	three_position := [2]int{0, 0}
	four_position := [2]int{0, 0}
	five_position := [2]int{0, 0}
	six_position := [2]int{0, 0}
	seven_position := [2]int{0, 0}
	eight_position := [2]int{0, 0}
	tail_position := [2]int{0, 0}
	var tail_positions [][2]int
	tail_positions = append(tail_positions, tail_position)
	for _, line := range lines {
		var line_array = strings.Split(line, " ")
		direction := line_array[0]
		amount, _ := strconv.Atoi(line_array[1])
		for i := 0; i < amount; i++ {
			// fmt.Println("Moving:", i)
			// For Sol One
			head_position = moveHead(head_position, direction)
			// Added for part 2
			one_position = calculateFollowing(head_position, one_position)
			two_position = calculateFollowing(one_position, two_position)
			three_position = calculateFollowing(two_position, three_position)
			four_position = calculateFollowing(three_position, four_position)
			five_position = calculateFollowing(four_position, five_position)
			six_position = calculateFollowing(five_position, six_position)
			seven_position = calculateFollowing(six_position, seven_position)
			eight_position = calculateFollowing(seven_position, eight_position)
			// End sol2 wackiness
			tail_position = calculateFollowing(eight_position, tail_position)
			tail_positions = append(tail_positions, tail_position)
			// fmt.Println("Positions", head_position,
			// 	one_position, two_position, three_position, four_position, five_position,
			// 	six_position, seven_position, eight_position, tail_position)
		}
		// fmt.Println("Positions", head_position,
		// 	one_position, two_position, three_position, four_position, five_position,
		// 	six_position, seven_position, eight_position, tail_position)
	}
	//fmt.Println(tail_positions)
	fmt.Println("Solution:", len(deduplicateSlice(tail_positions)))
}
