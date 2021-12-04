package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day3/input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
 
    scanner := bufio.NewScanner(file)
    counts := make(map[int]int)
    for scanner.Scan() {
        num := scanner.Text()
        for i, digit := range num {
            if _, exists := counts[i]; !exists {
                counts[i] = 0
            }
            val, _ := strconv.Atoi(string(digit))
            if val == 0 {
                counts[i]--
            } else {
                counts[i]++
            }
        }
    }

    var gamma = make([]string, 12)
    var eps = make([]string, 12)
    for i, count := range counts {
        if count > 0 {
            gamma[i] = "1"
            eps[i] = "0"
        } else {
            gamma[i] = "0"
            eps[i] = "1"
        }
    }

    gammaDec, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
    epsDec, _ := strconv.ParseInt(strings.Join(eps, ""), 2, 64)

    fmt.Println("Gamma:   " + fmt.Sprint(gammaDec))
    fmt.Println("Epsilon: " + fmt.Sprint(epsDec))
    fmt.Println("Power: " + fmt.Sprint(gammaDec * epsDec))
}