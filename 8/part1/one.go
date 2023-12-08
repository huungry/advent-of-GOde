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

var re = regexp.MustCompile("[A-Z]+")

var instructions []rune
var mapLines = map[Destination]Moves{}

var instructionsLength int
var startingDestination = Destination{"AAA"}
var endDestination = Destination{"ZZZ"}
var counter = 0

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

	run(startingDestination, 0)
}

func run(destination Destination, instructionIndex int) {
	counter++

	var nextDestination Destination

	if instructions[instructionIndex] == 'L' {
		nextDestination = mapLines[destination].left
	} else {
		nextDestination = mapLines[destination].right
	}

	// fmt.Println("Current destination:", destination, "Direction", instructions[instructionIndex], "Next destination:", nextDestination, "Counter:", counter)

	if nextDestination != endDestination {
		if instructionIndex+1 >= instructionsLength {
			run(nextDestination, 0)
		} else {
			run(nextDestination, instructionIndex+1)
		}
	} else {
		fmt.Println("Got it! Counter:", counter)
		return
	}

}
