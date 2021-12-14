package day14

import (
	"bufio"
	"os"
	"regexp"
)

func step(pairs map[string]uint64, rules map[string]string, counts map[string]uint64) map[string]uint64 {
	newMap := copyMap(pairs)
	for pair, count := range pairs {
		char := rules[pair]
		if char != "" && count > 0 {
			newMap[pair] -= count
			newMap[string(pair[0])+char] += count
			newMap[char+string(pair[1])] += count
			counts[char] += count
		}
	}
	return newMap
}

func copyMap(pairs map[string]uint64) map[string]uint64 {
	target := make(map[string]uint64)
	for pair, count := range pairs {
		target[pair] = count
	}
	return target
}

func getMaxMin(counts map[string]uint64) (max uint64, min uint64) {
	var maxCount uint64 = 0
	var minCount uint64 = 99999999999999999

	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
		if count < minCount {
			minCount = count
		}
	}
	return maxCount, minCount
}

func initialize(fileName string) (map[string]uint64, map[string]string, map[string]uint64) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	templatePattern := regexp.MustCompile(`[A-Z]+`)
	pairPattern := regexp.MustCompile(`([A-Z]{2}) -> ([A-Z])`)
	rules := make(map[string]string)
	pairs := make(map[string]uint64)
	counts := make(map[string]uint64)
	for scanner.Scan() {
		line := scanner.Text()
		matches := pairPattern.FindStringSubmatch(line)
		if matches != nil {
			rules[matches[1]] = matches[2]
		} else if matches = templatePattern.FindStringSubmatch(line); matches != nil {
			template := matches[0]
			for i := 0; i < len(template)-1; i++ {

				pairs[template[i:i+2]]++
				counts[string(template[i])]++
			}
			counts[string(template[len(template)-1])]++

		}
	}

	return pairs, rules, counts
}
