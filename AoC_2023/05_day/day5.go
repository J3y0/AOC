package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"05_day/parsing"
	"05_day/utils"
)

func main() {
	data, err := os.ReadFile("./data/day5.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while reading file: %v", err)
		os.Exit(1)
	}

	// -------- Part 1 --------
	start := time.Now()
	part1, err := computeMinLocation(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while solving part 1: %v", err)
		os.Exit(1)
	}
	total_p1 := time.Since(start)

	fmt.Printf("Result for part 1: %d\n", part1)
	fmt.Printf("Time for part 1: %s\n\n", total_p1)

	// -------- Part 2 --------
	start = time.Now()
	part2, err := computeMinLocationPart2(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while solving part 2: %v", err)
		os.Exit(1)
	}
	total_p2 := time.Since(start)

	fmt.Printf("Result for part 2: %d\n", part2)
	fmt.Printf("Time for part 2: %s\n", total_p2)
}

func computeMinLocation(data []byte) (uint, error) {
	blocks := strings.Split(string(data), "\n\n")
	seeds, err := parsing.ParseSeeds(blocks[0])
	if err != nil {
		return 0, err
	}
	transformationRanges, err := parsing.ParseTransformations(blocks[1:])
	if err != nil {
		return 0, err
	}

	var minLocation uint = math.MaxUint
	for _, seed := range seeds {
		var location uint
		var stepName string = "seed"
		var inputAtStep int = seed
		for stepName != "location" {
			location, stepName = transformationRanges[stepName].Transform(inputAtStep)
			inputAtStep = int(location)
		}

		// Update min if necessary
		if location < minLocation {
			minLocation = location
		}
	}

	return minLocation, nil
}

func computeMinLocationPart2(data []byte) (uint, error) {
	blocks := strings.Split(string(data), "\n\n")
	ranges, err := parsing.ParseSeedRanges(blocks[0])
	if err != nil {
		return 0, err
	}
	transformationRanges, err := parsing.ParseTransformations(blocks[1:])
	if err != nil {
		return 0, err
	}

	var stepName string = "seed"
	for stepName != "location" {
		shiftedRanges := make([]*utils.Range, 0)

		transformations := transformationRanges[stepName].Ranges
		for _, transformation := range transformations {
			ranges_cpy := ranges
			for _, rg := range ranges_cpy {
				overlapRange := rg.FindOverlap(transformation.Range)
				if overlapRange != nil {
					newRanges := rg.SplitRange(*overlapRange)
					ranges = utils.Remove(ranges, rg)
					ranges = append(ranges, newRanges...)

					// Shift occurs
					overlapRange.ShiftRange(transformation.Shift)
					shiftedRanges = append(shiftedRanges, overlapRange)
				}
			}
		}

		ranges = append(ranges, shiftedRanges...)
		stepName = transformations[0].DestRangeName
	}
	return utils.FindMinFromRanges(ranges), nil
}
