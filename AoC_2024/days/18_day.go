package days

import (
	"fmt"
	"math"
	"os"
	"strings"

	"aoc/utils"
)

type Day18 struct {
	bytes          []utils.Pos
	grid           [][]rune
	maxRow, maxCol int
}

func (d *Day18) Part1() (int, error) {
	grid := utils.CopyArr(d.grid)
	for i := range 1024 {
		grid[d.bytes[i].X][d.bytes[i].Y] = '#'
	}

	return d.findShortestPath(grid), nil
}

func (d *Day18) Part2() (int, error) {
	minByte := 1024
	maxByte := len(d.bytes)
	for i := range minByte {
		d.grid[d.bytes[i].X][d.bytes[i].Y] = '#'
	}

	fill := true
	half := 0
	for utils.Abs(minByte-maxByte) > 1 {
		half = (maxByte + minByte) / 2
		if fill {
			for i := minByte; i <= half; i++ {
				d.grid[d.bytes[i].X][d.bytes[i].Y] = '#'
			}
		} else {
			for i := maxByte; i >= half; i-- {
				d.grid[d.bytes[i].X][d.bytes[i].Y] = '.'
			}
		}

		lenPath := d.findShortestPath(d.grid)
		if lenPath == math.MaxInt {
			fill = false
			maxByte = half
		} else {
			fill = true
			minByte = half
		}
	}

	part2 := fmt.Sprintf("%d,%d", d.bytes[half].Y, d.bytes[half].X)
	fmt.Println("Part2 answer (not a number):", part2)
	return 0, nil
}

func (d *Day18) Parse() error {
	content, err := os.ReadFile("./data/18_day.txt")
	if err != nil {
		return err
	}

	d.maxRow = 70 + 1
	d.maxCol = 70 + 1
	d.grid = make([][]rune, d.maxRow)
	for i := range d.maxRow {
		l := make([]rune, d.maxCol)
		for j := range d.maxCol {
			l[j] = '.'
		}
		d.grid[i] = l
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	d.bytes = make([]utils.Pos, len(lines))
	for i, l := range lines {
		var x, y int
		_, err := fmt.Sscanf(l, "%d,%d", &y, &x)
		if err != nil {
			return err
		}

		d.bytes[i] = utils.Pos{X: x, Y: y}
	}
	return nil
}

func (d *Day18) findShortestPath(grid [][]rune) int {
	queue := []utils.Pos{{X: 0, Y: 0}}
	scores := make(map[utils.Pos]int)
	seen := make(map[utils.Pos]bool)
	minScore := math.MaxInt
	seen[utils.Pos{X: 0, Y: 0}] = true
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if scores[p] > minScore {
			continue
		}

		// End
		if p.X == d.maxRow-1 && p.Y == d.maxCol-1 {
			if minScore > scores[p] {
				minScore = scores[p]
			}
			continue
		}

		neighbors := [4]utils.Pos{
			{X: p.X + 1, Y: p.Y},
			{X: p.X - 1, Y: p.Y},
			{X: p.X, Y: p.Y + 1},
			{X: p.X, Y: p.Y - 1},
		}

		for _, n := range neighbors {
			if utils.OutOfGrid(n, d.maxRow, d.maxCol) || grid[n.X][n.Y] == '#' {
				continue
			}

			if seen[n] {
				continue
			}
			seen[n] = true

			scoreN, ok := scores[n]
			if ok && scoreN < scores[p]+1 {
				continue
			}
			scores[n] = scores[p] + 1
			queue = append(queue, n)
		}
	}

	return minScore
}
