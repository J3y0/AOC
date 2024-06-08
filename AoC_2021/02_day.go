package main

import (
	"fmt"
	"main/utils"
	"os"
)

func main() {
	lines, err := utils.ParseLines("./input/day2.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// ----- Part 1 -----
	part1, err := ComputeFinalPositionPart1(lines)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// ----- Part 2 -----
	part2, err := ComputeFinalPositionPart2(lines)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	utils.FormatAndPrintResultWithoutTime(part1, part2)
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
