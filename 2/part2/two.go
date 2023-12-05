package main

import (
	"aoc/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var availableColors = []string{
	"red",
	"green",
	"blue",
}

var gameRules = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	lines := common.ReadFile("./live")

	re := regexp.MustCompile("[0-9]+")

	var powerSum int

	for _, line := range lines {
		withoutSpaces := strings.ReplaceAll(line, " ", "")
		split := strings.Split(withoutSpaces, ":")

		gameId, _ := strconv.Atoi(re.FindString(split[0]))
		rounds := strings.Split(split[1], ";")

		minimumNumberOfCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		fmt.Println("Welcome to game", gameId)

		for _, round := range rounds {
			roundColors := strings.Split(round, ",")
			for _, roundColor := range roundColors {
				for _, availableColor := range availableColors {
					if strings.Contains(roundColor, availableColor) {
						numberOfCubesThisColor, _ := strconv.Atoi(re.FindString(roundColor))
						if numberOfCubesThisColor > minimumNumberOfCubes[availableColor] {
							minimumNumberOfCubes[availableColor] = numberOfCubesThisColor
						}
					}
				}
			}
		}

		fmt.Println("Minimum number of cubes:", minimumNumberOfCubes, "for game", gameId)

		powerOfThisGame := 1

		for _, minimumNumberOfCubesThisColor := range minimumNumberOfCubes {
			powerOfThisGame *= minimumNumberOfCubesThisColor
		}

		powerSum += powerOfThisGame

		fmt.Println("Power of this game:", powerOfThisGame)
	}

	fmt.Println("Power sum:", powerSum)
}
