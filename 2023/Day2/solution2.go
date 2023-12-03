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

func splitInput(input string) (int, int) {
	trimGameString := strings.Split(input, ":")
	gameNumber, _ := strconv.Atoi(trimGameString[0][5:])

	splitBySets := strings.Split(trimGameString[1], ";")
	minimumCube := Cubes{0, 0, 0}
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
		if current_cube.red > minimumCube.red {
			minimumCube.red = current_cube.red
		}
		if current_cube.blue > minimumCube.blue {
			minimumCube.blue = current_cube.blue
		}
		if current_cube.green > minimumCube.green {
			minimumCube.green = current_cube.green
		}
	}
	power := minimumCube.red * minimumCube.blue * minimumCube.green
	return gameNumber, power
}

func main() {
	lines, _ := readLines("data.txt")
	solution := 0
	for _, line := range lines {
		println("Raw:", line)
		gameNumber, power := splitInput(line)
		println("Game number:", gameNumber, power)
		solution += power
	}
	println("Solution:", solution)
}
