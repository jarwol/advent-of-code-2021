package day6

import (
	"bufio"
	"strconv"
	"strings"
)

func initialize(scanner *bufio.Scanner) []uint64 {
	fish := make([]uint64, 9)
	scanner.Scan()
	daysLeft := strings.Split(scanner.Text(), ",")

	for _, days := range daysLeft {
		numDays, _ := strconv.ParseInt(days, 10, 64)
		fish[numDays]++
	}
	return fish
}

func simulateDay(fish []uint64) []uint64 {
	nextDay := make([]uint64, len(fish))
	for daysLeft, count := range fish {
		if daysLeft == 0 {
			nextDay[8] += count
			nextDay[6] += count
		} else {
			nextDay[daysLeft-1] += count
		}
	}
	return nextDay
}

func countFish(fish []uint64) uint64 {
	var sum uint64 = 0
	for _, count := range fish {
		sum += count
	}
	return sum
}