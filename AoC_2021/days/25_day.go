package days

import (
	"main/utils"
)

type Day25 struct {
	grid   [][]rune
	height int
	width  int
}

func (d *Day25) Parse(input string) error {
	lines := utils.ParseLines(input)
	d.grid = make([][]rune, len(lines))
	for i, l := range lines {
		d.grid[i] = []rune(l)
	}
	d.height = len(lines)
	d.width = len(d.grid[0])
	return nil
}

func (d *Day25) Part1() (int, error) {
	nextGrid := make([][]rune, d.height)
	curGrid := make([][]rune, d.height)
	for i := range d.height {
		curGrid[i] = make([]rune, d.width)
		copy(curGrid[i], d.grid[i])
		nextGrid[i] = make([]rune, d.width)
		copy(nextGrid[i], d.grid[i])
	}

	step := 0
	for {
		hasMoved := false
		// move to east
		for r := range d.height {
			for c := range d.width {
				if curGrid[r][c] == '.' || curGrid[r][c] == 'v' || curGrid[r][(c+1)%d.width] != '.' {
					continue
				}
				nextGrid[r][c] = '.'
				nextGrid[r][(c+1)%d.width] = curGrid[r][c]
				hasMoved = true
			}
		}
		// update grid
		for i := range d.height {
			copy(curGrid[i], nextGrid[i])
		}
		// move to south
		for c := range d.width {
			for r := range d.height {
				if curGrid[r][c] == '.' || curGrid[r][c] == '>' || curGrid[(r+1)%d.height][c] != '.' {
					continue
				}
				nextGrid[r][c] = '.'
				nextGrid[(r+1)%d.height][c] = curGrid[r][c]
				hasMoved = true
			}
		}
		// update grid
		for i := range d.height {
			copy(curGrid[i], nextGrid[i])
		}
		step++
		if !hasMoved {
			break
		}
	}
	return step, nil
}

func (d *Day25) Part2() (int, error) {
	return 0, nil
}
