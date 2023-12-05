package main

import (
	"aoc/common"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Seed struct {
	value uint64
}

type AtoB struct {
	destinationRangeStart uint64
	sourceRangeStart      uint64
	rangeLength           uint64
}

var seeds []Seed

var mappers [][]AtoB = make([][]AtoB, 7)
var minLocation uint64 = math.MaxUint64

var re = regexp.MustCompile("[0-9]+")

func main() {
	lines := common.ReadFile("./live")

	// split by empty line
	split := strings.Split(strings.Join(lines, "\n"), "\n\n")

	seedsStr := split[0]
	seedsNumbersStr := re.FindAllString(seedsStr, -1)

	for _, seedNumberStr := range seedsNumbersStr {
		seedNumber, err := strconv.Atoi(seedNumberStr)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, Seed{uint64(seedNumber)})
	}

	for index, split := range split[1:] {
		buildMap(split, index+1)
	}

	for _, seed := range seeds {
		calculateForSeed(seed.value, 0)
	}

	fmt.Println("Min location is", minLocation)
}

func calculateForSeed(source uint64, index int) {
	nextSource := source
	for _, mapper := range mappers[index] {
		if isBetween(source, mapper.sourceRangeStart, mapper.sourceRangeStart+mapper.rangeLength-1) {
			nextSource = source + (mapper.destinationRangeStart - mapper.sourceRangeStart)
		}
	}

	if index < 6 {
		calculateForSeed(nextSource, index+1)
	} else {
		if nextSource < minLocation {
			minLocation = nextSource
		}
	}
}

func buildMap(split string, index int) {
	elementsWithoutTitle := strings.Split(split, "\n")[1:]
	build(elementsWithoutTitle, &mappers[index-1])
}

func build(lines []string, mapToUpdate *[]AtoB) {
	for _, line := range lines {
		elementStr := re.FindAllString(line, -1)

		destinationRangeStart, err := strconv.Atoi(elementStr[0])
		if err != nil {
			panic(err)
		}
		sourceRangeStart, err := strconv.Atoi(elementStr[1])
		if err != nil {
			panic(err)
		}
		rangeLength, err := strconv.Atoi(elementStr[2])
		if err != nil {
			panic(err)
		}

		*mapToUpdate = append(*mapToUpdate, AtoB{uint64(destinationRangeStart), uint64(sourceRangeStart), uint64(rangeLength)})
	}
}

func isBetween(num, min, max uint64) bool {
	return num >= min && num <= max
}
