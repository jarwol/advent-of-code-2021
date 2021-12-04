package day4

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type Board [][]int

func (board Board) Sum() int {
	sum := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] != -1 {
				sum += board[i][j]
			}
		}
	}
	return sum
}

func (board Board) Mark(num int) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == num {
				board[i][j] = -1
				if board.CheckRow(i) || board.CheckCol(j) {
					return true
				}
			}
		}
	}
	return false
}

func (board Board) CheckRow(row int) bool {
	for _, num := range board[row] {
		if num != -1 {
			return false
		}
	}
	return true
}

func (board Board) CheckCol(col int) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] != -1 {
			return false
		}
	}
	return true
}

func initializeBoards(scanner *bufio.Scanner) []Board {
	boards := []Board{}
	board := Board{}
	sep := regexp.MustCompile(" +")
    for scanner.Scan() {
        line := scanner.Text()
		if line == "" {
			if len(board) > 0 {
				boards = append(boards, board)
			}
			board = Board{}
		} else {
			row := toIntArray(sep.Split(strings.TrimSpace(line), -1))
			board = append(board, row)
		}
    }
	return boards
}

func toIntArray(arr []string) []int {
	ints := []int{}
	for _, str := range arr {
		num, _ := strconv.Atoi(str)
		ints = append(ints, num)
	}
	return ints
}