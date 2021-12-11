package day11

import (
	"fmt"
)

func Part1() {
	octopuses := initialize("day11/input.txt")
	flashCount := 0
	for i := 1; i <= 100; i++ {
		flashCount = executeStep(octopuses)
	}

	fmt.Println(flashCount)
}
