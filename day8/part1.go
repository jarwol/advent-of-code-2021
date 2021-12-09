package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part1() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalUniques := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		// signals := strings.Split(parts[0], " ")
		digits := strings.Split(parts[1], " ")
		totalUniques += countUniques(digits)
	}

	fmt.Println(totalUniques)
}

func countUniques(digits []string) int {
	sum := 0
	for _, digit := range digits {
		num := len(digit)
		if num == 2 || num == 3 || num == 4 || num == 7 {
			sum++
		}
	}
	return sum
}