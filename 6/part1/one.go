package main

import (
	"aoc/common"
	"fmt"
	"regexp"
	"strconv"
)

type TimeDistance struct {
	time     int
	distance int
}

var re = regexp.MustCompile("[0-9]+")
var result = 1

func main() {
	lines := common.ReadFile("./live")

	for _, line := range lines {
		fmt.Println(line)
	}

	timesStr := re.FindAllString(lines[0], -1)
	distancesStr := re.FindAllString(lines[1], -1)

	var times []int
	for _, timeStr := range timesStr {
		time, err := strconv.Atoi(timeStr)
		if err != nil {
			panic(err)
		}
		times = append(times, time)
	}

	var distances []int
	for _, distanceStr := range distancesStr {
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}
		distances = append(distances, distance)
	}

	var timeDistances []TimeDistance

	for i := 0; i < len(times); i++ {
		timeDistances = append(timeDistances, TimeDistance{times[i], distances[i]})
	}

	for _, timeDistance := range timeDistances {
		possibleWaysToWin := 0
		for chargingForTime := 1; chargingForTime < timeDistance.time; chargingForTime++ {
			travelingFor := timeDistance.time - chargingForTime
			distance := chargingForTime * travelingFor
			if distance > timeDistance.distance {
				possibleWaysToWin++
			}
		}
		result *= possibleWaysToWin
	}

	fmt.Println("Result", result)
}
