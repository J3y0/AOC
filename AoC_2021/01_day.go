package main

import (
	"fmt"
	"main/utils"
	"os"
	"strings"
)

func main() {
	measurements, err := parseMeasurements("./input/day1.txt")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	// ----- Part 1 -----
	totalIncrease := CountDepthIncrease(measurements)

	// ----- Part 2 -----
	totalIncreaseGrouped := Count3GroupMeasurements(measurements)

	utils.FormatAndPrintResultWithoutTime(totalIncrease, totalIncreaseGrouped)
}

func Count3GroupMeasurements(measurementsList []int) (total int) {
	previousSum := measurementsList[0] + measurementsList[1] + measurementsList[2]
	for i := 3; i < len(measurementsList); i++ {
		currentSum := measurementsList[i-2] + measurementsList[i-1] + measurementsList[i]
		if previousSum < currentSum {
			total++
		}
		previousSum = currentSum
	}
	return
}

func CountDepthIncrease(measurementsList []int) (total int) {
	for i := range measurementsList {
		if i == 0 {
			continue
		}
		if measurementsList[i] > measurementsList[i-1] {
			total++
		}
	}

	return
}

func parseMeasurements(path string) (measurements []int, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	for _, line := range lines {
		var measure int
		_, err = fmt.Sscanf(line, "%d", &measure)
		if err != nil {
			return nil, err
		}
		measurements = append(measurements, measure)
	}

	return
}
