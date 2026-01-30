package days

import (
	"main/utils"
	"slices"
)

type Day9 struct {
	heightmaps [][]rune
}

func (d *Day9) Parse(input string) error {
	lines := utils.ParseLines(input)
	d.heightmaps = make([][]rune, len(lines))
	for i, l := range lines {
		d.heightmaps[i] = []rune(l)
	}
	return nil
}

func (d *Day9) Part1() (int, error) {
	n, m := len(d.heightmaps), len(d.heightmaps[0])
	tot := 0
	for i := range n {
		for j := range m {
			if isLowPoint(i, j, d.heightmaps) {
				tot += int(d.heightmaps[i][j]-'0') + 1
			}
		}
	}
	return tot, nil
}

func (d *Day9) Part2() (int, error) {
	// get all start points
	startPoints := make([][2]int, 0)
	n, m := len(d.heightmaps), len(d.heightmaps[0])
	for i := range n {
		for j := range m {
			if isLowPoint(i, j, d.heightmaps) {
				startPoints = append(startPoints, [2]int{i, j})
			}
		}
	}

	sizes := make([]int, len(startPoints))
	for i, s := range startPoints {
		sizes[i] = getBasinSize(s, d.heightmaps)
	}

	// sort in descending order
	slices.SortFunc(sizes, func(i, j int) int {
		return j - i
	})

	return sizes[0] * sizes[1] * sizes[2], nil
}

func isLowPoint(i, j int, hmap [][]rune) bool {
	tile := hmap[i][j]
	ok := true
	if i > 0 {
		ok = ok && tile < hmap[i-1][j]
	}

	if j > 0 {
		ok = ok && tile < hmap[i][j-1]
	}

	if i < len(hmap)-1 {
		ok = ok && tile < hmap[i+1][j]
	}

	if j < len(hmap[i])-1 {
		ok = ok && tile < hmap[i][j+1]
	}

	return ok
}

func getBasinSize(start [2]int, hmap [][]rune) int {
	q := [][2]int{start}
	area := 0
	seen := make(map[[2]int]bool)
	// Explore neighbors if they are not equal to 9
	for {
		if len(q) == 0 {
			break
		}
		// unqueue
		p := q[0]
		q = q[1:]
		if seen[p] {
			continue
		}
		seen[p] = true
		area += 1

		// Add neighbors in the queue
		for _, n := range utils.OrthogonalNeighbors {
			ni := p[0] + n[0]
			nj := p[1] + n[1]
			if ni < 0 || nj < 0 || ni >= len(hmap) || nj >= len(hmap[0]) {
				continue
			}
			if hmap[ni][nj] == '9' {
				continue
			}
			q = append(q, [2]int{ni, nj})
		}
	}

	return area
}
