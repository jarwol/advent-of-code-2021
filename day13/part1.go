package day13

import (
	"fmt"
)

func Part1() {
	paper, folds := Initialize("day13/input.txt")
	
	paper = doFold(paper, folds[0])
	
	fmt.Println(countDots(paper))
}

func countDots(paper [][]bool) int {
	count := 0
	for _, row := range paper {
		for _, dot := range row {
			if dot {
				count++
			}
		}
	}
	return count
}
