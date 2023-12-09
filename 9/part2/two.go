package main

import (
	"aoc/common"
	"fmt"
	"regexp"
	"strconv"
)

type Line struct {
	values []int
}

type LineHistory struct {
	lines []Line
}

var lines []Line
var lineHistories = map[int]LineHistory{}

var result int

var re = regexp.MustCompile("-?[0-9]+")

func main() {
	inputLines := common.ReadFile("./live")

	for _, inputLine := range inputLines {
		numbersStr := re.FindAllString(inputLine, -1)
		thisLineValues := []int{}
		for _, numberStr := range numbersStr {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}
			thisLineValues = append(thisLineValues, number)
		}
		lines = append(lines, Line{thisLineValues})
	}

	for index, line := range lines {
		calculateHistory(line, index)
	}

	for index, lineHistory := range lineHistories {
		calculateSecret(index, lineHistory)
	}
}

func calculateHistory(line Line, index int) {
	differences := []int{}
	for i := 0; i < len(line.values)-1; i++ {
		differences = append(differences, line.values[i+1]-line.values[i])
	}

	allZeroes := true
	for _, difference := range differences {
		if difference != 0 {
			allZeroes = false
		}
	}

	lineHistories[index] = LineHistory{append(lineHistories[index].lines, line)}

	if allZeroes {
		lineHistories[index] = LineHistory{append(lineHistories[index].lines, Line{differences})}
		return
	} else {
		calculateHistory(Line{differences}, index)
	}
}

func calculateSecret(index int, lineHistory LineHistory) {
	linesInHistory := len(lineHistory.lines)
	lastLine := lineHistory.lines[linesInHistory-1]

	lastLine.values = append(lastLine.values, 0)

	for i := linesInHistory - 2; i >= 0; i-- {
		nextLineInHistory := lineHistory.lines[i+1]
		firstInNextLine := nextLineInHistory.values[0]

		firstInThisLine := lineHistory.lines[i].values[0]

		toAddInThisLine := firstInThisLine - firstInNextLine

		lineHistory.lines[i].values = append([]int{toAddInThisLine}, lineHistory.lines[i].values...)

		if i == 0 {
			result += toAddInThisLine
			fmt.Println("Next value in history", toAddInThisLine, "Result", result)
		}
	}
}
