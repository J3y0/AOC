package days

import (
	"main/utils"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day7 struct {
	crabs []int
	minX  int
	maxX  int
}

func (d *Day7) Part1() (int, error) {
	crabs, minX, maxX, err := parseCrabs("./input/07_day.txt")
	if err != nil {
		return 0, err
	}
	d.crabs = crabs
	d.minX = minX
	d.maxX = maxX

	return minFuel(d.crabs, d.minX, d.maxX, true), nil
}

func (d *Day7) Part2() (int, error) {
	if len(d.crabs) == 0 {
		crabs, minX, maxX, err := parseCrabs("./input/07_day.txt")
		if err != nil {
			return 0, err
		}
		d.crabs = crabs
		d.minX = minX
		d.maxX = maxX
	}

	return minFuel(d.crabs, d.minX, d.maxX, false), nil
}

func minFuel(crabs []int, minX, maxX int, part1 bool) int {
	// Find min fuel
	minFuel := math.MaxInt
	for x := minX; x < maxX; x++ {
		fuel := 0
		for _, crab := range crabs {
			if part1 {
				fuel += utils.AbsInt(crab - x)
			} else {
				steps := utils.AbsInt(crab - x)
				fuel += (steps * (steps + 1)) / 2
			}
		}

		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func parseCrabs(path string) ([]int, int, int, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, 0, 0, err
	}

	crabs := strings.Split(string(content), ",")

	parsedCrabs := make([]int, len(crabs))
	for i, crab := range crabs {
		crabInt, err := strconv.Atoi(crab)
		if err != nil {
			return nil, 0, 0, err
		}

		parsedCrabs[i] = crabInt
	}

	// Find min and max crab's horizontal pos
	minX := parsedCrabs[0]
	maxX := parsedCrabs[0]
	for _, crab := range parsedCrabs {
		if crab > maxX {
			maxX = crab
		} else if crab < minX {
			minX = crab
		}
	}

	return parsedCrabs, minX, maxX, nil
}
