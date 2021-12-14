package day14

import (
	"fmt"
)

func Part1() {
	pairs, rules, counts := initialize("day14/input.txt")

	for i := 1; i <= 10; i++ {
		pairs = step(pairs, rules, counts)
	}

	maxCount, minCount := getMaxMin(counts)
	
	fmt.Println(maxCount - minCount)
}

