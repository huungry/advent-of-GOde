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

	validGamesIdSum := 0

	for _, line := range lines {
		isGameValid := true

		withoutSpaces := strings.ReplaceAll(line, " ", "")
		split := strings.Split(withoutSpaces, ":")

		gameId, _ := strconv.Atoi(re.FindString(split[0]))
		rounds := strings.Split(split[1], ";")

		fmt.Println("Welcome to game", gameId)

		for _, round := range rounds {
			if isGameValid {
				roundColors := strings.Split(round, ",")
				for _, roundColor := range roundColors {
					for _, availableColor := range availableColors {
						if strings.Contains(roundColor, availableColor) && isGameValid {
							numberOfCubesThisColor, _ := strconv.Atoi(re.FindString(roundColor))
							if numberOfCubesThisColor > gameRules[availableColor] {
								fmt.Println("Found", numberOfCubesThisColor, "cubes of color", availableColor, "this round is invalid")
								isGameValid = false
								break
							}
						}
					}
				}
			}
		}

		if isGameValid {
			validGamesIdSum += gameId
			fmt.Println("This game is valid - adding to id sum. The sum is now", validGamesIdSum)
		}
	}

	fmt.Println("Sum of valid games ids:", validGamesIdSum)
}
