package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"aoc/utils"
)

type Day2 struct {
	reports [][]int
}

func (d *Day2) Parse() error {
	file, err := os.Open("./data/02_day.txt")
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
	d.reports = make([][]int, 0)
	for s.Scan() {
		line := s.Text()
		line = strings.TrimSpace(line)

		rep, err := utils.FromLine(line, " ")
		if err != nil {
			return err
		}
		d.reports = append(d.reports, rep)
	}
	return nil
}

func (d *Day2) Part1() (int, error) {
	safeNb := 0
	for _, rep := range d.reports {
		if d.verifyWithError(rep) {
			safeNb += 1
		}
	}
	return safeNb, nil
}

func (d *Day2) Part2() (int, error) {
	safeNb := 0
	for _, rep := range d.reports {
		for i := -1; i < len(rep); i += 1 {
			var repDel []int
			if i == -1 {
				repDel = rep
			} else {
				repDel = utils.OmitIndex(rep, i)
			}
			if d.verifyWithError(repDel) {
				safeNb += 1
				break
			}
		}
	}
	return safeNb, nil
}

func (d *Day2) checkDelta(prevDelta, delta int) bool {
	return prevDelta*delta < 0 || utils.Abs(delta) > 3 || utils.Abs(delta) < 1
}

func (d *Day2) verifyWithError(rep []int) bool {
	prevDelta := 0
	safe := true
	for i := 1; i < len(rep); i += 1 {
		delta := rep[i-1] - rep[i]

		// Different sign or difference too important
		if d.checkDelta(prevDelta, delta) {
			safe = false
			break
		}
		prevDelta = delta
	}
	return safe
}
