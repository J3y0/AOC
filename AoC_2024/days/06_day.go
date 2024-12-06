package days

import (
	"os"
	"strings"
)

type Day6 struct {
	start   [2]int
	grid    [][]rune
	visited map[[2]int]bool
}

func (d *Day6) Part1() (int, error) {
	count := 1
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	dir := 0
	pos := d.start

	d.visited = make(map[[2]int]bool)
	d.visited[d.start] = true
	for d.validPos(pos) {
		// Check next wall
		if d.grid[pos[0]][pos[1]] == '#' {
			pos[0] = pos[0] - directions[dir][0]
			pos[1] = pos[1] - directions[dir][1]
			dir = (dir + 1) % len(directions)
			continue
		}

		if _, ok := d.visited[pos]; !ok {
			d.visited[pos] = true
			count++
		}
		// Update guard position
		pos[0] = pos[0] + directions[dir][0]
		pos[1] = pos[1] + directions[dir][1]
	}
	return count, nil
}

func (d *Day6) Part2() (int, error) {
	if len(d.visited) == 0 {
		d.Part1()
	}

	totInfLoop := 0
	for pos := range d.visited {
		if pos == d.start || d.grid[pos[0]][pos[1]] == '#' {
			continue
		}
		d.grid[pos[0]][pos[1]] = '#'

		if d.isInfLoop() {
			totInfLoop++
		}

		d.grid[pos[0]][pos[1]] = '.'
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
				d.start = [2]int{i, j}
			}
			runes[j] = c
		}
		d.grid[i] = runes
	}

	return nil
}

func (d *Day6) validPos(pos [2]int) bool {
	return pos[0] < len(d.grid) && pos[1] < len(d.grid[0]) && pos[0] >= 0 && pos[1] >= 0
}

func (d *Day6) isInfLoop() bool {
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	dir := 0
	pos := d.start

	visited := make(map[[3]int]bool)
	for d.validPos(pos) {
		key := [3]int{pos[0], pos[1], dir}

		if visited[key] {
			return true
		}
		// Check next wall
		if d.grid[pos[0]][pos[1]] == '#' {
			pos[0] = pos[0] - directions[dir][0]
			pos[1] = pos[1] - directions[dir][1]
			dir = (dir + 1) % len(directions)
			continue
		}

		visited[key] = true
		// Update guard position
		pos[0] = pos[0] + directions[dir][0]
		pos[1] = pos[1] + directions[dir][1]
	}

	return false
}
