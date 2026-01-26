package days

import (
	"main/utils"
)

type Day11 struct {
	grid [][]int
}

func (d *Day11) Parse(input string) error {
	lines := utils.ParseLines(input)
	for _, l := range lines {
		parsed, err := utils.ParseLineToIntArray(l, "")
		if err != nil {
			return err
		}

		d.grid = append(d.grid, parsed)
	}
	return nil
}

func (d *Day11) Part1() (int, error) {
	localGrid := make([][]int, len(d.grid))
	for i := range d.grid {
		localGrid[i] = make([]int, len(d.grid[i]))
		copy(localGrid[i], d.grid[i])
	}

	countFlashes := 0
	for _ = range 100 {
		// Increase all
		flashQueue := make([][2]int, 0)
		for i := range localGrid {
			for j := range localGrid[i] {
				localGrid[i][j]++
				if localGrid[i][j] > 9 {
					flashQueue = append(flashQueue, [2]int{i, j})
				}
			}
		}

		flashmap := make([][]bool, len(localGrid))
		for i := range localGrid {
			flashmap[i] = make([]bool, len(localGrid[i]))
		}
		countFlashes += flashStep(flashQueue, localGrid, flashmap)
	}

	return countFlashes, nil
}

func (d *Day11) Part2() (int, error) {
	steps := 0
	for {
		// Increase all
		flashQueue := make([][2]int, 0)
		for i := range d.grid {
			for j := range d.grid[i] {
				d.grid[i][j]++
				if d.grid[i][j] > 9 {
					flashQueue = append(flashQueue, [2]int{i, j})
				}
			}
		}

		flashmap := make([][]bool, len(d.grid))
		for i := range d.grid {
			flashmap[i] = make([]bool, len(d.grid[i]))
		}
		countFlashes := flashStep(flashQueue, d.grid, flashmap)

		steps++
		// all jellyfishes flashed
		if countFlashes == len(d.grid)*len(d.grid[0]) {
			break
		}
	}

	return steps, nil
}

func flashStep(flashQueue [][2]int, grid [][]int, flashmap [][]bool) int {
	countFlashes := 0
	for {
		if len(flashQueue) == 0 {
			break
		}

		pi := flashQueue[0][0]
		pj := flashQueue[0][1]
		flashQueue = flashQueue[1:]

		if flashmap[pi][pj] {
			continue
		}
		flashmap[pi][pj] = true
		countFlashes++
		grid[pi][pj] = 0

		for _, n := range utils.Neighbors {
			ni := pi + n[0]
			nj := pj + n[1]
			if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) {
				continue
			}
			if flashmap[ni][nj] {
				continue
			}

			grid[ni][nj]++
			if grid[ni][nj] > 9 {
				flashQueue = append(flashQueue, [2]int{ni, nj})
			}
		}
	}

	return countFlashes
}
