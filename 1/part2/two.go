package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"aoc/common"
)

var replacements = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"0":     "0",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func main() {
	lines := common.ReadFile("./live")

	var numbers []int

	for _, line := range lines {
		firstIndex := math.MaxInt
		lastIndex := 0
		var first string
		var last string
		for toReplace, replacement := range replacements {
			thisFirstIndex := strings.Index(line, toReplace)
			thisLastIndex := strings.LastIndex(line, toReplace)
			if thisFirstIndex == -1 {
				continue
			}
			if thisFirstIndex < firstIndex {
				firstIndex = thisFirstIndex
				first = replacement
			}
			if thisLastIndex > lastIndex {
				lastIndex = thisLastIndex
				last = replacement
			}
		}

		if last == "" {
			last = first
		}

		number, _ := strconv.Atoi(first + last)
		numbers = append(numbers, number)
	}

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Println("Sum:", sum)
}
