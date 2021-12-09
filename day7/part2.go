package day7

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Part2() {
	line, err := ioutil.ReadFile("day7/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	crabs := toInts(strings.Split(string(line), ","))

	min := dist(crabs, 0)
	for i := 1; i < max(crabs); i++ {
		d := dist(crabs, i)
		if d < min {
			min = d
		}
	}

	fmt.Println(min)
}

func dist(nums []int, position int) int {
	sum := 0
	for _, num := range nums {
		for i := 1; i <= abs(num - position); i++ {
			sum += i
		}
	}
	return sum
}
