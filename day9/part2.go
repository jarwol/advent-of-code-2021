package day9

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Part2() {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	heights := initialize(scanner)
	lowPoints := getLowPoints(heights)
	basins := []int{}
	for _, pt := range lowPoints {
		basin := getBasin(heights, pt, make(map[Point]bool))
		basins = append(basins, len(basin))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	fmt.Println(basins[0] * basins[1] * basins[2])
}

func getBasin(heights [][]int, pt Point, visited map[Point]bool) map[Point]bool {
	if visited[pt] || pt.height == 9 {
		return visited
	}
	visited[pt] = true
	neighbors := getNeighbors(heights, pt)

	for _, neighbor := range neighbors {
		getBasin(heights, neighbor, visited)
	}
	return visited
}

func getLowPoints(heights [][]int) []Point {
	lowPoints := []Point{}
	for i, row := range heights {
		for j, height := range row {
			pt := Point{i, j, height}
			neighbors := getNeighbors(heights, pt)
			if isLowPoint(pt, neighbors) {
				lowPoints = append(lowPoints, pt)
			}
		}
	}
	return lowPoints
}