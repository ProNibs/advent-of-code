package main

import (
	"bufio"
	"fmt"
	"os"
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

type Rock struct {
	width  int
	height int
	rocks  [][]int
}

func makeHorizontalRock() *Rock {
	rock := new(Rock)
	rock.width = 4
	rock.height = 1
	rock.rocks = [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	return rock
}

func makePlusRock() *Rock {
	rock := new(Rock)
	rock.width = 3
	rock.height = 3
	rock.rocks = [][]int{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}
	return rock
}

func makeLShapedRock() *Rock {
	rock := new(Rock)
	rock.width = 3
	rock.height = 3
	rock.rocks = [][]int{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}
	return rock
}

func makeVerticalRock() *Rock {
	rock := new(Rock)
	rock.width = 1
	rock.height = 4
	rock.rocks = [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	return rock
}

func makeSquareRock() *Rock {
	rock := new(Rock)
	rock.width = 2
	rock.height = 2
	rock.rocks = [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	return rock
}

// Convert the rock to rock chamber coordinates
// Idea here is to then write collision detection as it falls
func convertRock(inputRock *Rock, top_rock int) [][]int {
	var output = inputRock.rocks
	for i, _ := range output {
		// Adjust 2 from left side
		output[i][0] += 2
		// Adjust 3 from the heighest rock
		output[i][1] += 3 + top_rock
	}
	// fmt.Println(output)
	return output
}

func pushRock(input [][]int, direction string, rock_chamber [][]int) [][]int {
	if direction == ">" {
		// Check to ensure it CAN be pushed right
		for _, item := range input {
			// Cannot be pushed right
			if item[0]+1 > 6 {
				return input
			}
		}
		for idx := range input {
			// Passed earlier check, push it right
			input[idx][0] += 1
		}
		if rockCollideCheck(input, rock_chamber) {
			// Actually, that'll hit another rock...
			fmt.Println("Reverting a push right")
			for idx := range input {
				input[idx][0] -= 1
			}
		}
	} else if direction == "<" {
		// Check to ensure it CAN be pushed left
		for _, item := range input {
			// Cannot be pushed left
			if item[0]-1 < 0 {
				return input
			}
		}
		for idx := range input {
			// Passed earlier check, push it left
			input[idx][0] -= 1
		}
		if rockCollideCheck(input, rock_chamber) {
			// Actually, that'll hit another rock...
			fmt.Println("Reverting a push Left")
			for idx := range input {
				input[idx][0] += 1
			}
		}
	}
	// fmt.Println("Push rock equals:", input)
	return input
}

func deepCopy(input [][]int) [][]int {
	output := make([][]int, len(input))
	for idx := range output {
		output[idx] = make([]int, len(input[idx]))
		copy(output[idx], input[idx])
	}
	return output
}

func attemptRockFall(input [][]int) [][]int {
	output := deepCopy(input)
	for idx := range output {
		output[idx][1] -= 1
	}
	return output
}

func rockCollideCheck(input [][]int, rock_chamber [][]int) bool {
	collision_detected := false
	for _, rock := range input {
		for _, item := range rock_chamber {
			if (item[0] == rock[0]) && (item[1] == rock[1]) {
				collision_detected = true
				break
			}
		}
		if collision_detected {
			break
		}
	}
	return collision_detected
}

func checkRockCollision(input [][]int, highest_rock int, rock_chamber [][]int) bool {
	output := deepCopy(input)
	collision_detected := false
	// fmt.Println("Checking input", output, highest_rock)
	for _, coordinates := range output {
		// fmt.Println("Checking coordinates:", coordinates)
		// Check for easier one, already on the floor
		if coordinates[1] < 0 || collision_detected {
			collision_detected = true
			break
		}
		// Possible collision, let's check against rock chamber
		if coordinates[1] <= highest_rock+3 {
			// fmt.Println("Attempting collision detect")
			for _, item := range rock_chamber {
				if (item[0] == coordinates[0]) && (item[1] == coordinates[1]) {
					collision_detected = true
					break
				}
			}
		}
	}
	// fmt.Println(collision_detected)
	return collision_detected
}

func nextRock(rock_rotation int, rock_list []string) *Rock {
	// Error checking first
	if rock_rotation >= len(rock_list) {
		fmt.Println("Tried to make rock", rock_rotation)
		os.Exit(1)
	}
	fmt.Println("Making rock", rock_list[rock_rotation])
	if rock_list[rock_rotation] == "horizontal" {
		return makeHorizontalRock()
	} else if rock_list[rock_rotation] == "plus" {
		return makePlusRock()
	} else if rock_list[rock_rotation] == "l" {
		return makeLShapedRock()
	} else if rock_list[rock_rotation] == "vertical" {
		return makeVerticalRock()
	} else if rock_list[rock_rotation] == "square" {
		return makeHorizontalRock()
	}
	fmt.Println("Didn't match rock list?")
	os.Exit(1)
	return makeHorizontalRock()
}

func calculateHighestRock(rock_chamber [][]int) int {
	highest_rock := 0
	for _, rock := range rock_chamber {
		// Y-coordinates start from 0...
		if rock[1]+1 > highest_rock {
			highest_rock = rock[1] + 1
		}
	}
	return highest_rock
}

func main() {
	lines, _ := readLines("data.txt")
	var rock_chamber [][]int
	highest_rock := 0
	current_rock := convertRock(makeHorizontalRock(), highest_rock)
	rock_rotation := 0
	rock_rotation_list := []string{"horizontal", "plus", "l", "vertical", "square"}
	// fmt.Println("Starting rock:", current_rock)
	rocks_dropped := 1
	max_rocks_dropped := 6
	for i := 0; rocks_dropped < max_rocks_dropped; i++ {
		for _, direction := range lines[0] {
			fmt.Println(string(direction))
			// fmt.Println("Current rock before push:", current_rock)
			current_rock = pushRock(current_rock, string(direction), rock_chamber)
			// fmt.Println("Post push", current_rock)
			possible_rock := attemptRockFall(current_rock)
			// fmt.Println("Possible rock after push and fall:", possible_rock)
			// fmt.Println("Chamber:", rock_chamber)
			if !checkRockCollision(possible_rock, highest_rock, rock_chamber) {
				// fmt.Println("No collision, set it!", current_rock, possible_rock)
				current_rock = possible_rock
			} else {
				// Collision would happen, so time to set it in stone
				for _, coordinates := range current_rock {
					rock_chamber = append(rock_chamber, coordinates)
				}
				highest_rock = calculateHighestRock(rock_chamber)
				// Re-start from horizontal rock
				rock_rotation += 1
				if rock_rotation >= len(rock_rotation_list) {
					rock_rotation = 0
				}
				current_rock = convertRock(nextRock(rock_rotation, rock_rotation_list), highest_rock)
				fmt.Println("Starting a rock at", current_rock)
				rocks_dropped += 1
				if rocks_dropped > max_rocks_dropped {
					break
				}
			}
		}
		if rocks_dropped >= max_rocks_dropped {
			break
		}
	}
	fmt.Println(rock_chamber)
	fmt.Println("Solution One:", highest_rock)
}
