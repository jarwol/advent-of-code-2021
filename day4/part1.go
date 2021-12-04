package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part1() {
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

	for _, num := range nums {
		for _, board := range boards {
			if board.Mark(num) {
				fmt.Println(board.Sum() * num)
				return
			}
		}
	}

}

