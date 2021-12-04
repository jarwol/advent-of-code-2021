package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BitPositions struct {
    zeros, ones []string
}

func Part2() {
	file, err := os.Open("day3/input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
 
    scanner := bufio.NewScanner(file)
    // countsByDigit := make(map[int]BitPositions)
    var oxygenNums []string
    var co2Nums []string
    for scanner.Scan() {
        num := scanner.Text()
        oxygenNums = append(oxygenNums, num)
        co2Nums = append(co2Nums, num)

    }
    size := len(oxygenNums[0])
    for i := 0; i < size && len(oxygenNums) > 1; i++ {
        zeros, ones := groupByDigit(oxygenNums, i)
        if len(zeros) > len(ones) {
            oxygenNums = zeros
        } else {
            oxygenNums = ones
        }
    }

    for i := 0; i < size && len(co2Nums) > 1; i++ {
        zeros, ones := groupByDigit(co2Nums, i)
        if len(zeros) > len(ones) {
            co2Nums = ones
        } else {
            co2Nums = zeros
        }
    }


    // for i, bits := range countsByDigit {
    //     if len(bits.ones) > len(bits.zeros) {
    //         bits.ones = filter(bits.ones, func (n string) bool { return string(n[i]) == "1" })
    //         bits.zeros = filter(bits.zeros, func (n string) bool { return string(n[i]) == "0" })
    //     } else if len(bits.ones) < len(bits.zeros) {
    //         bits.ones = filter(bits.ones, func (n string) bool { return string(n[i]) == "0" })
    //         bits.zeros = filter(bits.zeros, func (n string) bool { return string(n[i]) == "1" })
    //     }
    // }

    oxygenVal, _ := strconv.ParseInt(oxygenNums[0], 2, 64)
    co2Val, _ := strconv.ParseInt(co2Nums[0], 2, 64)

    fmt.Println("O2:   " + fmt.Sprint(oxygenVal))
    fmt.Println("CO2: " + fmt.Sprint(co2Val))
    fmt.Println("Rating: " + fmt.Sprint(oxygenVal * co2Val))
}

func filter(s []string, f func(string) bool) []string {
    filtered := []string{}
    for _, val := range s {
        if f(val) {
            filtered = append(filtered, val)
        }
    }
    return filtered
}

func groupByDigit(nums []string, digit int) ([]string, []string) {
    ones := []string{}
    zeros := []string{}
    for _, num := range nums {
        if string(num[digit]) == "0" {
            zeros = append(zeros, num)
        } else {
            ones = append(ones, num)
        }
    }
    return zeros, ones
}