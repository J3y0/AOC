package days

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"aoc/utils"
)

type Day5 struct {
	rules   map[string]bool
	updates [][]int
}

// if a b match a|b rule, they are in the right order
func (d *Day5) compare(a int, b int) int {
	if d.rules[fmt.Sprintf("%d|%d", a, b)] {
		return -1
	}
	return 1
}

func (d *Day5) Part1() (int, error) {
	sum := 0
	for _, upd := range d.updates {
		if slices.IsSortedFunc(upd, d.compare) {
			sum += upd[len(upd)/2]
		}
	}
	return sum, nil
}

func (d *Day5) Part2() (int, error) {
	sum := 0
	for _, upd := range d.updates {
		if !slices.IsSortedFunc(upd, d.compare) {
			slices.SortFunc(upd, d.compare)
			sum += upd[len(upd)/2]
		}
	}
	return sum, nil
}

func (d *Day5) Parse() error {
	file, err := os.Open("./data/05_day.txt")
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error while closing file")
		}
	}(file)

	s := bufio.NewScanner(file)
	var line string
	// Read rules
	d.rules = make(map[string]bool)
	for s.Scan() {
		line = strings.TrimSpace(s.Text())
		if line == "" {
			break
		}

		d.rules[line] = true
	}

	// Read updates
	d.updates = make([][]int, 0)
	for s.Scan() {
		line = s.Text()
		upd, err := utils.FromLine(line, ",")
		if err != nil {
			return err
		}

		d.updates = append(d.updates, upd)
	}
	return nil
}
