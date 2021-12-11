package day10

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var completionScores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func Part2() {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scores := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		remainder, corrupt := getRemainder(line)

		if !corrupt {
			completionChars := getClosingChars(remainder)
			scores = append(scores, calcCompletionScore(completionChars))
		}
	}

	sort.Ints(scores)


	fmt.Println(scores[len(scores) / int(2)])
}

func calcCompletionScore(chars []rune) int {
	sum := 0
	for _, ch := range chars {
		sum *= 5
		sum += completionScores[ch]
	}
	return sum
}

func getClosingChars(open *Stack) []rune {
	size := open.Len()
	closing := make([]rune, size)
	for i := 0; i < size; i++ {
		ch := open.Pop()
		closing[i] = getClosingMatch(ch)
	}
	return closing
}

func getRemainder(line string) (remainer *Stack, isCorrupt bool) {
	chars := &Stack{}

	for _, ch := range line {
		if isOpeningChar(ch) {
			chars = chars.Push(ch)
		} else {
			lastChar := chars.Pop()
			if !isMatch(lastChar, ch) {
				return nil, true
			}
		}
	}
	return chars, false
}

func getClosingMatch(open rune) rune {
	switch open {
	case '(':
		return ')'
	case '{':
		return '}'
	case '[':
		return ']'
	case '<':
		return '>'
	}
	return '0'
}
