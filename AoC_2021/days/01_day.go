package days

import (
	"fmt"
	"strings"
)

type Day1 struct {
	measurements []int
}

func (d *Day1) Parse(input string) (err error) {
	lines := strings.Split(input, "\n")
	measurements := make([]int, 0, len(lines))
	for _, line := range lines {
		var measure int
		_, err = fmt.Sscanf(line, "%d", &measure)
		if err != nil {
			return err
		}
		measurements = append(measurements, measure)
	}

	d.measurements = measurements
	return
}

func (d *Day1) Part1() (int, error) {
	return CountDepthIncrease(d.measurements), nil
}

func (d *Day1) Part2() (int, error) {
	return Count3GroupMeasurements(d.measurements), nil
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
