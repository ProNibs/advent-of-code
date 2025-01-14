package main

import "testing"

func TestCalcNeighborsLeftToRight(t *testing.T) {
	input_seed := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "C", "D"},
		{"B", "B", "C", "C"},
		{"E", "E", "E", "C"},
	}
	// Specifically Region "A"
	expected_output := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	}

	neighbors := calcNeighbors(input_seed, []int{0, 0})
	if len(expected_output) != len(neighbors) {
		t.Fatalf(`Amount of neighbors detected does not match amount expected.`)
	}
	for _, one_coordinate := range expected_output {
		matched := false
		for _, other_coordinate := range neighbors {
			if other_coordinate[0] == one_coordinate[0] && other_coordinate[1] == one_coordinate[1] {
				matched = true
				break
			}
		}
		if !matched {
			t.Fatalf(`This coordinate didn't receive a match: (%d,%d)`, one_coordinate[0], one_coordinate[1])
		}
	}
}

func TestCalcNeighborsRightToLeft(t *testing.T) {
	input_seed := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "C", "D"},
		{"B", "B", "C", "C"},
		{"E", "E", "E", "C"},
	}
	// Specifically Region "A"
	expected_output := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	}

	neighbors := calcNeighbors(input_seed, []int{0, 3})
	if len(expected_output) != len(neighbors) {
		t.Fatalf(`Amount of neighbors detected does not match amount expected.`)
	}
	for _, one_coordinate := range expected_output {
		matched := false
		for _, other_coordinate := range neighbors {
			if other_coordinate[0] == one_coordinate[0] && other_coordinate[1] == one_coordinate[1] {
				matched = true
				break
			}
		}
		if !matched {
			t.Fatalf(`This coordinate didn't receive a match: (%d,%d)`, one_coordinate[0], one_coordinate[1])
		}
	}
}

func TestCalcNeighborsDown(t *testing.T) {
	input_seed := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "B", "D"},
		{"B", "B", "B", "C"},
		{"E", "E", "E", "C"},
	}
	// Specifically a modified Region "C"
	expected_output := [][]int{
		{2, 3},
		{3, 3},
	}

	neighbors := calcNeighbors(input_seed, []int{2, 3})
	if len(expected_output) != len(neighbors) {
		t.Fatalf(`Amount of neighbors detected does not match amount expected.`)
	}
	for _, one_coordinate := range expected_output {
		matched := false
		for _, other_coordinate := range neighbors {
			if other_coordinate[0] == one_coordinate[0] && other_coordinate[1] == one_coordinate[1] {
				matched = true
				break
			}
		}
		if !matched {
			t.Fatalf(`This coordinate didn't receive a match: (%d,%d)`, one_coordinate[0], one_coordinate[1])
		}
	}
}

func TestCalcNeighborsUp(t *testing.T) {
	input_seed := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "B", "D"},
		{"B", "B", "B", "C"},
		{"E", "E", "E", "C"},
	}
	// Specifically a modified Region "C"
	expected_output := [][]int{
		{2, 3},
		{3, 3},
	}

	neighbors := calcNeighbors(input_seed, []int{3, 3})
	if len(expected_output) != len(neighbors) {
		t.Fatalf(`Amount of neighbors detected does not match amount expected.`)
	}
	for _, one_coordinate := range expected_output {
		matched := false
		for _, other_coordinate := range neighbors {
			if other_coordinate[0] == one_coordinate[0] && other_coordinate[1] == one_coordinate[1] {
				matched = true
				break
			}
		}
		if !matched {
			t.Fatalf(`This coordinate didn't receive a match: (%d,%d)`, one_coordinate[0], one_coordinate[1])
		}
	}
}

func TestCalcNeighborsSquare(t *testing.T) {
	input_seed := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "C", "D"},
		{"B", "B", "C", "C"},
		{"E", "E", "E", "C"},
	}
	// Specifically Region "B"
	expected_output := [][]int{
		{1, 0},
		{1, 1},
		{2, 0},
		{2, 1},
	}

	neighbors := calcNeighbors(input_seed, []int{1, 0})
	for _, one_coordinate := range expected_output {
		matched := false
		for _, other_coordinate := range neighbors {
			if other_coordinate[0] == one_coordinate[0] && other_coordinate[1] == one_coordinate[1] {
				matched = true
				break
			}
		}
		if !matched {
			t.Fatalf(`This coordinate didn't receive a match: (%d,%d)`, one_coordinate[0], one_coordinate[1])
		}
	}
	if len(expected_output) != len(neighbors) {
		t.Fatalf(`Amount of neighbors detected does not match amount expected.`)
	}
}

func TestCalcNeighborsRegionC(t *testing.T) {
	input_seed := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "C", "D"},
		{"B", "B", "C", "C"},
		{"E", "E", "E", "C"},
	}
	// Specifically Region "C"
	expected_output := [][]int{
		{1, 2},
		{2, 2},
		{2, 3},
		{3, 3},
	}

	neighbors := calcNeighbors(input_seed, []int{1, 2})
	for _, one_coordinate := range expected_output {
		matched := false
		for _, other_coordinate := range neighbors {
			if other_coordinate[0] == one_coordinate[0] && other_coordinate[1] == one_coordinate[1] {
				matched = true
				break
			}
		}
		if !matched {
			t.Fatalf(`This coordinate didn't receive a match: (%d,%d)`, one_coordinate[0], one_coordinate[1])
		}
	}
	if len(expected_output) != len(neighbors) {
		t.Fatalf(`Amount of neighbors detected does not match amount expected.`)
	}
}
