package sensor

import (
	"sort"
)

type Point struct {
	X int
	Y int
}

type Sensor struct {
	Position              Point
	BeaconPosition        Point
	ClosestBeaconDistance int
}

type Line struct {
	Start int
	End   int
}

func (l *Line) Length() int {
	if l.Start == l.End {
		return 0
	}
	return l.End - l.Start + 1
}

func (s *Sensor) OmitBeacon(line Line) Line {
	// Associated beacon must be on the edge of the diamond (or wouln't
	// be the closest beacon -> impossible)
	if s.BeaconPosition.X >= s.Position.X {
		line.End = s.BeaconPosition.X - 1
	} else {
		line.Start = s.BeaconPosition.X + 1
	}
	return line
}

func (s *Sensor) Coverage(yCoord int) (line Line) {
	if (yCoord <= s.Position.Y-s.ClosestBeaconDistance) || (yCoord >= s.Position.Y+s.ClosestBeaconDistance) {
		// Not in range
		return
	}

	// Compute by how much we shift the x_min coord (s.Position.x - s.ClosestBeaconDistance) for being on yCoord
	// Because Manhattan distance remains equal when what is removed from x coord is added to y coord
	shift := Abs(yCoord - s.Position.Y)
	// Return the line which is on yCoord and in sensor's coverage
	line = Line{
		Start: s.Position.X - s.ClosestBeaconDistance + shift,
		End:   s.Position.X + s.ClosestBeaconDistance - shift,
	}

	if s.BeaconPosition.Y == yCoord {
		line = s.OmitBeacon(line)
	}
	return
}

type IntervalValue struct {
	value          int
	closingBracket bool // true for closing bracket
}

// Klee's algorithm
func ComputeLength(lines []Line) int {
	var result int
	var n int = len(lines)
	points := make([]IntervalValue, 2*n)

	for i := 0; i < n; i++ {
		points[i*2] = IntervalValue{value: lines[i].Start, closingBracket: false}
		points[i*2+1] = IntervalValue{value: lines[i].End, closingBracket: true}
	}

	sort.Slice(points, func(i, j int) bool {
		// If equal, should behave as if 2 segments are overlapping on this particular value
		if points[i].value == points[j].value {
			return !points[i].closingBracket
		}
		return points[i].value < points[j].value
	})

	nbOpenSegment := 0
	for i := 0; i < 2*n; i++ {
		if i > 0 && nbOpenSegment > 0 && points[i].value > points[i-1].value {
			result += points[i].value - points[i-1].value
		}

		if points[i].closingBracket {
			nbOpenSegment--
		} else {
			nbOpenSegment++
		}
		if nbOpenSegment == 0 {
			result++
		}
	}

	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Manhattan(p1 Point, p2 Point) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y)
}
