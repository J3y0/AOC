package days

import (
	"os"
	"strings"

	"aoc/utils"
)

type Day10 struct {
	maxRow, maxCol int
	grid           [][]rune
	trailheads     []utils.Pos
}

func (d *Day10) Part1() (int, error) {
	return d.FindPath(1)
}

func (d *Day10) Part2() (int, error) {
	return d.FindPath(2)
}

func (d *Day10) Parse() error {
	content, err := os.ReadFile("./data/10_day.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	d.grid = make([][]rune, len(lines))
	d.trailheads = make([]utils.Pos, 0)
	for i, l := range lines {
		d.grid[i] = []rune(l)
		// Find trailheads
		for j, r := range l {
			if r == '0' {
				d.trailheads = append(d.trailheads, utils.Pos{X: i, Y: j})
			}
		}
	}

	d.maxCol = len(d.grid[0])
	d.maxRow = len(d.grid)
	return nil
}

func (d *Day10) possibleNeighbors(pos, n utils.Pos) bool {
	val := int(d.grid[pos.X][pos.Y] - '0')
	valNeighb := int(d.grid[n.X][n.Y] - '0')
	if val+1 == valNeighb {
		return true
	}
	return false
}

func (d *Day10) FindPath(part int) (int, error) {
	tot := 0
	distinctPaths := 0
	for _, t := range d.trailheads {
		var trailtail map[utils.Pos]bool
		if part == 1 {
			trailtail = make(map[utils.Pos]bool)
		}
		queue := []utils.Pos{{X: t.X, Y: t.Y}}
		for len(queue) > 0 {
			// Pop
			pos := queue[0]
			queue = queue[1:]
			// If on a 9, stop the trail
			if d.grid[pos.X][pos.Y] == '9' {
				if part == 1 {
					trailtail[pos] = true
				} else {
					distinctPaths++
				}
				continue
			}

			// Possible neighbors
			neighbors := []utils.Pos{
				{X: pos.X - 1, Y: pos.Y},
				{X: pos.X + 1, Y: pos.Y},
				{X: pos.X, Y: pos.Y - 1},
				{X: pos.X, Y: pos.Y + 1},
			}
			for _, n := range neighbors {
				if !utils.OutOfGrid(n, d.maxRow, d.maxCol) && d.possibleNeighbors(pos, n) {
					queue = append(queue, n)
				}
			}
		}

		tot += len(trailtail)
	}

	if part == 2 {
		return distinctPaths, nil
	}
	return tot, nil
}
