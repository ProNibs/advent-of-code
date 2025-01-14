package main

import "testing"

func TestAreaForTestdata1ForRegionA(t *testing.T) {
	// Specifically Region "A"
	input_seed := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	}

	area := calcArea(input_seed)
	if area != 4 {
		t.Fatalf(`"A" Area didn't return 4. It returned %d`, area)
	}
}

func TestAreaForTestdata1ForRegionB(t *testing.T) {
	input_seed := [][]int{
		{1, 0},
		{1, 1},
		{2, 0},
		{2, 1},
	}

	area := calcArea(input_seed)
	if area != 4 {
		t.Fatalf(`"B" Area didn't return 4. It returned %d`, area)
	}
}

func TestAreaForTestdata1ForRegionC(t *testing.T) {
	input_seed := [][]int{
		{1, 2},
		{2, 2},
		{2, 3},
		{3, 3},
	}

	area := calcArea(input_seed)
	if area != 4 {
		t.Fatalf(`"C" Area didn't return 4. It returned %d`, area)
	}
}

func TestAreaForTestdata1ForRegionD(t *testing.T) {
	input_seed := [][]int{
		{1, 3},
	}

	area := calcArea(input_seed)
	if area != 1 {
		t.Fatalf(`"D" Area didn't return 1. It returned %d`, area)
	}
}

func TestAreaForTestdata1ForRegionE(t *testing.T) {
	input_seed := [][]int{
		{3, 0},
		{3, 1},
		{3, 2},
	}

	area := calcArea(input_seed)
	if area != 3 {
		t.Fatalf(`"E" Area didn't return 3. It returned %d`, area)
	}
}
