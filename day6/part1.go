package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fish := initialize(scanner)

	for i := 0; i < 80; i++ {
		fish = simulateDay(fish)
	}
	fmt.Println(countFish(fish))
}

