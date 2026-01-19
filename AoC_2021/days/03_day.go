package days

import (
	"main/utils"
	"strconv"
)

type Day3 struct {
	binaries []string
}

func (d *Day3) Parse(input string) error {
	d.binaries = utils.ParseLines(input)
	return nil
}

func (d *Day3) Part1() (int, error) {
	return ComputePowerConsumption(d.binaries), nil
}

func (d *Day3) Part2() (int, error) {
	return ComputeLifeSupportRating(d.binaries)
}

func ComputeLifeSupportRating(lines []string) (result int, err error) {
	lengthBinary := len(lines[0])

	oxygenLines := lines
	var oxygenGeneratorRating int64
	for i := range lengthBinary {
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
	for i := range lengthBinary {
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
	for i := range lengthBinary {
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
