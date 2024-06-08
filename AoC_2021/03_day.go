package main

import (
	"fmt"
	"main/utils"
	"os"
	"strconv"
)

func main() {
	lines, err := utils.ParseLines("./input/day3.txt")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	// ----- Part1 -----
	part1 := ComputePowerConsumption(lines)

	// ----- Part2 -----
	part2, err := ComputeLifeSupportRating(lines)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	utils.FormatAndPrintResultWithoutTime(part1, part2)
}

func ComputeLifeSupportRating(lines []string) (result int, err error) {
	lengthBinary := len(lines[0])

	oxygenLines := lines
	var oxygenGeneratorRating int64
	for i := 0; i < lengthBinary; i++ {
		var nextOxygenLines []string
		var count0, count1 int
		for _, line := range oxygenLines {
			if line[i] == '0' {
				count0++
			} else {
				count1++
			}
		}

		var mostBitOxygen uint8
		if count0 > count1 {
			mostBitOxygen = '0'
		} else {
			mostBitOxygen = '1'
		}

		for _, line := range oxygenLines {
			if line[i] == mostBitOxygen {
				nextOxygenLines = append(nextOxygenLines, line)
			}
		}

		if len(nextOxygenLines) == 1 {
			oxygenGeneratorRating, err = strconv.ParseInt(nextOxygenLines[0], 2, 32)
			if err != nil {
				return 0, err
			}
			break
		}
		oxygenLines = nextOxygenLines
	}

	co2Lines := lines
	var co2ScrubberRating int64
	for i := 0; i < lengthBinary; i++ {
		var count0, count1 int
		for _, line := range co2Lines {
			if line[i] == '0' {
				count0++
			} else {
				count1++
			}
		}

		var mostBitCo2 uint8
		if count0 > count1 {
			mostBitCo2 = '1'
		} else {
			mostBitCo2 = '0'
		}
		var nextCo2Lines []string
		for _, line := range co2Lines {
			if line[i] == mostBitCo2 {
				nextCo2Lines = append(nextCo2Lines, line)
			}
		}

		if len(nextCo2Lines) == 1 {
			co2ScrubberRating, err = strconv.ParseInt(nextCo2Lines[0], 2, 32)
			if err != nil {
				return 0, err
			}
			break
		}

		co2Lines = nextCo2Lines
	}

	return int(co2ScrubberRating * oxygenGeneratorRating), nil
}

func ComputePowerConsumption(lines []string) int {
	lengthBinary := len(lines[0])
	var gammaRate int
	for i := 0; i < lengthBinary; i++ {
		var count0, count1 int
		for _, line := range lines {
			if line[i] == '0' {
				count0++
			} else {
				count1++
			}
		}

		if count0 < count1 {
			gammaRate += utils.PowInt(2, lengthBinary-i-1)
		}
	}
	epsilonRate := utils.PowInt(2, lengthBinary) - gammaRate - 1
	return gammaRate * epsilonRate
}
