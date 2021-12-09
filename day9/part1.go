package day9

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	heights := initialize(scanner)

	risk := 0
	for i, row := range heights {
		for j, height := range row {
			pt := Point{i, j, height}
			neighbors := getNeighbors(heights, pt)
			if isLowPoint(pt, neighbors) {
				risk += height + 1
			}
		}
	}

	fmt.Println(risk)
}
