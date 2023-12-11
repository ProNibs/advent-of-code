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

func calcWins(time int, win_distance int) int {
	possible_wins := 0
	// println("Distance to beat:", win_distance)
	for i := 0; i < time; i++ {
		// fmt.Println("Distance travels", i*(time-i))
		if win_distance < i*(time-i) {
			possible_wins += 1
		}
	}
	return possible_wins
}

func main() {
	lines, _ := readLines("data.txt")
	solution := 1
	var time_array []int
	var distance_array []int
	for i, line := range lines {
		println("Raw:", line)
		// Simple input, so we can make assumptions for once
		if i == 0 {
			// Time
			temp_array := strings.Split(line, " ")
			// fmt.Println(temp_array)
			for _, one_int := range temp_array {
				current_int, err := strconv.Atoi(one_int)
				if err == nil {
					time_array = append(time_array, current_int)
				}
			}
		} else if i == 1 {
			// Distance
			temp_array := strings.Split(line, " ")
			// fmt.Println(temp_array)
			for _, one_int := range temp_array {
				current_int, err := strconv.Atoi(one_int)
				if err == nil {
					distance_array = append(distance_array, current_int)
				}
			}
		}
	}
	fmt.Println(time_array)
	fmt.Println(distance_array)
	for i := range time_array {
		// println(calcWins(time_array[i], distance_array[i]))
		solution *= calcWins(time_array[i], distance_array[i])
	}
	println("Solution:", solution)
}
