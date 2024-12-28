package days

import (
	"fmt"
	"os"
	"strings"

	"aoc/utils"
)

type Day15 struct {
	maxRow, maxCol int
	grid           [][]rune
	moves          []rune
	start          utils.Pos
}

func (d *Day15) Part1() (int, error) {
	pos := d.start
	grid := utils.CopyArr(d.grid)
	for _, m := range d.moves {
		switch m {
		case 'v':
			pos = d.move(pos, utils.Pos{X: 1, Y: 0}, grid)
		case '>':
			pos = d.move(pos, utils.Pos{X: 0, Y: 1}, grid)
		case '<':
			pos = d.move(pos, utils.Pos{X: 0, Y: -1}, grid)
		case '^':
			pos = d.move(pos, utils.Pos{X: -1, Y: 0}, grid)
		default:
			return 0, fmt.Errorf("Undefined move: %c", m)
		}
	}

	// Compute answer
	tot := 0
	for i := range d.maxRow {
		for j := range d.maxCol {
			if grid[i][j] == 'O' {
				tot += 100*i + j
			}
		}
	}
	return tot, nil
}

func (d *Day15) Part2() (int, error) {
	newGrid := d.expand(d.grid)

	pos := utils.Pos{X: d.start.X, Y: 2 * d.start.Y}
	for _, m := range d.moves {
		switch m {
		case 'v':
			pos = d.movePart2(pos, utils.Pos{X: 1, Y: 0}, newGrid)
		case '>':
			pos = d.movePart2(pos, utils.Pos{X: 0, Y: 1}, newGrid)
		case '<':
			pos = d.movePart2(pos, utils.Pos{X: 0, Y: -1}, newGrid)
		case '^':
			pos = d.movePart2(pos, utils.Pos{X: -1, Y: 0}, newGrid)
		default:
			return 0, fmt.Errorf("Undefined move: %c", m)
		}
	}

	// Compute answer
	tot := 0
	for i := range len(newGrid) {
		for j := range len(newGrid[0]) {
			if newGrid[i][j] == '[' {
				tot += 100*i + j
			}
		}
	}
	return tot, nil
}

func (d *Day15) Parse() error {
	content, err := os.ReadFile("./data/15_day.txt")
	if err != nil {
		return err
	}

	split := strings.Split(strings.TrimSpace(string(content)), "\n\n")
	lines := strings.Split(split[0], "\n")
	moves := split[1]
	d.moves = []rune(strings.ReplaceAll(moves, "\n", ""))

	d.grid = make([][]rune, len(lines))
	for i, line := range lines {
		d.grid[i] = []rune(line)
		for j := range line {
			if line[j] == '@' {
				d.start = utils.Pos{X: i, Y: j}
			}
		}
	}

	d.maxRow = len(lines)
	d.maxCol = len(lines[0])
	return nil
}

func (d *Day15) move(pos, dir utils.Pos, grid [][]rune) utils.Pos {
	newPos := utils.Pos{X: pos.X + dir.X, Y: pos.Y + dir.Y}
	if grid[newPos.X][newPos.Y] == '#' {
		return pos
	}

	// rocks
	if grid[newPos.X][newPos.Y] == 'O' {
		i := 0
		for grid[newPos.X+i*dir.X][newPos.Y+i*dir.Y] == 'O' {
			i++
		}
		// Can't move as wall behind
		if grid[newPos.X+i*dir.X][newPos.Y+i*dir.Y] == '#' {
			return pos
		}
		// All good, just move first rock to last pos
		grid[newPos.X][newPos.Y] = '.'
		grid[newPos.X+i*dir.X][newPos.Y+i*dir.Y] = 'O'
	}

	return newPos
}

func (d *Day15) movePart2(pos, dir utils.Pos, grid [][]rune) utils.Pos {
	seen := make(map[utils.Pos]bool)
	blocks := []utils.Pos{pos}

	canMove := true
	for len(blocks) > 0 {
		p := blocks[0]
		blocks = blocks[1:]

		newPos := utils.Pos{X: p.X + dir.X, Y: p.Y + dir.Y}
		if seen[p] {
			continue
		}
		seen[p] = true
		val := grid[newPos.X][newPos.Y]
		if val == '#' {
			canMove = false
			break
		}

		if val == '[' {
			blocks = append(blocks, utils.Pos{X: newPos.X, Y: newPos.Y})
			blocks = append(blocks, utils.Pos{X: newPos.X, Y: newPos.Y + 1})
		}

		if val == ']' {
			blocks = append(blocks, utils.Pos{X: newPos.X, Y: newPos.Y - 1})
			blocks = append(blocks, utils.Pos{X: newPos.X, Y: newPos.Y})
		}
	}

	if !canMove {
		return pos
	}

	gcopy := utils.CopyArr(grid)
	for k := range seen {
		grid[k.X][k.Y] = '.'
	}

	for k := range seen {
		grid[k.X+dir.X][k.Y+dir.Y] = gcopy[k.X][k.Y]
	}

	return utils.Pos{X: pos.X + dir.X, Y: pos.Y + dir.Y}
}

func (d *Day15) expand(grid [][]rune) [][]rune {
	newGrid := make([][]rune, d.maxRow)

	for i := range d.maxRow {
		line := make([]rune, 2*d.maxCol)
		for j := range d.maxCol {
			if grid[i][j] == '#' {
				line[2*j] = '#'
				line[2*j+1] = '#'
			}

			if grid[i][j] == '.' {
				line[2*j] = '.'
				line[2*j+1] = '.'
			}

			if grid[i][j] == 'O' {
				line[2*j] = '['
				line[2*j+1] = ']'
			}

			if grid[i][j] == '@' {
				line[2*j] = '@'
				line[2*j+1] = '.'
			}
		}
		newGrid[i] = line
	}
	return newGrid
}

func printGrid(grid [][]rune) {
	for i := range len(grid) {
		for j := range len(grid[0]) {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println()
	}
}
