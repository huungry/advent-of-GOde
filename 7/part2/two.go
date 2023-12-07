package main

import (
	"aoc/common"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Bet struct {
	value int
}

type Round struct {
	cards []int
	bet   Bet
}

type RoundResult struct {
	cards        []int
	bet          Bet
	fiveMatches  int
	fullHouse    []int
	fourMatches  int
	threeMatches int
	twoPairs     []int
	twoMatches   int
	oneMatches   []int
}

var cardsStrengths = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var rounds []Round
var roundsResults []RoundResult

var re = regexp.MustCompile("[0-9]+")

var result uint64

func main() {
	lines := common.ReadFile("./live")

	for _, line := range lines {
		split := strings.Split(line, " ")

		cardsStr := split[0]
		cardsRunes := []rune(cardsStr)

		betStr := re.FindString(split[1])
		bet, err := strconv.Atoi(betStr)
		if err != nil {
			panic(err)
		}

		var cardsStrength []int
		for _, card := range cardsRunes {
			cardsStrength = append(cardsStrength, cardsStrengths[card])
		}

		rounds = append(rounds, Round{cards: cardsStrength, bet: Bet{bet}})
	}

	for _, round := range rounds {
		fiveMatches, jokersUsed := isNumberOfKind(5, round.cards, []int{}, 0)
		fourMatches, jokersUsed := isNumberOfKind(4, round.cards, fiveMatches, jokersUsed)
		threeMatches, jokersUsed := isNumberOfKind(3, round.cards, append(fiveMatches, fourMatches...), jokersUsed)
		twoMatches, jokersUsed := isNumberOfKind(2, round.cards, append(fiveMatches, append(fourMatches, threeMatches...)...), jokersUsed)
		oneMatches, _ := isNumberOfKind(1, round.cards, append(fiveMatches, append(fourMatches, append(threeMatches, twoMatches...)...)...), jokersUsed)

		fullHouse := findFullHouse(threeMatches, twoMatches)
		twoPairs := findTwoPairs(twoMatches)

		fiveMatchesFinal := 0
		if len(fiveMatches) > 0 {
			fiveMatchesFinal = fiveMatches[0]
		}

		fourMatchesFinal := 0
		if len(fourMatches) > 0 {
			fourMatchesFinal = fourMatches[0]
		}

		threeMatchesFinal := 0
		if len(threeMatches) > 0 {
			threeMatchesFinal = threeMatches[0]
		}

		twoMatchesFinal := 0
		if len(twoMatches) > 0 {
			twoMatchesFinal = twoMatches[0]
		}

		roundsResults = append(roundsResults, RoundResult{
			cards:        round.cards,
			bet:          round.bet,
			fiveMatches:  fiveMatchesFinal,
			fourMatches:  fourMatchesFinal,
			fullHouse:    fullHouse,
			threeMatches: threeMatchesFinal,
			twoPairs:     twoPairs,
			twoMatches:   twoMatchesFinal,
			oneMatches:   oneMatches,
		})
	}

	roundsResults = sortByCards(roundsResults)

	fiveOfKindsRounds := []RoundResult{}
	fourOfKindsRounds := []RoundResult{}
	fullHousesRounds := []RoundResult{}
	threeOfKindsRounds := []RoundResult{}
	twoOfKindsRounds := []RoundResult{}
	twoPairsRounds := []RoundResult{}
	oneOfKindsRounds := []RoundResult{}

	for _, roundResult := range roundsResults {
		if roundResult.fiveMatches > 0 {
			fiveOfKindsRounds = append(fiveOfKindsRounds, roundResult)
		} else if roundResult.fourMatches > 0 {
			fourOfKindsRounds = append(fourOfKindsRounds, roundResult)
		} else if len(roundResult.fullHouse) > 0 {
			fullHousesRounds = append(fullHousesRounds, roundResult)
		} else if roundResult.threeMatches > 0 {
			threeOfKindsRounds = append(threeOfKindsRounds, roundResult)
		} else if len(roundResult.twoPairs) > 0 {
			twoPairsRounds = append(twoPairsRounds, roundResult)
		} else if roundResult.twoMatches > 0 {
			twoOfKindsRounds = append(twoOfKindsRounds, roundResult)
		} else if len(roundResult.oneMatches) > 0 {
			oneOfKindsRounds = append(oneOfKindsRounds, roundResult)
		}
	}

	mergedAll := []RoundResult{}
	mergedAll = append(mergedAll, oneOfKindsRounds...)
	mergedAll = append(mergedAll, twoOfKindsRounds...)
	mergedAll = append(mergedAll, twoPairsRounds...)
	mergedAll = append(mergedAll, threeOfKindsRounds...)
	mergedAll = append(mergedAll, fullHousesRounds...)
	mergedAll = append(mergedAll, fourOfKindsRounds...)
	mergedAll = append(mergedAll, fiveOfKindsRounds...)

	for index, roundResult := range mergedAll {
		bet := roundResult.bet.value
		result += uint64(bet) * (uint64(index) + 1)
	}

	fmt.Println("Result: ", result)
}

func isNumberOfKind(targetCount int, cards []int, alreadyTargetedHigher []int, jokersUsed int) ([]int, int) {
	occurrences := make(map[int]int)
	matches := []int{}
	cardsToProcess := []int{}

	jokers := 0
	for _, card := range cards {
		if card == 1 {
			jokers++
		}
	}

	jokers = jokers - jokersUsed
	
	if jokers == 5 {
		return []int{1}, 5 // edge case if all jokers
	}

	for _, card := range cards {
		if !contains(alreadyTargetedHigher, card) {
			cardsToProcess = append(cardsToProcess, card)
		}
	}

	for i := 0; i < len(cardsToProcess); i++ {
		occurrences[cardsToProcess[i]]++
	}

	for card, occurance := range occurrences {
		if card != 1 {
			for jokersToUse := 0; jokersToUse <= jokers; jokersToUse++ {
				if occurance+jokersToUse == targetCount {
					jokersUsed = jokersUsed + jokersToUse
					jokers = jokers - jokersToUse
					if jokers > -1 {
						matches = append(matches, card)
					}
				}
			}
		}
	}
	return matches, jokersUsed
}

func findFullHouse(threes []int, twos []int) []int {
	for _, three := range threes {
		for _, two := range twos {
			if three != two {
				return []int{three, two}
			}
		}
	}

	return []int{}
}

func findTwoPairs(twos []int) []int {
	for _, two1 := range twos {
		for _, two2 := range twos {
			if two1 != two2 {
				return []int{two1, two2}
			}
		}
	}

	return []int{}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func sortByCards(rounds []RoundResult) []RoundResult {

	sort.SliceStable(rounds, func(i, j int) bool {
		for k := range rounds[i].cards {
			if rounds[i].cards[k] != rounds[j].cards[k] {
				return rounds[i].cards[k] < rounds[j].cards[k]
			}
		}
		return len(rounds[i].cards) < len(rounds[j].cards)
	})

	return rounds
}
