package main

import (
	"aoc/common"
	"fmt"
)

type Position struct {
	lineIndex int
	charIndex int
}

type Line struct {
	values []rune
}

var lines = map[int]Line{}
var startingLineIndexAndPos Position
var routesFromStart = []Position{}
var alreadyVisited = map[Position][]Position{}       // starting position -> position
var reachedStartingPosition = map[Position]bool{}    // startingposition -> bool
var startingPositionCounter = map[Position]int{}     // startingposition -> counter
var startingPositionsMaxCounter = map[Position]int{} // startingposition -> max value counter

var counter = 0

func main() {
	inputLines := common.ReadFile("./live")

	for lineIndex, line := range inputLines {
		lines[lineIndex] = Line{values: []rune(line)}
	}

	for lineIndex, line := range lines {
		for charIndex, char := range line.values {
			if char == 'S' {
				startingLineIndexAndPos = Position{lineIndex: lineIndex, charIndex: charIndex}
			}
		}
	}

	// for _, line := range lines {
	// 	fmt.Printf("%s\n", string(line.values))
	// }

	listRoutesFromStart()
	// for _, route := range routesFromStart {
	// 	fmt.Printf("%d %d\n", route.lineIndex, route.charIndex)
	// }

	// for _, startingPos := range routesFromStart {
	// 	startingPositionCounter[startingPos] = 0
	// 	calculateNextStep(startingPos, startingPos.lineIndex, startingPos.charIndex)
	// }

	fmt.Println("Starting position:", startingLineIndexAndPos)
	calculateNextStep(startingLineIndexAndPos, startingLineIndexAndPos.lineIndex, startingLineIndexAndPos.charIndex)

	fmt.Println(counter / 2)
}

func calculateNextStep(startingPos Position, lineIndex int, charIndex int) {
	left := Position{lineIndex: lineIndex, charIndex: charIndex - 1}
	right := Position{lineIndex: lineIndex, charIndex: charIndex + 1}
	up := Position{lineIndex: lineIndex - 1, charIndex: charIndex}
	down := Position{lineIndex: lineIndex + 1, charIndex: charIndex}

	fmt.Printf("Current position: %d %d\n", lineIndex, charIndex)

	currentPos := Position{lineIndex, charIndex}
	currentChar := lines[lineIndex].values[charIndex]

	if currentPos != startingPos {
		alreadyVisited[startingPos] = append(alreadyVisited[startingPos], Position{lineIndex, charIndex})
	}

	canGoLeftFromCurrentChar := currentChar == '-' || currentChar == 'J' || currentChar == '7'
	canGoRightFromCurrentChar := currentChar == '-' || currentChar == 'L' || currentChar == 'F'
	canGoUpFromCurrentChar := currentChar == '|' || currentChar == 'L' || currentChar == 'J'
	canGoDownFromCurrentChar := currentChar == '|' || currentChar == 'F' || currentChar == '7'

	if startingPos == currentPos && startingPositionCounter[startingPos] > 0 {
		fmt.Printf("Reached starting position: %d %d. Counter %d\n", lineIndex, charIndex, startingPositionCounter[startingPos])
		if (startingPositionCounter[startingPos]) > counter {
			counter = startingPositionCounter[startingPos]
		}
		reachedStartingPosition[startingPos] = true
		return
	}

	startingPositionCounter[startingPos] = startingPositionCounter[startingPos] + 1

	leftChar := 'X'
	if left.charIndex >= 0 {
		wasVisited := false
		for _, alreadyVisitedPos := range alreadyVisited[startingPos] {
			if alreadyVisitedPos == left {
				wasVisited = true
			}
		}
		if !wasVisited {
			leftChar = lines[left.lineIndex].values[left.charIndex]
		}
	}
	if canGoLeftFromCurrentChar && leftChar == '-' || leftChar == 'L' || leftChar == 'F' {
		fmt.Printf("Going left because char is: %s\n", string(leftChar))
		calculateNextStep(startingPos, lineIndex, left.charIndex)
	}

	rightChar := 'X'
	if right.charIndex < len(lines[right.lineIndex].values) {
		wasVisited := false
		for _, alreadyVisitedPos := range alreadyVisited[startingPos] {
			if alreadyVisitedPos == right {
				fmt.Printf("Already visited right: %d %d\n", right.lineIndex, right.charIndex)
				wasVisited = true
			}
		}

		if !wasVisited {
			rightChar = lines[right.lineIndex].values[right.charIndex]
		}
	}
	if canGoRightFromCurrentChar && rightChar == '-' || rightChar == 'J' || rightChar == '7' {
		fmt.Printf("Going right because char is: %s\n", string(rightChar))
		calculateNextStep(startingPos, right.lineIndex, right.charIndex)
	}

	upChar := 'X'
	if up.lineIndex >= 0 {
		wasVisited := false
		for _, alreadyVisitedPos := range alreadyVisited[startingPos] {
			if alreadyVisitedPos == up {
				wasVisited = true
			}
		}

		if !wasVisited {
			upChar = lines[up.lineIndex].values[up.charIndex]
		}
	}
	if canGoUpFromCurrentChar && upChar == '|' || upChar == '7' || upChar == 'F' {
		fmt.Printf("Going up because char is: %s\n", string(upChar))
		calculateNextStep(startingPos, up.lineIndex, up.charIndex)
	}

	downChar := 'X'
	if down.lineIndex < len(lines) {
		wasVisited := false
		for _, alreadyVisitedPos := range alreadyVisited[startingPos] {
			if alreadyVisitedPos == down {
				wasVisited = true
			}
		}

		if !wasVisited {
			downChar = lines[down.lineIndex].values[down.charIndex]
		}
	}
	if canGoDownFromCurrentChar && downChar == '|' || downChar == 'J' || downChar == 'L' {
		fmt.Printf("Going down because char is: %s\n", string(downChar))
		calculateNextStep(startingPos, down.lineIndex, down.charIndex)
	}

	startingPositionCounter[startingPos] = startingPositionCounter[startingPos] - 1
	if startingPositionCounter[startingPos] > startingPositionsMaxCounter[startingPos] {
		startingPositionsMaxCounter[startingPos] = startingPositionCounter[startingPos]
	}

	fmt.Printf("No route found for: %d %d Current counter %d\n", lineIndex, charIndex, startingPositionCounter[startingPos])
}

func listRoutesFromStart() {
	left := Position{lineIndex: startingLineIndexAndPos.lineIndex, charIndex: startingLineIndexAndPos.charIndex - 1}
	right := Position{lineIndex: startingLineIndexAndPos.lineIndex, charIndex: startingLineIndexAndPos.charIndex + 1}
	up := Position{lineIndex: startingLineIndexAndPos.lineIndex - 1, charIndex: startingLineIndexAndPos.charIndex}
	down := Position{lineIndex: startingLineIndexAndPos.lineIndex + 1, charIndex: startingLineIndexAndPos.charIndex}

	leftChar := 'X'
	if left.charIndex >= 0 {
		leftChar = lines[left.lineIndex].values[left.charIndex]
	}
	if leftChar == '-' || leftChar == 'L' || leftChar == 'F' {
		routesFromStart = append(routesFromStart, left)
	}

	rightChar := 'X'
	if right.charIndex < len(lines[right.lineIndex].values) {
		rightChar = lines[right.lineIndex].values[right.charIndex]
	}
	if rightChar == '-' || rightChar == 'J' || rightChar == '7' {
		routesFromStart = append(routesFromStart, right)
	}

	upChar := 'X'
	if up.lineIndex >= 0 {
		upChar = lines[up.lineIndex].values[up.charIndex]
	}
	if upChar == '|' || upChar == '7' || upChar == 'F' {
		routesFromStart = append(routesFromStart, up)
	}

	downChar := 'X'
	if down.lineIndex < len(lines) {
		downChar = lines[down.lineIndex].values[down.charIndex]
	}
	if downChar == '|' || downChar == 'J' || downChar == 'L' {
		routesFromStart = append(routesFromStart, down)
	}

	replacement := 'X'

	routeFromStart1 := routesFromStart[0]
	routeFromStart2 := routesFromStart[1]

	if routeFromStart1 == left && routeFromStart2 == right || routeFromStart1 == right && routeFromStart2 == left {
		replacement = '-'
	} else if routeFromStart1 == up && routeFromStart2 == down || routeFromStart1 == down && routeFromStart2 == up {
		replacement = '|'
	} else if routeFromStart1 == left && routeFromStart2 == up || routeFromStart1 == up && routeFromStart2 == left {
		replacement = 'J'
	} else if routeFromStart1 == left && routeFromStart2 == down || routeFromStart1 == down && routeFromStart2 == left {
		replacement = '7'
	} else if routeFromStart1 == right && routeFromStart2 == up || routeFromStart1 == up && routeFromStart2 == right {
		replacement = 'L'
	} else if routeFromStart1 == right && routeFromStart2 == down || routeFromStart1 == down && routeFromStart2 == right {
		replacement = 'F'
	}

	fmt.Printf("Replacing %d %d with %s\n", startingLineIndexAndPos.lineIndex, startingLineIndexAndPos.charIndex, string(replacement))
	lines[startingLineIndexAndPos.lineIndex].values[startingLineIndexAndPos.charIndex] = replacement
}
