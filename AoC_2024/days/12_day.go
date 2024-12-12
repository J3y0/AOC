package days

import (
	"aoc/utils"
	"os"
	"strings"
)

type Day12 struct {
	grid           [][]rune
	maxRow, maxCol int
}

func (d *Day12) Part1() (int, error) {
	seen := make(map[utils.Pos]bool)
	tot := 0
	for i := range len(d.grid) {
		for j := range len(d.grid[0]) {
			if seen[utils.Pos{X: i, Y: j}] {
				continue
			}
			a, p := d.findAreaPerimeter(utils.Pos{X: i, Y: j}, seen)
			tot += a * p
		}
	}
	return tot, nil
}

func (d *Day12) Part2() (int, error) {
	seen := make(map[utils.Pos]bool)
	tot := 0
	for i := range len(d.grid) {
		for j := range len(d.grid[0]) {
			if seen[utils.Pos{X: i, Y: j}] {
				continue
			}
			a, s := d.findAreaSides(utils.Pos{X: i, Y: j}, seen)
			tot += a * s
		}
	}
	return tot, nil
}

func (d *Day12) Parse() error {
	content, err := os.ReadFile("./data/12_day.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	d.grid = make([][]rune, len(lines))
	for i, l := range lines {
		d.grid[i] = []rune(l)
	}

	d.maxRow = len(lines)
	d.maxCol = len(lines[0])
	return nil
}

func (d *Day12) findAreaPerimeter(start utils.Pos, seen map[utils.Pos]bool) (area int, perimeter int) {
	queue := []utils.Pos{start}
	seen[start] = true
	area = 1
	for len(queue) > 0 {
		// pop
		pos := queue[0]
		queue = queue[1:]

		neighbors := [4]utils.Pos{
			{X: pos.X + 1, Y: pos.Y},
			{X: pos.X - 1, Y: pos.Y},
			{X: pos.X, Y: pos.Y + 1},
			{X: pos.X, Y: pos.Y - 1},
		}

		for _, n := range neighbors {
			// Continue if not valid neighbor
			if utils.OutOfGrid(n, d.maxRow, d.maxCol) || d.grid[n.X][n.Y] != d.grid[pos.X][pos.Y] {
				/*
					If no valid neighbor, add 1 to perimeter. Consider this example:
						 _ _ _
						|a a a|  every bar is an invalid neighbor
						 - - -
				*/
				perimeter++
				continue
			}

			if !seen[n] {
				queue = append(queue, n)
				seen[n] = true
				// Increase region area
				area++
			}
		}
	}
	return
}

func (d *Day12) findAreaSides(start utils.Pos, seen map[utils.Pos]bool) (area int, sides int) {
	queue := []utils.Pos{start}
	seen[start] = true
	area = 1

	/*
			edge directions
			Take a region:
		              0 _
					1 |A| 3
		             2 -
	*/
	const (
		up = iota
		left
		down
		right
	)

	edges := make(map[[3]int]bool)
	for len(queue) > 0 {
		// pop
		pos := queue[0]
		queue = queue[1:]

		neighbors := [4]utils.Pos{
			{X: pos.X + 1, Y: pos.Y},
			{X: pos.X - 1, Y: pos.Y},
			{X: pos.X, Y: pos.Y + 1},
			{X: pos.X, Y: pos.Y - 1},
		}

		for _, n := range neighbors {
			if utils.OutOfGrid(n, d.maxRow, d.maxCol) || d.grid[n.X][n.Y] != d.grid[pos.X][pos.Y] {
				if n.X < pos.X {
					edges[[3]int{pos.X, pos.Y, up}] = true
				} else if n.X > pos.X {
					edges[[3]int{pos.X, pos.Y, down}] = true
				} else if n.Y > pos.Y {
					edges[[3]int{pos.X, pos.Y, right}] = true
				} else {
					edges[[3]int{pos.X, pos.Y, left}] = true
				}
				continue
			}

			if !seen[n] {
				queue = append(queue, n)
				seen[n] = true
				// Increase region area
				area++
			}
		}
	}

	// Return early
	if area == 1 {
		sides = 4
		return
	}
	sides = d.countSides(edges)

	return
}

func (d *Day12) countSides(edges map[[3]int]bool) (sides int) {
	// Idea is to count as seen edges that share a direction with the edge we examine
	// Then, we count one for the set of edges and not cardinal(set)
	edgesSeen := make(map[[3]int]bool)
	for k := range edges {
		if edgesSeen[k] {
			continue
		}
		edgesSeen[k] = true
		sides++
		px, py, dir := k[0], k[1], k[2]
		if dir%2 == 0 {
			// Search horizontally
			for _, d := range [2]int{-1, 1} {
				newPosY := py + d
				for {
					next := [3]int{px, newPosY, dir}
					if !edges[next] {
						break
					}
					edgesSeen[next] = true
					newPosY += d
				}
			}
		} else {
			// Search vertically
			for _, d := range [2]int{-1, 1} {
				newPosX := px + d
				for {
					next := [3]int{newPosX, py, dir}
					if !edges[next] {
						break
					}
					edgesSeen[next] = true
					newPosX += d
				}
			}
		}
	}
	return
}
