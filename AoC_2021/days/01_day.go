package days

import (
	"fmt"
	"os"
	"strings"
)

type Day1 struct {
	measurements []int
}

func (d *Day1) Part1() (int, error) {
	measurements, err := parseMeasurements("./input/01_day.txt")
	if err != nil {
		return 0, err
	}
	d.measurements = measurements

	return CountDepthIncrease(measurements), nil
}

func (d *Day1) Part2() (int, error) {
	if len(d.measurements) == 0 {
		measurements, err := parseMeasurements("./input/01_day.txt")
		if err != nil {
			return 0, err
		}
		d.measurements = measurements
	}

	return Count3GroupMeasurements(d.measurements), nil
}

var Solution1 = Day1{}

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
