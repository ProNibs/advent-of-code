package main

func calcPerimeter(input [][]int) int {
	// Assume only values of a similar type were given and adjacent already
	output := 0
	for i, coordinates := range input {
		// It's easier to assume it's standalone and then subtract if there is a neighbor
		output += 4
		for j, neighbors := range input {
			if i == j {
				continue
			}
			// Check for neighbors to the left and/or right
			// Check if on same row first
			if coordinates[1] == neighbors[1] {
				// Neighbor to the right
				if coordinates[0]+1 == neighbors[0] {
					output -= 1
				}
				// Neighbor to the left
				if coordinates[0]-1 == neighbors[0] {
					output -= 1
				}
			}
			// Check for neighbors to the up and/or down
			// Check if same column first
			if coordinates[0] == neighbors[0] {
				// Neighbor up
				if coordinates[1]-1 == neighbors[1] {
					output -= 1
				}
				// Neighbor down
				if coordinates[1]+1 == neighbors[1] {
					output -= 1
				}
			}
		}
		// println(i)
	}
	return output
}
