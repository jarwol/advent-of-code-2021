package day8

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Part2() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		signals := mapStrings(strings.Split(parts[0], " "), sortString)
		output := mapStrings(strings.Split(parts[1], " "), sortString)
		decodedSignals, decodedDigits := decodeUniques(signals)
		decodedSignals = decodeSignals(decodedSignals, decodedDigits)
		sum += decodeOutput(decodedSignals, output)
	}

	fmt.Println(sum)
}

func decodeSignals(signalMap map[string]int, digitMap map[int]string) map[string]int {
	for signal, val := range signalMap {
		if val == -1 {
			if len(signal) == 5 && containsChars(signal, digitMap[7]) {
				signalMap[signal] = 3
				digitMap[3] = signal
			} else if len(signal) == 6 {
				if containsChars(signal, digitMap[4]) {
					signalMap[signal] = 9
					digitMap[9] = signal
				} else if !containsChars(signal, digitMap[1]) {
					signalMap[signal] = 6
					digitMap[6] = signal
				} else {
					signalMap[signal] = 0
					digitMap[0] = signal
				}
			}
		}
	}

	for signal, val := range signalMap {
		if val == -1 {
			if containsChars(digitMap[6], signal) {
				signalMap[signal] = 5
				digitMap[5] = signal
			} else {
				signalMap[signal] = 2
				digitMap[2] = signal
			}
		}
	}
	return signalMap
}

func decodeUniques(signals []string) (map[string]int, map[int]string) {
	decoded := make(map[string]int)
	reverse := make(map[int]string)
	for _, signal := range signals {
		num := len(signal)
		if num == 2 {
			decoded[signal] = 1
			reverse[1] = signal
		} else if num == 4 {
			decoded[signal] = 4
			reverse[4] = signal
		} else if num == 3 {
			decoded[signal] = 7
			reverse[7] = signal
		} else if num == 7 {
			decoded[signal] = 8
			reverse[8] = signal
		} else {
			decoded[signal] = -1
		}
	}
	return decoded, reverse
}

func decodeOutput(signals map[string]int, output []string) int {
	res := make([]int, 4)
	for i, digit := range output {
		res[i] = signals[digit]
	}
	return res[0] * 1000 + res[1] * 100 + res[2] * 10 + res[3]
}

func containsChars(str string, chars string) bool {
	for _, char := range chars {
		if !strings.ContainsRune(str, char) {
			return false
		}
	}
	return true
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func mapStrings(strs []string, f func (string) string) []string {
	res := make([]string, len(strs))
	for i, str := range strs {
		res[i] = f(str)
	}
	return res
}