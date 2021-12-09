package day9

import (
	"bufio"
	"strings"
	"com.jarwol/advent-of-code-2021/common"
)

type Point struct {
	i int
	j int
	height int
}

func initialize(scanner *bufio.Scanner) [][]int {
	heights := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		vals := common.ToIntSlice(strings.Split(line, ""))
		heights = append(heights, vals)
	}
	return heights
}

func isLowPoint(pt Point, neighbors []Point) bool {
	for _, neighbor := range neighbors {
		if neighbor.height <= pt.height {
			return false
		}
	}
	return true
}

func getNeighbors(heights [][]int, pt Point) []Point {
	neighbors := []Point{}
	
	if pt.i > 0 {
		neighbors = append(neighbors, Point{pt.i - 1, pt.j, heights[pt.i - 1][pt.j]})
	}
	if pt.j > 0 {
		neighbors = append(neighbors, Point{pt.i, pt.j - 1, heights[pt.i][pt.j - 1]})
	}
	if pt.i < len(heights) - 1 {
		neighbors = append(neighbors, Point{pt.i + 1, pt.j, heights[pt.i + 1][pt.j]})
	}
	if pt.j < len(heights[0]) - 1 {
		neighbors = append(neighbors, Point{pt.i, pt.j + 1, heights[pt.i][pt.j + 1]})
	}
	return neighbors
}