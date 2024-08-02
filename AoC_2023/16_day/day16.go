package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Direction byte

const (
	UP Direction = 1 << iota
	RIGHT
	DOWN
	LEFT
)

type Tile struct {
	Char     rune
	BeamsDir Direction
}

func (t *Tile) IsEnergized() bool {
	return t.BeamsDir > 0
}

type ContraptionMap [][]Tile

func (cm ContraptionMap) clone() ContraptionMap {
	newContraption := make([][]Tile, len(cm))
	for i := range cm {
		newContraption[i] = slices.Clone(cm[i])
	}
	return newContraption
}

func (cm ContraptionMap) TraceBeam(row, col int, direction Direction) int {
	// If out of bound or already visited with the same beam
	if row >= len(cm) || row < 0 || col < 0 || col >= len(cm[0]) || (cm[row][col].BeamsDir&direction) > 0 {
		return 0
	}

	energized := 0
	if !cm[row][col].IsEnergized() {
		energized++
	}
	// Register that a beam traverses the tile with direction
	cm[row][col].BeamsDir |= direction

	switch cm[row][col].Char {
	case '/':
		switch direction {
		case UP:
			direction = RIGHT
		case DOWN:
			direction = LEFT
		case RIGHT:
			direction = UP
		case LEFT:
			direction = DOWN
		}
	case '\\':
		switch direction {
		case UP:
			direction = LEFT
		case DOWN:
			direction = RIGHT
		case RIGHT:
			direction = DOWN
		case LEFT:
			direction = UP
		}
	case '-':
		if direction == UP || direction == DOWN {
			energized += cm.TraceBeam(row, col-1, LEFT)
			energized += cm.TraceBeam(row, col+1, RIGHT)
			return energized
		}
	case '|':
		if direction == LEFT || direction == RIGHT {
			energized += cm.TraceBeam(row-1, col, UP)
			energized += cm.TraceBeam(row+1, col, DOWN)
			return energized
		}
	}

	switch direction {
	case UP:
		row--
	case DOWN:
		row++
	case RIGHT:
		col++
	case LEFT:
		col--
	default:
		panic("Unknown direction")
	}
	return energized + cm.TraceBeam(row, col, direction)
}

func (cm ContraptionMap) MaximizeEnergy() int {
	var maxEnergy int

	// Try for each lines
	energy := 0
	for i := range cm {
		energy = cm.clone().TraceBeam(i, 0, RIGHT)
		maxEnergy = max(maxEnergy, energy)

		energy = cm.clone().TraceBeam(i, len(cm[0])-1, LEFT)
		maxEnergy = max(maxEnergy, energy)
	}

	// Try for each column
	for j := range cm[0] {
		energy = cm.clone().TraceBeam(0, j, DOWN)
		maxEnergy = max(maxEnergy, energy)

		energy = cm.clone().TraceBeam(len(cm)-1, j, UP)
		maxEnergy = max(maxEnergy, energy)
	}

	return maxEnergy
}

func main() {
	contraptionMap, err := ReadMap("./data/day16.txt")
	if err != nil {
		panic(err)
	}

	tStart := time.Now()
	// ----- Part 1 -----
	part1 := contraptionMap.clone().TraceBeam(0, 0, RIGHT)
	fmt.Println("Result for part 1:", part1)
	fmt.Println("Time for part 1:", time.Since(tStart))

	// ----- Part 2 -----
	part2 := contraptionMap.MaximizeEnergy()
	fmt.Println("Result for part 2:", part2)
	fmt.Println("Time for part 2:", time.Since(tStart))
}

func ReadMap(path string) (ContraptionMap, error) {
	file, err := os.Open(path)
	if err != nil {
		return ContraptionMap{}, err
	}
	defer file.Close()

	var contraptionMap = make(ContraptionMap, 0)

	s := bufio.NewScanner(file)
	for s.Scan() {
		readLine := strings.TrimSpace(s.Text())

		var tiles = make([]Tile, len(readLine))
		for i, char := range readLine {
			tiles[i] = Tile{
				Char:     char,
				BeamsDir: Direction(0),
			}
		}
		contraptionMap = append(contraptionMap, tiles)
	}

	return contraptionMap, nil
}
