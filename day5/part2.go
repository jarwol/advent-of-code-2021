package day5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part2() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	area := init2(scanner)

	intersections := 0
	for _, count := range area {
		if count > 1 {
			intersections++
		}
	}
	fmt.Println(intersections)
}

func getSlope(pt1 Point, pt2 Point) Fraction {
	numerator := pt2.y - pt1.y
	denominator := pt2.x - pt1.x

	gcd := abs(gcd(numerator, denominator))
	if denominator == 0 {
		if numerator < 0 {
			return Fraction{-1, 0}
		} else {
			return Fraction{1, 0}
		}
	} else if numerator == 0 {
		if denominator < 0 {
			return Fraction{0, -1}
		} else {
			return  Fraction{0, 1}
		}
	}
	return Fraction{numerator / gcd, denominator / gcd}
}

func getNextPoint(pt Point, slope Fraction) Point {
	return Point{pt.x + slope.denominator, pt.y + slope.numerator}
}

func init2(scanner *bufio.Scanner) map[Point]int {
	area := make(map[Point]int)
	p := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	for scanner.Scan() {
		matches := p.FindStringSubmatch(scanner.Text())
		x1, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		x2, _ := strconv.Atoi(matches[3])
		y2, _ := strconv.Atoi(matches[4])
		a := Point{x1, y1}
		b := Point{x2, y2}
		m := getSlope(a, b)

		for pt := a; pt != b; pt = getNextPoint(pt, m) {
			area[pt]++
		}
		area[b]++
	}
	return area
}