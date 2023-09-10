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

func coefCurves(sx1, sy1, sx2, sy2, r1, r2 int) (int, int, int, int) {
	// Curves of the form x+a or -x+b forming the bounds of the diamond
	a1 := -(sx1 + r1) + sy1
	a3 := -(sx1 - r1) + sy1
	b2 := sx2 + sy2 - r2
	b4 := sx2 + sy2 + r2
	return a1, a3, b2, b4
}

func AddIntersectionPoints(s1 Sensor, s2 Sensor, allIntersectionPoints []Point) []Point {
	// There is 8 intersection points possible between 2 sensors
	r1 := s1.ClosestBeaconDistance + 1
	r2 := s2.ClosestBeaconDistance + 1
	sx1 := s1.Position.X
	sy1 := s1.Position.Y
	sx2 := s2.Position.X
	sy2 := s2.Position.Y
	// Constants are named as follows: a for positive gradient, b for negative
	// first number is the number of the curve (imagine a diamond, the first curve would be the top-right one)
	// second number is the number of the sensor used to define the point: either the first one or the second one
	// we need to change sensors' order as values differ
	a11, a31, b22, b42 := coefCurves(sx1, sy1, sx2, sy2, r1, r2)
	a12, a32, b21, b41 := coefCurves(sx2, sy2, sx1, sy1, r2, r1)

	// Points are named as follows:
	// point_12 is the intersection point between curve 1 and 2
	point_12 := Point{X: (b22 - a11) / 2, Y: (b22 + a11) / 2}
	point_21 := Point{X: (b21 - a12) / 2, Y: (b21 + a12) / 2}

	point_14 := Point{X: (b42 - a11) / 2, Y: (b42 + a11) / 2}
	point_41 := Point{X: (b41 - a12) / 2, Y: (b41 + a12) / 2}

	point_32 := Point{X: (b22 - a31) / 2, Y: (b22 + a31) / 2}
	point_23 := Point{X: (b21 - a32) / 2, Y: (b21 + a32) / 2}

	point_34 := Point{X: (b42 - a31) / 2, Y: (b42 + a31) / 2}
	point_43 := Point{X: (b41 - a32) / 2, Y: (b41 + a32) / 2}
	var allPoints []Point = []Point{point_12, point_21, point_14, point_41, point_32, point_23, point_34, point_43}

	allPoints = filterBounds(allPoints)

	allIntersectionPoints = append(allIntersectionPoints, allPoints...)
	return allIntersectionPoints
}

func isWithinBounds(p Point) bool {
	return p.X >= 0 && p.X <= 4000000 && p.Y >= 0 && p.Y <= 4000000
}

func filterBounds(allPoints []Point) []Point {
	var result []Point
	for _, p := range allPoints {
		if isWithinBounds(p) {
			result = append(result, p)
		}
	}

	return result
}

func removeDuplicate(points []Point) (result []Point) {
	if len(points) < 1 {
		result = points
		return
	}
	allKeys := make(map[Point]bool)
	for _, p := range points {
		if _, added := allKeys[p]; !added {
			result = append(result, p)
			allKeys[p] = true
		}
	}

	return
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
