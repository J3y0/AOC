package days

import (
	"os"
	"strings"

	"aoc/utils"
)

type Day8 struct {
	antennas       map[byte][]utils.Pos
	maxRow, maxCol int
}

func (d *Day8) Part1() (int, error) {
	antinodes := 0
	seen := make(map[utils.Pos]bool)
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
	seen := make(map[utils.Pos]bool)
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
	d.antennas = make(map[byte][]utils.Pos)
	for i, l := range lines {
		for j := range l {
			c := lines[i][j]
			if c == '.' {
				continue
			}

			if _, ok := d.antennas[c]; ok {
				d.antennas[c] = append(d.antennas[c], utils.Pos{X: i, Y: j})
			} else {
				d.antennas[c] = make([]utils.Pos, 1)
				d.antennas[c][0] = utils.Pos{X: i, Y: j}
			}
		}
	}
	return nil
}

func (d *Day8) countAntinodes(ant1, ant2 utils.Pos, seen map[utils.Pos]bool, part int) (valid int) {
	dy := ant2.Y - ant1.Y
	dx := ant2.X - ant1.X

	i := part % 2 // Need 0 for part2 and 1 for part1
	for {
		anti1 := utils.Pos{X: ant1.X - i*dx, Y: ant1.Y - i*dy}
		anti2 := utils.Pos{X: ant2.X + i*dx, Y: ant2.Y + i*dy}

		anti1Out := utils.OutOfGrid(anti1, d.maxRow, d.maxCol)
		anti2Out := utils.OutOfGrid(anti2, d.maxRow, d.maxCol)

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
