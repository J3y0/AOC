package utils

import "math"

func Max(y, x int) int {
	if x > y {
		return x
	}
	return y
}

func Min(y, x int) int {
	if x < y {
		return x
	}
	return y
}

type Range struct {
	RangeStart int
	RangeEnd   int
}

func RangeFromBound(start int, end int) *Range {
	if start >= end {
		return nil
	}
	return &Range{RangeStart: start, RangeEnd: end}
}

func RangeFromSize(start int, size int) *Range {
	return &Range{RangeStart: start, RangeEnd: start + size - 1}
}

func (r *Range) ShiftRange(shift int) {
	r.RangeStart += shift
	r.RangeEnd += shift
}

func (r *Range) FindOverlap(otherRange Range) *Range {
	overlapStart := Max(r.RangeStart, otherRange.RangeStart)
	overlapEnd := Min(r.RangeEnd, otherRange.RangeEnd)
	return RangeFromBound(overlapStart, overlapEnd)
}

func (r *Range) SplitRange(overlapped Range) []*Range {
	var result []*Range

	if overlapped.RangeStart > r.RangeStart {
		result = append(result, RangeFromBound(r.RangeStart, overlapped.RangeStart))
	}

	if overlapped.RangeEnd < r.RangeEnd {
		result = append(result, RangeFromBound(overlapped.RangeEnd, r.RangeEnd))
	}
	return result
}

func FindMinFromRanges(ranges []*Range) uint {
	var min int = math.MaxInt
	for _, rg := range ranges {
		min = Min(min, rg.RangeStart)
	}

	return uint(min)
}

func Remove(arr []*Range, elt *Range) []*Range {
	var result []*Range

	for _, r := range arr {
		if r != elt {
			result = append(result, r)
		}
	}

	return result
}

type TransformationItem struct {
	Range         Range
	Shift         int
	SrcRangeName  string
	DestRangeName string
}

type TransformationRanges struct {
	Ranges []TransformationItem
}

func (tr *TransformationRanges) Transform(input int) (uint, string) {
	destRangeName := tr.Ranges[0].DestRangeName
	for _, mapsItem := range tr.Ranges {
		if input >= mapsItem.Range.RangeStart && input < mapsItem.Range.RangeEnd {
			return uint(input + mapsItem.Shift), destRangeName
		}
	}
	return uint(input), destRangeName
}
