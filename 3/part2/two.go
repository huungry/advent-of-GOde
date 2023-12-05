package main

import (
	"aoc/common"
	"fmt"
	"regexp"
	"strconv"
)

type LineNumber struct {
	value int
}

type PotentialGearIndex struct {
	index int
}

type PartNumber struct {
	value         int
	previousIndex int
	nextIndex     int
}

var linesPotentialGearIndexes = make(map[LineNumber][]PotentialGearIndex)
var partNumbers = make(map[LineNumber][]PartNumber)
var gearRatio int

func main() {
	lines := common.ReadFile("./live")

	reGear := regexp.MustCompile(`\*`)
	reNumbers := regexp.MustCompile(`\d+`)

	for lineNumber, line := range lines {
		lineNumberTyped := LineNumber{lineNumber}
		// indexes of symbols
		gearIndexesPairs := reGear.FindAllStringIndex(line, -1)
		var potentialGearIndexes []PotentialGearIndex
		for _, gearIndexesPair := range gearIndexesPairs {
			potentialGearIndexes = append(potentialGearIndexes, PotentialGearIndex{gearIndexesPair[0]}) // only one letter expected
		}
		linesPotentialGearIndexes[lineNumberTyped] = append(linesPotentialGearIndexes[lineNumberTyped], potentialGearIndexes...)

		// numbers
		numberPairs := reNumbers.FindAllString(line, -1)
		var numberPairsInt []int
		for _, numberPair := range numberPairs {
			asInt, _ := strconv.Atoi(numberPair)
			numberPairsInt = append(numberPairsInt, asInt)
		}

		// indexes of numbers
		numberIndexesPairs := reNumbers.FindAllStringIndex(line, -1)

		// zip to PartNumber - numbers with extended to previous and next indexes
		for index, numberIndexesPair := range numberIndexesPairs {
			previousNeighbour := 0
			if numberIndexesPair[0] > 0 {
				previousNeighbour = numberIndexesPair[0] - 1
			}
			partNumbers[lineNumberTyped] = append(partNumbers[lineNumberTyped], PartNumber{numberPairsInt[index], previousNeighbour, numberIndexesPair[1]})
		}
	}

	for lineNumber, potentialGearIndexes := range linesPotentialGearIndexes {
		for _, potentialGearIndex := range potentialGearIndexes {
			var gearPartNumbersValues []int

			previousLinePartNumbers := partNumbers[LineNumber{lineNumber.value - 1}]
			thisLinePartNumbers := partNumbers[LineNumber{lineNumber.value}]
			nextLinePartNumbers := partNumbers[LineNumber{lineNumber.value + 1}]

			for _, previousLinePartNumber := range previousLinePartNumbers {
				if isBetween(potentialGearIndex.index, previousLinePartNumber.previousIndex, previousLinePartNumber.nextIndex) {
					gearPartNumbersValues = append(gearPartNumbersValues, previousLinePartNumber.value)
				}
			}

			for _, thisLinePartNumber := range thisLinePartNumbers {
				if isBetween(potentialGearIndex.index, thisLinePartNumber.previousIndex, thisLinePartNumber.nextIndex) {
					gearPartNumbersValues = append(gearPartNumbersValues, thisLinePartNumber.value)
				}
			}

			for _, nextLinePartNumber := range nextLinePartNumbers {
				if isBetween(potentialGearIndex.index, nextLinePartNumber.previousIndex, nextLinePartNumber.nextIndex) {
					gearPartNumbersValues = append(gearPartNumbersValues, nextLinePartNumber.value)
				}
			}

			if len(gearPartNumbersValues) == 2 {
				thisPartRatioValue := gearPartNumbersValues[0] * gearPartNumbersValues[1]
				gearRatio += thisPartRatioValue
			}
		}
	}

	fmt.Println(gearRatio)
}

func isBetween(num, min, max int) bool {
	return num >= min && num <= max
}
