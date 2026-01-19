package days

import (
	"fmt"
	"main/utils"
)

type Day2 struct {
	instructions []string
}

func (d *Day2) Parse(input string) error {
	d.instructions = utils.ParseLines(input)
	return nil
}

func (d *Day2) Part1() (int, error) {
	return ComputeFinalPositionPart1(d.instructions)
}

func (d *Day2) Part2() (int, error) {
	return ComputeFinalPositionPart2(d.instructions)
}

func ComputeFinalPositionPart1(lines []string) (int, error) {
	var horizontal, depth int
	for _, line := range lines {
		var (
			instruction string
			times       int
		)
		_, err := fmt.Sscanf(line, "%s %d", &instruction, &times)
		if err != nil {
			return 0, err
		}

		switch instruction {
		case "forward":
			horizontal += times
		case "down":
			depth += times
		case "up":
			depth -= times
		}
	}

	return horizontal * depth, nil
}

func ComputeFinalPositionPart2(lines []string) (int, error) {
	var horizontal, depth, aim int
	for _, line := range lines {
		var (
			instruction string
			times       int
		)
		_, err := fmt.Sscanf(line, "%s %d", &instruction, &times)
		if err != nil {
			return 0, err
		}

		switch instruction {
		case "forward":
			horizontal += times
			depth += aim * times
		case "down":
			aim += times
		case "up":
			aim -= times
		}
	}

	return horizontal * depth, nil
}
