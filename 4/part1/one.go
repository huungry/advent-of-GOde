package main

import (
	"aoc/common"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers []int
	bet            []int
}

var result int

func main() {
	lines := common.ReadFile("./live")
	re := regexp.MustCompile("[0-9]+")

	var cards []Card

	for _, line := range lines {
		cardAndNumbers := strings.Split(line, ":")
		card := cardAndNumbers[0]
		cardIdStr := strings.Replace(card, "Card ", "", 1)
		cardId, _ := strconv.Atoi(cardIdStr)

		numbers := cardAndNumbers[1]
		numbersSplit := strings.Split(numbers, "|")

		winningNumbersStr := numbersSplit[0]
		winningNumbersArrayStr := re.FindAllString(winningNumbersStr, -1)

		var winningNumbers []int

		for _, winningNumberStr := range winningNumbersArrayStr {
			winningNumber, _ := strconv.Atoi(winningNumberStr)
			winningNumbers = append(winningNumbers, winningNumber)
		}

		betStr := numbersSplit[1]
		betArrayStr := re.FindAllString(betStr, -1)

		var bet []int

		for _, betStr := range betArrayStr {
			betNumber, _ := strconv.Atoi(betStr)
			bet = append(bet, betNumber)
		}

		cards = append(cards, Card{cardId, winningNumbers, bet})
	}

	for _, card := range cards {
		wins := 0
		for _, betNumber := range card.bet {
			if contains(card.winningNumbers, betNumber) {
				wins++
			}
		}
		result += 1 * int(math.Pow(2, float64(wins-1)))
	}

	println(result)
}

func contains(numbers []int, number int) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}
