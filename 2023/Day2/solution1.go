package main

import (
	"bufio"
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

type Cubes struct {
	red   int
	blue  int
	green int
}

// Return false is exceeds limits
func checkOneSet(input Cubes) bool {
	if input.red > 12 {
		return false
	} else if input.blue > 14 {
		return false
	} else if input.green > 13 {
		return false
	} else {
		return true
	}
}

func splitInput(input string) (int, bool) {
	trimGameString := strings.Split(input, ":")
	gameNumber, _ := strconv.Atoi(trimGameString[0][5:])

	invalid_game := false
	splitBySets := strings.Split(trimGameString[1], ";")
	for _, sets := range splitBySets {
		current_cube := Cubes{0, 0, 0}
		for _, one_color := range strings.Split(sets, ",") {
			number_color := strings.Split(one_color, " ")
			if number_color[2] == "red" {
				current_cube.red, _ = strconv.Atoi(number_color[1])
			} else if number_color[2] == "blue" {
				current_cube.blue, _ = strconv.Atoi(number_color[1])
			} else if number_color[2] == "green" {
				current_cube.green, _ = strconv.Atoi(number_color[1])
			}
		}
		if !checkOneSet(current_cube) {
			invalid_game = true
			return gameNumber, invalid_game
		}
	}
	return gameNumber, invalid_game
}

func main() {
	lines, _ := readLines("data.txt")
	solution_one := 0
	for _, line := range lines {
		println("Raw:", line)
		gameNumber, invalid_game := splitInput(line)
		println("Game number:", gameNumber, invalid_game)
		if !invalid_game {
			solution_one += gameNumber
		}
	}
	println("Solution:", solution_one)
}
