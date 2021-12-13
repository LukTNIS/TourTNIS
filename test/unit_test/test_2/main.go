package main

import "math"

func Sum(x int, y int) int {
	return x + y
}

func SquareRoot(input float64) float64 {
	return math.Sqrt(input)
}

func main() {
	Sum(5, 5)
	SquareRoot(4)
}
