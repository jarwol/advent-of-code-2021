package common

import "strconv"

func ToIntSlice(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		val, err := strconv.Atoi(str)
		if err != nil {
			panic("Cannot convert " + str + " to int.")
		}
		ints[i] = val
	}
	return ints
}