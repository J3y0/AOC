package days

import (
	"fmt"
	"main/utils"
)

type Day2 struct {
	Instructions []string
}

func (d *Day2) Part1() (int, error) {
	instructions, err := utils.ParseLines("./input/02_day.txt")
	if err != nil {
		return 0, err
	}
	d.Instructions = instructions

	return ComputeFinalPositionPart1(d.Instructions)
}

func (d *Day2) Part2() (int, error) {
	if len(d.Instructions) == 0 {
		instructions, err := utils.ParseLines("./input/02_day.txt")
		if err != nil {
			return 0, err
		}
		d.Instructions = instructions
	}

	return ComputeFinalPositionPart2(d.Instructions)
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
