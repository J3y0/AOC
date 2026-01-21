package days

import (
	"bufio"
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

/*
 * CAUTION: start does not mean it has the lowest coordinates
 * If trying to, there is always the case of the antidiagonal where min x may be the start but not min y
 */
type Line struct {
	start, end Point
}

func NewLine(x1, y1, x2, y2 int) Line {
	p1 := Point{x: x1, y: y1}
	p2 := Point{x: x2, y: y2}

	return Line{
		start: p1,
		end:   p2,
	}
}

func (l *Line) IsHorizontal() bool {
	return l.start.x == l.end.x
}

func (l *Line) IsVertical() bool {
	return l.start.y == l.end.y
}

type Day5 struct {
	lines []Line
}

func (d *Day5) Parse(input string) error {
	s := bufio.NewScanner(strings.NewReader(input))
	lines := make([]Line, 0)
	for s.Scan() {
		var x1, y1, x2, y2 int
		line := s.Text()
		_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &y1, &x1, &y2, &x2)
		if err != nil {
			return err
		}
		lines = append(lines, NewLine(x1, y1, x2, y2))
	}
	d.lines = lines
	return nil
}

func (d *Day5) Part1() (int, error) {
	seen := make(map[Point]int)
	for _, l := range d.lines {
		if !l.IsHorizontal() && !l.IsVertical() {
			continue
		}

		dx := sign(l.end.x - l.start.x)
		dy := sign(l.end.y - l.start.y)
		x, y := l.start.x, l.start.y
		for {
			seen[Point{x, y}]++
			if x == l.end.x && y == l.end.y {
				break
			}
			x += dx
			y += dy
		}
	}

	return countOverlaps(seen), nil
}

func (d *Day5) Part2() (int, error) {
	seen := make(map[Point]int)
	for _, l := range d.lines {
		dx := sign(l.end.x - l.start.x)
		dy := sign(l.end.y - l.start.y)
		x, y := l.start.x, l.start.y
		for {
			seen[Point{x, y}]++
			if x == l.end.x && y == l.end.y {
				break
			}
			x += dx
			y += dy
		}
	}

	return countOverlaps(seen), nil
}

func sign(val int) int {
	switch {
	case val > 0:
		return 1
	case val < 0:
		return -1
	default:
		return 0
	}
}

func countOverlaps(seen map[Point]int) int {
	tot := 0
	for _, count := range seen {
		if count >= 2 {
			tot += 1
		}
	}

	return tot
}
