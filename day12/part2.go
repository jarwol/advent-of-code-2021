package day12

import (
	"fmt"
)

func Part2() {
	caves := Initialize("day12/input.txt")
	for _, cave := range caves {
		cave.Print()
	}
	paths := visit2(caves["start"], []string{}, make(map[string]int))

	fmt.Println(len(paths))
	for _, path := range paths {
		fmt.Println(path)
	}
}

func visit2(cave *Cave, path []string, visited map[string]int) [][]string {
	newPath := make([]string, len(path))
	copy(newPath, path)
	newPath = append(newPath, cave.name)
	newPaths := [][]string{}
	if cave.name == "end" {
		return append(newPaths, newPath)
	}
	visitedCopy := copyMap2(visited)
	if !cave.isLarge {
		if visitedCopy[cave.name] > 1 {
			return [][]string{}
		} else {
			visitedCopy[cave.name]++
		}
	}

	for _, neighbor := range cave.neighbors {
		if neighbor.name != "start" {
			newPaths = append(newPaths, visit2(neighbor, newPath, visitedCopy)...)
		}
	}
	return newPaths
}

func copyMap2(source map[string]int) map[string]int {
	target := make(map[string]int)
	for key, val := range source {
		target[key] = val
	}
	return target
}

func contains(strs []string, target string) bool {
	for _, str := range strs {
		if str == target {
			return true
		}
	}
	return false
}