package day7

import "strconv"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(nums []int) int {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func toInts(vals []string) []int {
	ints := make([]int, len(vals))
	for i, str := range vals {
		val, _ := strconv.Atoi(str)
		ints[i] = val
	}
	return ints
}