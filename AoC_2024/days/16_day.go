package days

import (
	"math"
	"os"
	"strings"

	"aoc/utils"
)

type Day16 struct {
	grid  [][]rune
	start utils.Pos
}

type node struct {
	pos utils.Pos
	dir utils.Pos
}

type state struct {
	node  node
	score int
	path  []utils.Pos
}

func (d *Day16) Part1() (int, error) {
	return d.bfs(1), nil
}

func (d *Day16) Part2() (int, error) {
	return d.bfs(2), nil
}

func (d *Day16) Parse() error {
	content, err := os.ReadFile("./data/16_day.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	d.grid = make([][]rune, 0)
	for i, l := range lines {
		d.grid = append(d.grid, []rune(l))

		for j, c := range l {
			if c == 'S' {
				d.start = utils.Pos{X: i, Y: j}
			}
		}
	}

	return nil
}

func (d *Day16) getNeighbors(current state) (neighbors []node) {
	neighbors = make([]node, 0, 4)
	currDir, currPos := current.node.dir, current.node.pos
	oppositeDir := utils.Pos{X: -currDir.X, Y: -currDir.Y}

	directions := []utils.Pos{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
	}
	for _, dir := range directions {
		if dir == oppositeDir {
			continue
		}
		nPos := utils.Pos{X: currPos.X + dir.X, Y: currPos.Y + dir.Y}
		neighbors = append(neighbors, node{pos: nPos, dir: dir})
	}

	return
}

func (d *Day16) bfs(part int) int {
	queue := []state{
		{
			score: 0,
			node:  node{pos: d.start, dir: utils.Pos{X: 0, Y: 1}},
			path:  []utils.Pos{d.start},
		},
	}

	scores := make(map[node]int)
	minPaths := make(map[int][]utils.Pos)
	minScore := math.MaxInt
	for len(queue) > 0 {
		currState := queue[0]
		queue = queue[1:]

		if currState.score > minScore {
			continue
		}

		if d.grid[currState.node.pos.X][currState.node.pos.Y] == 'E' {
			if currState.score <= minScore {
				minScore = currState.score
				minPaths[minScore] = append(minPaths[minScore], currState.path...)
			}
			continue
		}

		neighbors := d.getNeighbors(currState)
		for _, n := range neighbors {
			if d.grid[n.pos.X][n.pos.Y] == '#' {
				continue
			}

			score := currState.score + 1
			if n.dir != currState.node.dir {
				score += 1000
			}

			previous, ok := scores[n]
			if ok && previous < score {
				continue
			}

			scores[n] = score

			nPath := make([]utils.Pos, len(currState.path))
			copy(nPath, currState.path)

			queue = append(queue, state{
				node:  n,
				score: score,
				path:  append(nPath, n.pos),
			})
		}
	}

	if part == 1 {
		return minScore
	}

	countSeats := make(map[utils.Pos]bool)
	for _, path := range minPaths[minScore] {
		countSeats[path] = true
	}
	return len(countSeats)
}
