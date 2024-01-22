package main

import (
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func (r *Range) Len() uint {
	return uint(r.End - r.Start + 1)
}

func (r *Range) IsEmpty() bool {
	return r.End < r.Start
}

func NewRange(start, end int) *Range {
	return &Range{
		Start: start,
		End:   end,
	}
}

func (pr PartRange) ComputeWaysFromRange() int {
	var total = 1
	for _, r := range pr {
		total *= int(r.Len())
	}
	return total
}

type PartRange map[rune]Range

func (pr PartRange) Clone() PartRange {
	newPr := make(PartRange)
	for k, v := range pr {
		newPr[k] = v
	}
	return newPr
}

func Part2(partRange PartRange, workflows *Workflows, moduleName string) int {
	if moduleName == "R" {
		return 0
	}
	if moduleName == "A" {
		return partRange.ComputeWaysFromRange()
	}

	var total int
	for _, instr := range (*workflows)[moduleName].Rules {
		step := strings.Split(instr, ":")
		rule := step[0]
		characteristic := rune(rule[0])

		value, _ := strconv.Atoi(rule[2:])
		var (
			passPartRangeCpy = make(PartRange)
			pass             *Range
			failed           *Range
		)
		switch rule[1] {
		case '<':
			pass = NewRange(partRange[characteristic].Start, value-1) // Range that goes to next module
			failed = NewRange(value, partRange[characteristic].End)   // Range that continues to next rule
		case '>':
			pass = NewRange(value+1, partRange[characteristic].End)   // Range that goes to next module
			failed = NewRange(partRange[characteristic].Start, value) // Range that continues to next rule
		}

		// If range is not empty, we pursue to the next module
		if !pass.IsEmpty() {
			passPartRangeCpy = partRange.Clone()
			passPartRangeCpy[characteristic] = *pass
			total += Part2(passPartRangeCpy, workflows, step[1])
		}

		// If failed range is empty, we don't have to continue for next rule as nothing goes to next rule
		if failed.IsEmpty() {
			break
		}
		// Update the range for the characteristic that didn't pass the test and go to the next rule with this
		// updated map
		partRange[characteristic] = *failed
	}

	// Total is the sum of accepted sub-ranges that pass the conditions (obtained
	// in previous recursive calls) and sub-ranges that did not and fall into the fallback function
	return total + Part2(partRange, workflows, (*workflows)[moduleName].Fallback)
}
