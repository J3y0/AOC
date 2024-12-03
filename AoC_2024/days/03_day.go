package days

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Day3 struct {
	memory string
}

func (d *Day3) Parse() error {
	content, err := os.ReadFile("./data/03_day.txt")
	if err != nil {
		return err
	}

	d.memory = "do()" + strings.ReplaceAll(string(content), "\n", "") + "don't()"
	return nil
}

func (d *Day3) Part1() (int, error) {
	return sumMulFromString(d.memory)
}

func (d *Day3) Part2() (int, error) {
	sum := 0

	rValid, err := regexp.Compile("do\\(\\).*?don't\\(\\)")
	if err != nil {
		return 0, err
	}

	validSegs := rValid.FindAllStringIndex(d.memory, -1)
	for _, valid := range validSegs {
		toAdd, err := sumMulFromString(d.memory[valid[0]:valid[1]])
		if err != nil {
			return 0, err
		}

		sum += toAdd
	}

	return sum, nil
}

func sumMulFromString(s string) (int, error) {
	r, err := regexp.Compile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
	if err != nil {
		return 0, err
	}
	matches := r.FindAllStringSubmatch(s, -1)

	sum := 0
	for _, m := range matches {
		left, _ := strconv.Atoi(m[1])
		right, _ := strconv.Atoi(m[2])

		sum += left * right
	}
	return sum, nil
}
