package main

import (
	"fmt"
	"regexp"
	"strconv"

	"aoc/common"
)

func main() {
	lines := common.ReadFile("./live")

	re := regexp.MustCompile("[0-9]")
	var numbers []int

	for _, line := range lines {
		numbersString := re.FindAllString(line, -1)

		if len(numbersString) > 0 {
			first := numbersString[0]
			last := numbersString[len(numbersString)-1]
			numberPair, _ := strconv.Atoi(first + last)
			numbers = append(numbers, numberPair)
		}
	}

	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("Sum:", sum)
}
