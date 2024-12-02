package days

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"aoc/utils"
)

type Day1 struct {
	row1 []int
	row2 []int
}

func (d *Day1) Part1() (int, error) {
	slices.Sort(d.row1)
	slices.Sort(d.row2)
	sum := 0
	for i := range d.row1 {
		sum += utils.Abs(d.row1[i] - d.row2[i])
	}
	return sum, nil
}

func (d *Day1) Part2() (int, error) {
	similarity := 0
	for _, id := range d.row1 {
		count := 0
		for _, id2 := range d.row2 {
			if id == id2 {
				count += 1
			}
		}
		similarity += count * id
	}
	return similarity, nil
}

func (d *Day1) Parse() error {
	file, err := os.Open("./data/01_day.txt")
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while closing file")
		}
	}(file)

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		line = strings.TrimSpace(line)

		ids := strings.Split(line, "   ")
		id, _ := strconv.Atoi(ids[0])
		d.row1 = append(d.row1, id)
		id, _ = strconv.Atoi(ids[1])
		d.row2 = append(d.row2, id)
	}

	return nil
}
