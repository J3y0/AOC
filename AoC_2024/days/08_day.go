package days

import (
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

type Day8 struct {
	antennas       map[byte][]Pos
	maxRow, maxCol int
}

func (d *Day8) Part1() (int, error) {
	antinodes := 0
	seen := make(map[Pos]bool)
	for _, antennas := range d.antennas {
		for i := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				antinodes += d.countAntinodes(antennas[i], antennas[j], seen, 1)
			}
		}
	}
	return antinodes, nil
}

func (d *Day8) Part2() (int, error) {
	antinodes := 0
	seen := make(map[Pos]bool)
	for _, antennas := range d.antennas {
		for i := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				antinodes += d.countAntinodes(antennas[i], antennas[j], seen, 2)
			}
		}
	}
	return antinodes, nil
}

func (d *Day8) Parse() error {
	content, err := os.ReadFile("./data/08_day.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	d.maxRow = len(lines)
	d.maxCol = len(lines[0])
	d.antennas = make(map[byte][]Pos)
	for i, l := range lines {
		for j := range l {
			c := lines[i][j]
			if c == '.' {
				continue
			}

			if _, ok := d.antennas[c]; ok {
				d.antennas[c] = append(d.antennas[c], Pos{x: i, y: j})
			} else {
				d.antennas[c] = make([]Pos, 1)
				d.antennas[c][0] = Pos{x: i, y: j}
			}
		}
	}
	return nil
}

func (d *Day8) countAntinodes(ant1, ant2 Pos, seen map[Pos]bool, part int) (valid int) {
	dy := ant2.y - ant1.y
	dx := ant2.x - ant1.x

	i := part % 2 // Need 0 for part2 and 1 for part1
	for {
		anti1 := Pos{x: ant1.x - i*dx, y: ant1.y - i*dy}
		anti2 := Pos{x: ant2.x + i*dx, y: ant2.y + i*dy}

		anti1Out := d.isOut(anti1)
		anti2Out := d.isOut(anti2)

		if anti1Out && anti2Out {
			break
		}

		if !anti1Out && !seen[anti1] {
			seen[anti1] = true
			valid++
		}

		if !anti2Out && !seen[anti2] {
			seen[anti2] = true
			valid++
		}

		if part == 1 {
			break
		}

		i++
	}
	return
}

func (d *Day8) isOut(p Pos) bool {
	return p.x < 0 || p.y < 0 || p.x >= d.maxRow || p.y >= d.maxCol
}
