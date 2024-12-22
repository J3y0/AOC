package days

import (
	"os"
	"strings"

	"aoc/utils"
)

type Day20 struct {
	grid           [][]rune
	start          utils.Pos
	maxRow, maxCol int
}

type cheatPos struct {
	dist int
	p    utils.Pos
}

func (d *Day20) Part1() (int, error) {
	return d.searchCheat(100, 2), nil
}

func (d *Day20) Part2() (int, error) {
	return d.searchCheat(100, 20), nil
}

func (d *Day20) Parse() error {
	content, err := os.ReadFile("./data/20_day.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	d.grid = make([][]rune, len(lines))
	for i, l := range lines {
		d.grid[i] = []rune(l)
		for j, r := range l {
			if r == 'S' {
				d.start = utils.Pos{X: i, Y: j}
			}
		}
	}

	d.maxRow = len(d.grid)
	d.maxCol = len(d.grid[0])
	return nil
}

func (d *Day20) initDist() map[utils.Pos]int {
	queue := []utils.Pos{d.start}
	dist := make(map[utils.Pos]int)
	dist[d.start] = 0
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if d.grid[p.X][p.Y] == 'E' {
			break
		}

		neighbors := d.getNeighbors(p)

		for _, n := range neighbors {
			if d.grid[n.X][n.Y] == '#' {
				continue
			}
			if _, ok := dist[n]; ok {
				continue
			}
			dist[n] = dist[p] + 1
			queue = append(queue, n)
		}
	}

	return dist
}

func (d *Day20) searchCheat(threshold int, cheatDuration int) int {
	tot := 0

	dist := d.initDist() // BFS for having base distance
	seen := make(map[utils.Pos]bool)
	seen[d.start] = true

	queue := []utils.Pos{d.start}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if d.grid[p.X][p.Y] == 'E' {
			continue
		}

		// compute time delta for valid positions after cheat
		afterCheatPos := d.explore(p, cheatDuration)
		for _, cp := range afterCheatPos {
			if dist[cp.p]-cp.dist-dist[p] >= threshold {
				tot++
			}
		}

		neighbors := d.getNeighbors(p)
		for _, n := range neighbors {
			if d.grid[n.X][n.Y] == '#' {
				continue
			}

			if seen[n] {
				continue
			}
			seen[n] = true
			queue = append(queue, n)
		}
	}

	return tot
}

func (d *Day20) getNeighbors(p utils.Pos) []utils.Pos {
	res := make([]utils.Pos, 0)

	neighbors := [4]utils.Pos{
		{X: p.X + 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y},
		{X: p.X, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
	}

	for _, n := range neighbors {
		if utils.OutOfGrid(n, d.maxRow, d.maxCol) {
			continue
		}
		res = append(res, n)
	}

	return res
}

func (d *Day20) explore(pos utils.Pos, width int) []cheatPos {
	res := make([]cheatPos, 0)
	for i := -width; i <= width; i++ {
		for j := -width; j <= width; j++ {
			if utils.Abs(i)+utils.Abs(j) > width {
				continue
			}
			n := utils.Pos{X: pos.X + i, Y: pos.Y + j}
			if utils.OutOfGrid(n, d.maxRow, d.maxCol) {
				continue
			}

			if d.grid[pos.X+i][pos.Y+j] != '#' {
				res = append(res, cheatPos{
					p:    n,
					dist: utils.Abs(i) + utils.Abs(j),
				})
			}
		}
	}
	return res
}
