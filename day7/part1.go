package day7

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Part1() {
	line, err := ioutil.ReadFile("day7/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	crabs := toInts(strings.Split(string(line), ","))

	min := distance(crabs, 0)
	for i := 1; i < max(crabs); i++ {
		d := distance(crabs, i)
		if d < min {
			min = d
		}
	}

	fmt.Println(min)
}

func distance(nums []int, position int) int {
	sum := 0
	for _, num := range nums {
		sum += abs(num - position)
	}
	return sum
}
