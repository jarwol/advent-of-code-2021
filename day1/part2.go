package day1

import (
    "fmt"
    "os"
	"bufio"
	"strconv"
)

func Part2() {
	file, err := os.Open("day1/depths.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
 
    scanner := bufio.NewScanner(file)
    var depths []int
    
    for scanner.Scan() {
        depth, _ := strconv.Atoi(scanner.Text())
        depths = append(depths, depth)
    }


    increases := 0
    prevDepth := 999999999999999
    for i := 0; i <= len(depths) - 3; i++ {
        depth := calcWindowSum(depths, i, 3)
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

func calcWindowSum(slice []int, start int, windowSize int) int {
    sum := 0
    for i := start; i < start + windowSize; i++ {
        sum += slice[i]
    }
    return sum
}