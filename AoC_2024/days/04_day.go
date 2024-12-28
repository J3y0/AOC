package days

import (
	"os"
	"strings"

	"aoc/utils"
)

type Day4 struct {
	tab [][]rune
}

func (d *Day4) Part1() (int, error) {
	count := 0
	want := "XMAS"
	for i := range d.tab {
		for j := range d.tab[i] {
			if d.tab[i][j] == 'X' {
				// Check lines
				count += d.checkLines(i, j, want)
				// Check columns
				count += d.checkCols(i, j, want)
				// Check diag
				count += d.checkDiag(i, j, want)
				count += d.checkOtherDiag(i, j, want)
			}
		}
	}
	return count, nil
}

func (d *Day4) Part2() (int, error) {
	count := 0
	for i := range d.tab {
		for j := range d.tab[i] {
			if d.tab[i][j] == 'A' {
				masCol := d.checkDiag(i, j, "AM") > 0 && d.checkDiag(i, j, "AS") > 0
				masOtherCol := d.checkOtherDiag(i, j, "AM") > 0 && d.checkOtherDiag(i, j, "AS") > 0
				if masCol && masOtherCol {
					count += 1
				}
			}
		}
	}
	return count, nil
}

func (d *Day4) Parse() error {
	content, err := os.ReadFile("./data/04_day.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	// Remove empty line
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		d.tab = append(d.tab, []rune(line))
	}

	return nil
}

func (d *Day4) checkLines(i, j int, want string) int {
	count := 0
	size := len(want)
	if j <= len(d.tab[i])-size && string(d.tab[i][j:j+size]) == want {
		count += 1
	}
	// Check reverse
	if j >= size-1 && string(d.tab[i][j-size+1:j+1]) == utils.Reverse(want) {
		count += 1
	}
	return count
}

func (d *Day4) checkCols(i, j int, want string) int {
	count := 0
	size := len(want)
	if i <= len(d.tab)-size {
		match := true
		for off := range want {
			if d.tab[i+off][j] != rune(want[off]) {
				match = false
				break
			}
		}

		if match {
			count += 1
		}
	}

	// Check reverse
	if i >= size-1 {
		match := true
		for off := size - 1; off >= 0; off -= 1 {
			if d.tab[i-off][j] != rune(want[off]) {
				match = false
				break
			}
		}

		if match {
			count += 1
		}
	}
	return count
}

func (d *Day4) checkDiag(i, j int, want string) int {
	count := 0
	size := len(want)
	// \
	//  \ diag
	if i <= len(d.tab)-size && j <= len(d.tab)-size {
		match := true
		for off := range want {
			if d.tab[i+off][j+off] != rune(want[off]) {
				match = false
				break
			}
		}

		if match {
			count += 1
		}
	}

	// Check reverse
	if i >= size-1 && j >= size-1 {
		match := true
		for off := size - 1; off >= 0; off -= 1 {
			if d.tab[i-off][j-off] != rune(want[off]) {
				match = false
				break
			}
		}

		if match {
			count += 1
		}
	}
	return count
}

func (d *Day4) checkOtherDiag(i, j int, want string) int {
	count := 0
	size := len(want)
	//  /
	// /  diag
	if i >= size-1 && j <= len(d.tab)-size {
		match := true
		for off := range want {
			if d.tab[i-off][j+off] != rune(want[off]) {
				match = false
				break
			}
		}

		if match {
			count += 1
		}
	}

	// Check reverse
	if i <= len(d.tab)-size && j >= size-1 {
		match := true
		for off := size - 1; off >= 0; off -= 1 {
			if d.tab[i+off][j-off] != rune(want[off]) {
				match = false
				break
			}
		}

		if match {
			count += 1
		}
	}
	return count
}
