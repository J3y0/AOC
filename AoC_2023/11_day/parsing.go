package main

import (
	"slices"
)

type GalaxyPos struct {
	X int
	Y int
}

func ParseFile(path string, part int) (galaxies []GalaxyPos, err error) {
	lines, err := ReadLines(path)
	if err != nil {
		return
	}

	var times int
	if part == 1 {
		times = 1
	} else if part == 2 {
		times = 1000000 - 1
	}

	galaxies = ParseGalaxies(lines)
	ExpandSpace(lines, galaxies, times)

	return
}

func ParseGalaxies(lines []string) []GalaxyPos {
	var galaxies []GalaxyPos
	for i, line := range lines {
		for j := range line {
			if string(line[j]) == "#" {
				galaxies = append(galaxies, GalaxyPos{X: i, Y: j})
			}
		}
	}
	return galaxies
}

func ExpandSpace(lines []string, galaxies []GalaxyPos, times int) {
	// Sort according to X coordinate
	slices.SortStableFunc(galaxies, func(a, b GalaxyPos) int {
		return a.X - b.X
	})

	// Expand according to X coordinate (line)
	var (
		nbExpandedLine   int
		nbExpandedColumn int
		galIdx           int
	)
	for _, line := range lines {
		nbOnLine := NbGalaxyOnLine(line)
		if nbOnLine == 0 {
			nbExpandedLine += times
			continue
		}

		for k := 0; k < nbOnLine; k++ {
			galaxies[galIdx].X += nbExpandedLine
			galIdx++
		}
	}

	// Sort according to Y coordinate
	slices.SortStableFunc(galaxies, func(a, b GalaxyPos) int { return a.Y - b.Y })
	galIdx = 0

	// Expand according to Y coordinate (column)
	for j := 0; j < len(lines[0]); j++ {
		nbOnColumn := NbGalaxyOnColumn(lines, j)
		if nbOnColumn == 0 {
			nbExpandedColumn += times
			continue
		}

		for k := 0; k < nbOnColumn; k++ {
			galaxies[galIdx].Y += nbExpandedColumn
			galIdx++
		}
	}
}

func NbGalaxyOnLine(line string) int {
	var counter int
	for _, char := range line {
		if string(char) == "#" {
			counter++
		}
	}
	return counter
}

func NbGalaxyOnColumn(lines []string, colIdx int) int {
	var counter int
	for _, line := range lines {
		if string(line[colIdx]) == "#" {
			counter++
		}
	}
	return counter
}
