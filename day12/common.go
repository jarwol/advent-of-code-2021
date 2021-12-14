package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cave struct {
	name      string
	isLarge   bool
	neighbors []*Cave
}

func Initialize(fileName string) map[string]*Cave {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	caves := make(map[string]*Cave)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		cave1 := caves[parts[0]]
		cave2 := caves[parts[1]]

		if cave1 == nil {
			cave1 = &Cave{
				name:      parts[0],
				isLarge:   isUpperCase(parts[0]),
				neighbors: []*Cave{}}
			caves[cave1.name] = cave1
		}
		if cave2 == nil {
			cave2 = &Cave{
				name:      parts[1],
				isLarge:   isUpperCase(parts[1]),
				neighbors: []*Cave{}}
			caves[cave2.name] = cave2
		}
		cave1.neighbors = append(cave1.neighbors, cave2)
		cave2.neighbors = append(cave2.neighbors, cave1)
	}
	return caves
}

func isUpperCase(str string) bool {
	return str == strings.ToUpper(str)
}

func (cave *Cave) Print() {
	neighbors := []string{}
	for _, neighbor := range cave.neighbors {
		neighbors = append(neighbors, neighbor.name)
	}
	fmt.Println(cave.name + " - " + strings.Join(neighbors, ","))
}
