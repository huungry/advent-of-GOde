package main

import (
	"aoc/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers []int
	bet            []int
}

var result uint64

func main() {
	lines := common.ReadFile("./live")
	re := regexp.MustCompile("[0-9]+")

	var cards []Card

	for _, line := range lines {
		cardAndNumbers := strings.Split(line, ":")
		card := cardAndNumbers[0]
		cardIdStr := re.FindString(card)
		cardId, err := strconv.Atoi(cardIdStr)
		if err != nil {
			panic(err)
		}

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

	cardsCount := make(map[int]uint64)

	for _, card := range cards {
		cardsCount[card.id] = 1
	}

	maxCardId := 0

	for _, card := range cards {
		if card.id > maxCardId {
			maxCardId = card.id
		}
	}

	for _, card := range cards {
		thisCardCount := cardsCount[card.id]

		for i := uint64(1); i <= thisCardCount; i++ {
			wins := 0
			for _, betNumber := range card.bet {
				if contains(card.winningNumbers, betNumber) {
					wins++
				}
			}

			for win := 1; win <= wins; win++ {
				if card.id+win <= maxCardId {
					cardsCount[card.id+win]++
				}
			}
		}
	}

	for _, cardCount := range cardsCount {
		result += cardCount
	}

	fmt.Println(result)
}

func contains(numbers []int, number int) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}
