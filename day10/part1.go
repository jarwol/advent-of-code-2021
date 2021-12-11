package day10

import (
	"bufio"
	"fmt"
	"os"
)

var errors = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func Part1() {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	errorSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		corrupt, ch := isCorrupt(line)

		if corrupt {
			errorSum += errors[ch]
		}
	}

	fmt.Println(errorSum)
}

func isCorrupt(line string) (isCorrupt bool, firstCorruptChar rune) {
	chars := &Stack{}

	for _, ch := range line {
		if isOpeningChar(ch) {
			chars = chars.Push(ch)
		} else {
			lastChar := chars.Pop()
			if !isMatch(lastChar, ch) {
				return true, ch
			}
		}
	}
	return false, '0'
}
