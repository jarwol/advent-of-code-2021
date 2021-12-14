package day13

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Fold struct {
	line string
	val int
}

func Initialize(fileName string) ([][]bool, []Fold) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dotPattern := regexp.MustCompile(`(\d+),(\d+)`)
	foldPattern := regexp.MustCompile(`fold along (x|y)=(\d+)`)
	var maxX, maxY int
	points := []Point{}
	folds := []Fold{}
	for scanner.Scan() {
		line := scanner.Text()
		matches := dotPattern.FindStringSubmatch(line)
		if matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			points = append(points, Point{x: x, y: y})
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		} else {
			matches = foldPattern.FindStringSubmatch(line)
			if matches != nil {
				val, _ := strconv.Atoi(matches[2])
				folds = append(folds, Fold{line: matches[1], val: val})
			}
		}
	}

	paper := make([][]bool, maxY + 1)
	for i := 0; i <= maxY; i++ {
		paper[i] = make([]bool, maxX + 1)
	}

	for _, pt := range points {
		paper[pt.y][pt.x] = true
	}

	return paper, folds
}

func doFold(paper [][]bool, fold Fold) [][]bool {
	var folded [][]bool
	if fold.line == "x" {
		folded = make([][]bool, len(paper))
		for i := 0; i < len(paper); i++ {
			folded[i] = make([]bool, fold.val)
		}
		for y := 0; y < len(paper); y++ {
			for x := 0; x < len(paper[y]); x++ {
				if x > fold.val {
					distanceFromLine := x - fold.val
					newX := fold.val - distanceFromLine
					folded[y][newX] = paper[y][newX] || paper[y][x]
				} else if x < fold.val {
					folded[y][x] = paper[y][x]
				}
			}
		}
	} else {
		folded = make([][]bool, fold.val)
		for i := 0; i < fold.val; i++ {
			folded[i] = make([]bool, len(paper[0]))
		}
		for y := 0; y < len(paper); y++ {
			for x := 0; x < len(paper[y]); x++ {
				if y > fold.val {
					distanceFromLine := y - fold.val
					newY := fold.val - distanceFromLine
					folded[newY][x] = paper[newY][x] || paper[y][x]
				} else if y < fold.val {
					folded[y][x] = paper[y][x]
				}
			}
		}
	}
	return folded
}

func printPaper(paper [][]bool) {
	for _, row := range paper {
		line := []string{}
		for _, dot := range row {
			if dot {
				line = append(line, "#")
			} else {
				line = append(line, ".")
			}
		}
		fmt.Println(strings.Join(line, " "))
	}
}