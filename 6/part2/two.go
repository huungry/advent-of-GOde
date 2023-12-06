package main

import (
	"aoc/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type TimeDistance struct {
	time     int
	distance int
}

var re = regexp.MustCompile("[0-9]+")

func main() {
	lines := common.ReadFile("./live")

	for _, line := range lines {
		fmt.Println(line)
	}

	timesNoSpaces := strings.ReplaceAll(lines[0], " ", "")
	distancesNoSpaces := strings.ReplaceAll(lines[1], " ", "")

	timeStr := re.FindAllString(timesNoSpaces, -1)
	distanceStr := re.FindAllString(distancesNoSpaces, -1)

	time, err := strconv.Atoi(timeStr[0])
	if err != nil {
		panic(err)
	}
	distance, err := strconv.Atoi(distanceStr[0])
	if err != nil {
		panic(err)
	}

	timeDistance := TimeDistance{time, distance}

	possibleWaysToWin := 0
	for chargingForTime := 1; chargingForTime < timeDistance.time; chargingForTime++ {
		travelingFor := timeDistance.time - chargingForTime
		distance := chargingForTime * travelingFor
		if distance > timeDistance.distance {
			possibleWaysToWin++
		}
	}

	fmt.Println("Result", possibleWaysToWin)
}
