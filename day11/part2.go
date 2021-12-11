package day11

import (
	"fmt"
)

func Part2() {
	octopuses := initialize("day11/input.txt")
	flashCount := 0
	step := 0
	for flashCount != 100 {
		flashCount = executeStep(octopuses)
		step++

	}

	fmt.Println(step)
}
