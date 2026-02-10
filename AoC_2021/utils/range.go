package utils

type Range struct {
	Start, End int
}

func (r *Range) Contains(v int) bool {
	return r.Start <= v && v <= r.End
}

func Intersect(r1, r2 Range) (Range, bool) {
	if r2.Start > r1.End || r2.End < r1.Start {
		return Range{}, false
	}
	interStart := max(r1.Start, r2.Start)
	interEnd := min(r1.End, r2.End)
	return Range{Start: interStart, End: interEnd}, true
}
