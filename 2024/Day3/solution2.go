package main

import (
	"bufio"
	"os"
	"regexp"
	"sort"
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

func multiplyStuff(input string) int {
	// Assume we just get a mul(X,X) string
	string_split := strings.Split(input, ",")
	re := regexp.MustCompile(`\d+`)
	left_number, _ := strconv.Atoi(re.FindStringSubmatch(string_split[0])[0])
	right_number, _ := strconv.Atoi(re.FindStringSubmatch(string_split[1])[0])
	// fmt.Println(left_number, right_number)
	return left_number * right_number
}

func findDeadZones(enabled_array [][]int, disabled_array [][]int) [][]int {
	// fmt.Println(enabled_array)
	// fmt.Println(disabled_array)
	var starting_enabled_array []int
	var starting_disabled_array []int
	// Only need the left side of each array
	for _, small_array := range enabled_array {
		starting_enabled_array = append(starting_enabled_array, small_array[0])
	}
	for _, small_array := range disabled_array {
		starting_disabled_array = append(starting_disabled_array, small_array[0])
	}
	// Sort just to be sure
	sort.Ints(starting_enabled_array)
	sort.Ints(starting_disabled_array)
	// fmt.Println("Enable:", starting_enabled_array, "Disabled:", starting_disabled_array)
	// Hopefully each line starts enabled, otherwise yeah...
	// On Second though, puzzle input isn't too large, can just make it a big one liner
	// Loop through Disabled array, find next viable Enable location and append to end result
	var dead_zone_array [][]int
	for _, start_of_dead_zone := range starting_disabled_array {
		for _, potential_end_of_dead_zone := range starting_enabled_array {
			if potential_end_of_dead_zone > start_of_dead_zone {
				dead_zone_array = append(dead_zone_array, []int{start_of_dead_zone, potential_end_of_dead_zone})
				break
			}
		}
	}
	return dead_zone_array
}

func checkIfNotInDeadZone(dead_zone_array [][]int, check_index int) bool {
	for _, one_dead_zone := range dead_zone_array {
		if check_index > one_dead_zone[0] && check_index < one_dead_zone[1] {
			return false
		}
	}
	return true
}

func main() {
	lines, _ := readLines("data.txt")
	solution := 0
	regex_matcher := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	enabled_matches := regexp.MustCompile(`do\(\)`)
	disabled_matches := regexp.MustCompile(`don't\(\)`)
	for _, line := range lines {
		// println("Raw:", line)
		// mul_array := regex_matcher.FindAllStringSubmatch(line, -1)
		mul_array_index := regex_matcher.FindAllStringSubmatchIndex(line, -1)
		// fmt.Println(mul_array)
		// fmt.Println("Array indexes:", mul_array_index)
		// fmt.Println(enabled_matches.FindAllStringSubmatchIndex(line, -1))
		// fmt.Println(disabled_matches.FindAllStringSubmatchIndex(line, -1))
		dead_zone_array := findDeadZones(enabled_matches.FindAllStringSubmatchIndex(line, -1), disabled_matches.FindAllStringSubmatchIndex(line, -1))
		for i, one_string := range regex_matcher.FindAllStringSubmatch(line, -1) {
			// fmt.Println(one_string[0], "has left index", mul_array_index[i][0])
			// fmt.Println("Dead zones:", dead_zone_array)
			// fmt.Println(checkIfNotInDeadZone(dead_zone_array, mul_array_index[i][0]))
			if checkIfNotInDeadZone(dead_zone_array, mul_array_index[i][0]) {
				solution += multiplyStuff(one_string[0])
			}
		}
	}
	println("Solution:", solution)
}
