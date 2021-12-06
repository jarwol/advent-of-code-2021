package day5

import "fmt"

type Point struct {
	x int
	y int
}

type Fraction struct {
	numerator   int
	denominator int
}

func printMatrix(matrix [][]int) {
	for x := range matrix {
		fmt.Println(matrix[x])
	}
}

func mapToMatrix(area map[Point]int) [][]int {
	minX := 0
	minY := 0
	maxX := 0
	maxY := 0
	for pt := range area {
		if pt.x < minX {
			minX = pt.x
		}
		if pt.x > maxX {
			maxX = pt.x
		}
		if pt.y < minY {
			minY = pt.y
		}
		if pt.y > maxY {
			maxY = pt.y
		}
	}

	matrix := make([][]int, maxY+1)
	for y := 0; y <= maxY; y++ {
		matrix[y] = make([]int, maxX+1)
	}

	for pt := range area {
		matrix[pt.y][pt.x]++
	}
	return matrix
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x % y
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}