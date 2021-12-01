package day1

import (
    "fmt"
    "os"
	"bufio"
	"strconv"
)

func Part1() {
	file, err := os.Open("day1/depths.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
 
    scanner := bufio.NewScanner(file)
	prevDepth := 999999999999999
	increases := 0
    for scanner.Scan() {
        depth, _ := strconv.Atoi(scanner.Text())
		if depth > prevDepth {
			increases++
		}
		prevDepth = depth
    }
 
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    } else {
		fmt.Println(increases)
	}
}