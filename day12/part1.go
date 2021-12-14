package day12

import (
	"fmt"
)

func Part1() {
	caves := Initialize("day12/input.txt")
	for _, cave := range caves {
		cave.Print()
	}
	paths := visit(caves["start"], []string{}, make(map[string]bool))

	fmt.Println(len(paths))
}

func visit(cave *Cave, path []string, visited map[string]bool) [][]string {
	if contains(path, "end") && cave.name == "end" {
		fmt.Println("HERE")
	}
	path = append(path, cave.name)
	newPaths := [][]string{}
	if cave.name == "end" {
		return append(newPaths, path)
	}
	visitedCopy := copyMap(visited)
	if !cave.isLarge {
		if visitedCopy[cave.name] {
			return [][]string{}
		} else {
			visitedCopy[cave.name] = true
		}
	}

	for _, neighbor := range cave.neighbors {
		newPaths = append(newPaths, visit(neighbor, path, visitedCopy)...)
	}
	return newPaths
}

func copyMap(source map[string]bool) map[string]bool {
	target := make(map[string]bool)
	for key, val := range source {
		target[key] = val
	}
	return target
}