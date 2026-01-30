package days

import (
	"main/utils"
	"math"
	"strconv"
	"strings"
)

type Day7 struct {
	crabs []int
	minX  int
	maxX  int
}

func (d *Day7) Parse(input string) error {
	input = strings.TrimSpace(input)
	crabs := strings.Split(input, ",")

	parsedCrabs := make([]int, len(crabs))
	for i, crab := range crabs {
		crabInt, err := strconv.Atoi(crab)
		if err != nil {
			return err
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
	d.minX = minX
	d.maxX = maxX
	d.crabs = parsedCrabs

	return nil
}

func (d *Day7) Part1() (int, error) {
	return minFuel(d.crabs, d.minX, d.maxX, true), nil
}

func (d *Day7) Part2() (int, error) {
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
