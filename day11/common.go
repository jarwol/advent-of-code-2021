package day11

import (
	"bufio"
	"os"
	"strings"

	"com.jarwol/advent-of-code-2021/common"
)

type Stack struct {
	items []common.Point
}

func (stack *Stack) Len() int {
	return len(stack.items)
}

func (stack *Stack) Push(item common.Point) *Stack {
	stack.items = append(stack.items, item)
	return stack
}

func (stack *Stack) Pop() common.Point {
	item := stack.items[len(stack.items) - 1]
	stack.items = stack.items[:len(stack.items) - 1]
	return item
}


func initialize(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := common.ToIntSlice(strings.Split(line, ""))
		matrix = append(matrix, row)
	}
	return matrix
}

func getNeighbors(matrix [][]int, pt common.Point) []common.Point {
	neighbors := []common.Point{}
	
	if pt.I > 0 {
		neighbors = append(neighbors, common.Point{pt.I - 1, pt.J})
		if pt.J > 0 {
			neighbors = append(neighbors, common.Point{pt.I - 1, pt.J - 1})
		}
	}
	if pt.J > 0 {
		neighbors = append(neighbors, common.Point{pt.I, pt.J - 1})
		if pt.I < len(matrix) - 1 {
			neighbors = append(neighbors, common.Point{pt.I + 1, pt.J - 1})
		}
	}
	if pt.I < len(matrix) - 1 {
		neighbors = append(neighbors, common.Point{pt.I + 1, pt.J})
		if pt.J < len(matrix[0]) - 1 {
			neighbors = append(neighbors, common.Point{pt.I + 1, pt.J + 1})
		}
	}
	if pt.J < len(matrix[0]) - 1 {
		neighbors = append(neighbors, common.Point{pt.I, pt.J + 1})
		if pt.I > 0 {
			neighbors = append(neighbors, common.Point{pt.I - 1, pt.J + 1})
		}
	}
	return neighbors
}

func executeStep(octopuses [][]int) int {
	toFlash := incrementEnergy(octopuses)

	flashed := make(map[common.Point]bool)
	
	for toFlash.Len() > 0 {
		pt := toFlash.Pop()
		toFlash = flash(octopuses, pt, flashed, toFlash)
	}
	resetFlashed(octopuses, flashed)
	return len(flashed)
}

func resetFlashed(octopuses [][]int, flashed map[common.Point]bool) {
	for pt := range flashed {
		octopuses[pt.I][pt.J] = 0
	}
}

func incrementEnergy(octopuses [][]int) *Stack {
	toFlash := &Stack{}
	for i, row := range octopuses {
		for j := range row {
			octopuses[i][j]++
			if octopuses[i][j] > 9 {
				toFlash = toFlash.Push(common.Point{I: i, J: j})
			}
		}
	}
	return toFlash
}

func flash(octopuses [][]int, pt common.Point, flashed map[common.Point]bool, toFlash *Stack) *Stack {
	if !flashed[pt] {
		flashed[pt] = true
		neighbors := getNeighbors(octopuses, pt)
		
		for _, neighbor := range neighbors {
			octopuses[neighbor.I][neighbor.J]++
			if octopuses[neighbor.I][neighbor.J] > 9 {
				toFlash = toFlash.Push(neighbor)
			}
		}
	}

	return toFlash
}