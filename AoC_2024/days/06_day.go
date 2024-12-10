package days

import (
	"aoc/utils"
	"os"
	"strings"
)

type Day6 struct {
	start   utils.Pos
	grid    [][]rune
	visited map[utils.Pos]bool
}

func (d *Day6) Part1() (int, error) {
	count := 1
	directions := [4]utils.Pos{
		{X: -1, Y: 0},
		{X: 0, Y: 1},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
	}
	dir := 0
	pos := d.start

	d.visited = make(map[utils.Pos]bool)
	d.visited[d.start] = true
	for !utils.OutOfGrid(pos, len(d.grid), len(d.grid[0])) {
		// Check next wall
		if d.grid[pos.X][pos.Y] == '#' {
			pos.X = pos.X - directions[dir].X
			pos.Y = pos.Y - directions[dir].Y
			dir = (dir + 1) % len(directions)
			continue
		}

		if _, ok := d.visited[pos]; !ok {
			d.visited[pos] = true
			count++
		}
		// Update guard position
		pos.X = pos.X + directions[dir].X
		pos.Y = pos.Y + directions[dir].Y
	}
	return count, nil
}

func (d *Day6) Part2() (int, error) {
	if len(d.visited) == 0 {
		d.Part1()
	}

	totInfLoop := 0
	for pos := range d.visited {
		if pos == d.start || d.grid[pos.X][pos.Y] == '#' {
			continue
		}
		d.grid[pos.X][pos.Y] = '#'

		if d.isInfLoop() {
			totInfLoop++
		}

		d.grid[pos.X][pos.Y] = '.'
	}
	return totInfLoop, nil
}

func (d *Day6) Parse() error {
	content, err := os.ReadFile("./data/06_day.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	d.grid = make([][]rune, len(lines))
	for i, l := range lines {
		l = strings.TrimSpace(l)
		runes := make([]rune, len(l))
		for j, c := range l {
			if c == '^' {
				d.start = utils.Pos{X: i, Y: j}
			}
			runes[j] = c
		}
		d.grid[i] = runes
	}

	return nil
}

func (d *Day6) isInfLoop() bool {
	directions := [4]utils.Pos{
		{X: -1, Y: 0},
		{X: 0, Y: 1},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
	}
	dir := 0
	pos := d.start

	visited := make(map[[3]int]bool)
	for !utils.OutOfGrid(pos, len(d.grid), len(d.grid[0])) {
		key := [3]int{pos.X, pos.Y, dir}

		if visited[key] {
			return true
		}
		// Check next wall
		if d.grid[pos.X][pos.Y] == '#' {
			pos.X = pos.X - directions[dir].X
			pos.Y = pos.Y - directions[dir].Y
			dir = (dir + 1) % len(directions)
			continue
		}

		visited[key] = true
		// Update guard position
		pos.X = pos.X + directions[dir].X
		pos.Y = pos.Y + directions[dir].Y
	}

	return false
}
