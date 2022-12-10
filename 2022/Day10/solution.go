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

func main() {
	lines, _ := readLines("data.txt")
	var clock_cycle = 0
	var register_x = 1
	var clock_strength = make(map[int]int)
	for _, line := range lines {
		if line == "noop" {
			clock_cycle += 1
			if clock_cycle == 20 {
				clock_strength[20] = 20 * register_x
			}
			if clock_cycle == 60 {
				clock_strength[60] = 60 * register_x
			}
			if clock_cycle == 100 {
				clock_strength[100] = 100 * register_x
			}
			if clock_cycle == 140 {
				clock_strength[140] = 140 * register_x
			}
			if clock_cycle == 180 {
				clock_strength[180] = 180 * register_x
			}
			if clock_cycle == 220 {
				clock_strength[220] = 220 * register_x
			}

		} else {
			var addx, _ = strconv.Atoi(strings.Split(line, " ")[1])
			clock_cycle += 2
			if clock_cycle == 20 || clock_cycle == 21 {
				clock_strength[20] = 20 * register_x
			}
			if clock_cycle == 60 || clock_cycle == 61 {
				clock_strength[60] = 60 * register_x
			}
			if clock_cycle == 100 || clock_cycle == 101 {
				clock_strength[100] = 100 * register_x
			}
			if clock_cycle == 140 || clock_cycle == 141 {
				clock_strength[140] = 140 * register_x
			}
			if clock_cycle == 180 || clock_cycle == 181 {
				clock_strength[180] = 180 * register_x
			}
			if clock_cycle == 220 || clock_cycle == 221 {
				clock_strength[220] = 220 * register_x
			}
			register_x += addx
		}
	}
	fmt.Println(clock_cycle, register_x, clock_strength)
	var strength_sum = 0
	for _, v := range clock_strength {
		strength_sum += v
	}
	fmt.Println("Solution One:", strength_sum)

}
