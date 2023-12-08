package main

import (
	"aoc/common"
	"fmt"
	"regexp"
)

type Destination struct {
	name string
}

type Moves struct {
	left  Destination
	right Destination
}

var re = regexp.MustCompile(`[A-Za-z0-9]+`)

var instructions []rune
var mapLines = map[Destination]Moves{}

var instructionsLength int
var startingDestinations []Destination

var stepsToZ = map[Destination]uint64{}

func main() {
	lines := common.ReadFile("./live")

	instructionsStr := lines[0]
	instructions = []rune(instructionsStr)
	instructionsLength = len(instructions)
	mappersStr := lines[2:]

	for _, mapperStr := range mappersStr {
		split := re.FindAllString(mapperStr, -1)

		name := split[0]
		left := split[1]
		right := split[2]

		mapLines[Destination{name}] = Moves{Destination{left}, Destination{right}}
	}

	for lineDestinationName, _ := range mapLines {
		if []rune(lineDestinationName.name)[2] == 'A' {
			startingDestinations = append(startingDestinations, lineDestinationName)
		}
	}

	// calculate how many steps to Z for each starting destination
	for _, startingDestination := range startingDestinations {
		calculateAndSaveSteps(startingDestination, 0)
	}

	// calculate LCM for all starting destinations
	lcmResult := uint64(1)
	for i := 0; i < len(startingDestinations); i++ {
		lcmResult = calculateLcm(lcmResult, stepsToZ[startingDestinations[i]])
	}

	fmt.Println("Result:", lcmResult)
}

func calculateGcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func calculateLcm(a, b uint64) uint64 {
	return a * b / calculateGcd(a, b)
}

func calculateAndSaveSteps(destination Destination, instructionIndex int) {
	var startingDestination = destination
	var counter uint64 = 0
	for {
		counter++

		var nextDestination Destination

		if instructions[instructionIndex] == 'L' {
			nextDestination = mapLines[destination].left
		} else {
			nextDestination = mapLines[destination].right
		}

		if []rune(nextDestination.name)[2] == 'Z' {
			stepsToZ[startingDestination] = counter
			break
		}

		if instructionIndex+1 >= instructionsLength {
			instructionIndex = 0
		} else {
			instructionIndex++
		}

		destination = nextDestination
	}
}
