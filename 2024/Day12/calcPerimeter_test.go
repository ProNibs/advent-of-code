package main

import "testing"

func TestPerimeterForTestdata1ForRegionA(t *testing.T) {
	// Specifically Region "A"
	input_seed := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	}

	perimeter := calcPerimeter(input_seed)
	if perimeter != 10 {
		t.Fatalf(`"A" Perimeter didn't return 10. It returned %d`, perimeter)
	}
}

func TestPerimeterForTestdata1ForRegionB(t *testing.T) {
	input_seed := [][]int{
		{1, 0},
		{1, 1},
		{2, 0},
		{2, 1},
	}

	perimeter := calcPerimeter(input_seed)
	if perimeter != 8 {
		t.Fatalf(`"B" Perimeter didn't return 8. It returned %d`, perimeter)
	}
}

func TestPerimeterForTestdata1ForRegionC(t *testing.T) {
	input_seed := [][]int{
		{1, 2},
		{2, 2},
		{2, 3},
		{3, 3},
	}

	perimeter := calcPerimeter(input_seed)
	if perimeter != 10 {
		t.Fatalf(`"C" Perimeter didn't return 10. It returned %d`, perimeter)
	}
}

func TestPerimeterForTestdata1ForRegionD(t *testing.T) {
	input_seed := [][]int{
		{1, 3},
	}

	perimeter := calcPerimeter(input_seed)
	if perimeter != 4 {
		t.Fatalf(`"4" Perimeter didn't return 4. It returned %d`, perimeter)
	}
}

func TestPerimeterForTestdata1ForRegionE(t *testing.T) {
	input_seed := [][]int{
		{3, 0},
		{3, 1},
		{3, 2},
	}

	perimeter := calcPerimeter(input_seed)
	if perimeter != 8 {
		t.Fatalf(`"E" Perimeter didn't return 8. It returned %d`, perimeter)
	}
}
