package main

import "fmt"

func walkLeft(input [][]string, start_location []int) [][]int {
	var output [][]int
	start_location_value := input[start_location[0]][start_location[1]]
	for i := start_location[1] - 1; i > -1; i-- {
		if input[start_location[0]][i] == start_location_value {
			output = append(output, []int{start_location[0], i})
		} else {
			// We broke the walk, get out
			break
		}
	}
	return output
}

func walkRight(input [][]string, start_location []int) [][]int {
	var output [][]int
	start_location_value := input[start_location[0]][start_location[1]]
	for i := start_location[1] + 1; i < len(input[0]); i++ {
		if input[start_location[0]][i] == start_location_value {
			output = append(output, []int{start_location[0], i})
		} else {
			// We broke the walk, get out
			break
		}
	}
	return output
}

func walkUp(input [][]string, start_location []int) [][]int {
	var output [][]int
	start_location_value := input[start_location[0]][start_location[1]]
	for j := start_location[0] - 1; j > -1; j-- {
		if input[j][start_location[1]] == start_location_value {
			output = append(output, []int{j, start_location[1]})
		} else {
			// We broke the walk, get out
			break
		}
	}
	return output
}

func walkDown(input [][]string, start_location []int) [][]int {
	var output [][]int
	start_location_value := input[start_location[0]][start_location[1]]
	for j := start_location[0] + 1; j < len(input); j++ {
		if input[j][start_location[1]] == start_location_value {
			output = append(output, []int{j, start_location[1]})
		} else {
			// We broke the walk, get out
			break
		}
	}
	return output
}

func calcNeighbors(input [][]string, start_location []int) [][]int {
	var output [][]int
	// start_location_value := input[start_location[0]][start_location[1]]
	// Obviously, starting location is a match
	output = append(output, start_location)
	// Walk left to see if we can find matches
	output = append(output, walkLeft(input, start_location)...)
	// Walk right to see if we can find matches
	output = append(output, walkRight(input, start_location)...)
	// Walk up to see if we can find matches
	output = append(output, walkUp(input, start_location)...)
	// Walk down to see if we can find matches
	output = append(output, walkDown(input, start_location)...)

	fmt.Println("Basic left,right,up,down", output)
	// To see if we can get even more neighbors
	// Loop over our current matches via changing the start_locations
	last_output_length := len(output)
	for _, new_start_location := range output {
		var new_start_output [][]int
		new_start_output = append(new_start_output, new_start_location)
		// Walk left to see if we can find matches
		new_start_output = append(new_start_output, walkLeft(input, new_start_location)...)
		// Walk right to see if we can find matches
		new_start_output = append(new_start_output, walkRight(input, new_start_location)...)
		// Walk up to see if we can find matches
		new_start_output = append(new_start_output, walkUp(input, new_start_location)...)
		// Walk down to see if we can find matches
		new_start_output = append(new_start_output, walkDown(input, new_start_location)...)
		// fmt.Println("With new start coordinate:", new_start_location)
		fmt.Println(new_start_output, last_output_length, len(new_start_output))
		// Check if any of the new stuff is in the old output
		for _, one_pair := range new_start_output {
			skipped := false
			for _, other_pair := range output {
				fmt.Println("New one", one_pair, "Compare to", other_pair)
				// It's already present, let's skip
				if one_pair[0] == other_pair[0] && one_pair[1] == other_pair[1] {
					skipped = true
					break

				}
			}
			// Never found a pair, so it must be new
			if !skipped {
				output = append(output, []int{one_pair[0], one_pair[1]})
			}
		}
	}
	// fmt.Println("Crazy loops", output)
	return output
}
