package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
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

func getHeight(input uint8) int {
	if input < 96 {
		// Must be uppercase
		fmt.Println("Why you try to pass an uppercase letter?", input)
		os.Exit(1)
	}
	return int(input - 96)
}

func checkOneDirection(input int, possible_direction int) bool {
	if possible_direction <= input+1 {
		return true
	}
	return false
}

func chooseDirection(left bool, right bool, up bool, down bool) string {
	// Create array of valid directions
	var directions []string
	if left {
		directions = append(directions, "left")
	}
	if right {
		directions = append(directions, "right")
	}
	if up {
		directions = append(directions, "up")
	}
	if down {
		directions = append(directions, "down")
	}
	rand.Seed(time.Now().Unix())
	//fmt.Println(directions)
	randomIndex := rand.Intn(len(directions))
	return directions[randomIndex]
}

func oneAttempt(grid_input [][]int, start_position [2]int, end_position [2]int, max_steps int) int {
	var current_position = [2]int{0, 0}
	var number_of_steps = 0
	var stepped_on_grid = make([][]int, len(grid_input))
	for current_position != end_position {
		// Check for possible directions and do boundary detection
		var x = current_position[0]
		var y = current_position[1]
		var position_number = grid_input[y][x]
		var possible_left = false
		var possible_right = false
		var possible_up = false
		var possible_down = false
		// Boundary checks
		//fmt.Println(x, y, len(grid_input), len(grid_input[0]))
		if x != 0 {
			possible_left = checkOneDirection(position_number, grid_input[y][x-1])
		}
		if y != 0 {
			possible_up = checkOneDirection(position_number, grid_input[y-1][x])
		}
		if x < len(grid_input[0])-1 {
			possible_right = checkOneDirection(position_number, grid_input[y][x+1])
		}
		if y < len(grid_input)-1 {
			possible_down = checkOneDirection(position_number, grid_input[y+1][x])
		}
		number_of_steps += 1
		// Randomly pick one of the step options
		// No dead ends as you can always backtrack...
		var newDirection = chooseDirection(possible_left, possible_right, possible_up, possible_down)
		//fmt.Println(newDirection)
		if newDirection == "left" {
			current_position[0] = x - 1
		} else if newDirection == "right" {
			current_position[0] = x + 1
		} else if newDirection == "up" {
			current_position[1] = y - 1
		} else if newDirection == "down" {
			current_position[1] = y + 1
		}
		//fmt.Println(number_of_steps, "Started at", position_number, "Went to", current_position)
		if number_of_steps == max_steps {
			// Stop trying, took too many steps
			return 0
		}
	}
	return number_of_steps
}
func main() {
	lines, _ := readLines("data.txt")
	var height_grid = make([][]int, len(lines))
	var start_position [2]int
	var end_position [2]int
	for idx, line := range lines {
		//fmt.Println(line)
		height_grid[idx] = make([]int, len(line))
		for index, letter := range line {
			// 0 is start, 1-26 is lowercase letters, 27 is end
			if string(letter) == "S" {
				height_grid[idx][index] = 0
				start_position[0] = idx
				start_position[1] = index
			} else if string(letter) == "E" {
				height_grid[idx][index] = 27
				end_position[0] = idx
				end_position[1] = index
			} else {
				//fmt.Println(getHeight(uint8(letter)))
				height_grid[idx][index] = getHeight(uint8(letter))
			}
		}
	}
	fmt.Println("Start", start_position, "End", end_position)
	fmt.Println("Height Grid")
	for _, v := range height_grid {
		fmt.Println(v)
	}
	// Start with most steps being I went through every single grid item
	least_steps := len(height_grid) * len(height_grid[0])
	for i := 0; i < 100000; i++ {
		possible_count := oneAttempt(height_grid, start_position, end_position, least_steps)
		if possible_count != 0 && possible_count < least_steps {
			least_steps = possible_count
		}
	}
	fmt.Println(least_steps)
	// Solution One
	// Solution Two
}
