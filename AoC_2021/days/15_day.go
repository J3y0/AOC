package days

import (
	"container/heap"
	"main/utils"
)

type Day15 struct {
	grid   [][]rune
	height int
	width  int
}

func (d *Day15) Parse(input string) error {
	lines := utils.ParseLines(input)
	d.height = len(lines)
	d.width = len(lines[0])
	d.grid = make([][]rune, d.height)

	for i, l := range lines {
		d.grid[i] = []rune(l)
	}
	return nil
}

func (d *Day15) Part1() (int, error) {
	return djikstra(d.grid, d.height, d.width, 1, [2]int{0, 0}, [2]int{d.height - 1, d.width - 1}), nil
}

func (d *Day15) Part2() (int, error) {
	return djikstra(d.grid, d.height, d.width, 5, [2]int{0, 0}, [2]int{5*d.height - 1, 5*d.width - 1}), nil
}

// h: height of the default grid
// w: width of the default grid
// m: number of times the grid is repeated
// start: start position
// end: end position
func djikstra(grid [][]rune, h, w, m int, start, end [2]int) int {
	seen := make(map[[2]int]bool)
	risks := make(map[[2]int]int)
	pq := &utils.PriorityQueue{
		&utils.PriorityQueueItem{
			Pos:      start,
			Priority: 0,
		},
	}
	heap.Init(pq)

	risks[start] = 0
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*utils.PriorityQueueItem)
		p := item.Pos

		if seen[p] {
			continue
		}
		seen[p] = true

		// end
		if p == end {
			return risks[p]
		}

		for _, n := range utils.OrthogonalNeighbors {
			nx := p[0] + n[0]
			ny := p[1] + n[1]
			// out of grid
			if nx < 0 || nx >= m*h || ny < 0 || ny >= m*w {
				continue
			}
			npos := [2]int{nx, ny}
			if seen[npos] {
				continue
			}
			// update score if better than known
			oldRisk, ok := risks[npos]
			newRisk := risks[p] + getRiskLevel(grid, h, w, nx, ny)
			if !ok || newRisk < oldRisk {
				risks[npos] = newRisk
				heap.Push(pq, &utils.PriorityQueueItem{
					Pos:      npos,
					Priority: newRisk,
				})
			}
		}
	}

	return -1
}

func getRiskLevel(grid [][]rune, h, w, x, y int) int {
	val := int(grid[x%h][y%w] - '0')
	return (val-1+x/h+y/w)%9 + 1
}
