package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part2() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	nums := toIntArray(strings.Split(scanner.Text(), ","))
	boards := initializeBoards(scanner)
	completedMap := make(map[int]bool)
	var lastCompleted Board
	var lastNum int

	for _, num := range nums {
		for i, board := range boards {
			if !completedMap[i] && board.Mark(num) {
				completedMap[i] = true
				lastCompleted = board
				lastNum = num
				if len(completedMap) == len(boards) {
					fmt.Println(lastCompleted.Sum() * lastNum)
					os.Exit(0)
				}
			}
		}
	}

	fmt.Println(lastCompleted.Sum() * lastNum)
}
