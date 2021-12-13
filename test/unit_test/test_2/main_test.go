package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}

func TestSquareRoot(t *testing.T) {
	tables := []struct {
		x float64
		y float64
	}{
		{1, 1},
		{4, 2},
		{16, 4},
		{9, 3},
	}

	for _, table := range tables {
		total := SquareRoot(table.x)
		if total != table.y {
			t.Errorf("Sum of %v was incorrect, got: %v, want: %v.", table.x, table.y, total)
		}
	}
}
