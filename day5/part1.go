package day5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	area := initialize(scanner)
	fmt.Println(area)
	intersections := 0
	for _, count := range area {
		if count > 1 {
			intersections++
		}
	}
	fmt.Println(intersections)
}

func initialize(scanner *bufio.Scanner) map[Point]int {
	area := make(map[Point]int)
	p := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	for scanner.Scan() {
		matches := p.FindStringSubmatch(scanner.Text())
		x1, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		x2, _ := strconv.Atoi(matches[3])
		y2, _ := strconv.Atoi(matches[4])

		minX := x1
		maxX := x2
		minY := y1
		maxY := y2
		if x2 < x1 {
			minX = x2
			maxX = x1
		}
		if y2 < y1 {
			minY = y2
			maxY = y1
		}

		if x1 == x2 || y1 == y2 {
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					pt := Point{x, y}
					area[pt]++
				}
			}
		}
	}
	return area
}
