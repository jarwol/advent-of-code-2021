package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day2/path.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
 
    scanner := bufio.NewScanner(file)
    x := 0
    z := 0
    for scanner.Scan() {
        parts := strings.Split(scanner.Text(), " ")
        direction := parts[0]
        val, _ := strconv.Atoi(parts[1])
		if direction == "forward" {
			x += val
		} else if direction == "up" {
            z -= val
        } else if direction == "down" {
            z += val
        } else {
            fmt.Println("Invalid direction: " + direction)
            return
        }
    }
 
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    } else {
		fmt.Println(x * z)
	}
}