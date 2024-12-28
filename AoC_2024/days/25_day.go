package days

import (
	"fmt"
	"os"
	"strings"
)

type Day25 struct {
	keys  [][]int
	locks [][]int
}

func (d *Day25) Part1() (int, error) {
	pairs := 0
	for _, lock := range d.locks {
		for _, key := range d.keys {
			valid := true
			for i := range key {
				if lock[i]+key[i] > 5 {
					valid = false
					break
				}
			}

			if valid {
				pairs++
			}
		}
	}
	return pairs, nil
}

func (d *Day25) Part2() (int, error) {
	fmt.Println("No part 2 :)")
	return 0, nil
}

func (d *Day25) Parse() error {
	content, err := os.ReadFile("./data/25_day.txt")
	if err != nil {
		return err
	}

	blocks := strings.Split(strings.TrimSpace(string(content)), "\n\n")

	for _, b := range blocks {
		lines := strings.Split(b, "\n")
		h := make([]int, len(lines[0]))
		for col := range lines[0] {
			height := 0
			for row := range lines {
				if lines[row][col] == '#' {
					height++
				}
			}

			h[col] = height - 1
		}

		if lines[0][0] == '.' {
			d.keys = append(d.keys, h)
		} else {
			d.locks = append(d.locks, h)
		}
	}
	return nil
}
