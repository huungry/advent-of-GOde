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

type SymbolIndex struct {
	index int
}

type PartNumber struct {
	value         int
	previousIndex int
	nextIndex     int
}

var linesSymbolsIndexes = make(map[LineNumber][]SymbolIndex)
var partNumbers = make(map[LineNumber][]PartNumber)
var sum int

func main() {
	lines := common.ReadFile("./live")

	reSymbol := regexp.MustCompile(`[^\w\s.]`)
	reNumbers := regexp.MustCompile(`\d+`)

	for lineNumber, line := range lines {
		lineNumberTyped := LineNumber{lineNumber}
		// indexes of symbols
		symbolIndexesPairs := reSymbol.FindAllStringIndex(line, -1)
		var symbolIndexes []SymbolIndex
		for _, symbolIndexesPair := range symbolIndexesPairs {
			symbolIndexes = append(symbolIndexes, SymbolIndex{symbolIndexesPair[0]}) // only one letter expected
		}
		linesSymbolsIndexes[lineNumberTyped] = append(linesSymbolsIndexes[lineNumberTyped], symbolIndexes...)

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

	for lineNumber, partNumbers := range partNumbers {
		for _, partNumber := range partNumbers {
			shouldBeCounted := false

			previousLineSymbolIndexes := linesSymbolsIndexes[LineNumber{lineNumber.value - 1}]
			thisLineSymbolIndexes := linesSymbolsIndexes[LineNumber{lineNumber.value}]
			nextLineSymbolIndexes := linesSymbolsIndexes[LineNumber{lineNumber.value + 1}]

			partNumberMinIndex := partNumber.previousIndex
			partNumberMaxIndex := partNumber.nextIndex

			for _, previousLineSymbolIndex := range previousLineSymbolIndexes {
				if previousLineSymbolIndex.index >= partNumberMinIndex && previousLineSymbolIndex.index <= partNumberMaxIndex {
					shouldBeCounted = true
					break
				}
			}

			for _, thisLineSymbolIndex := range thisLineSymbolIndexes {
				if thisLineSymbolIndex.index == partNumberMinIndex || thisLineSymbolIndex.index == partNumberMaxIndex {
					shouldBeCounted = true
					break
				}

			}

			for _, nextLineSymbolIndex := range nextLineSymbolIndexes {
				if nextLineSymbolIndex.index >= partNumberMinIndex && nextLineSymbolIndex.index <= partNumberMaxIndex {
					shouldBeCounted = true
					break
				}
			}

			if shouldBeCounted {
				sum += partNumber.value
			}
		}
	}

	fmt.Println(sum)
}
