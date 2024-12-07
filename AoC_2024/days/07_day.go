package days

import (
	"os"
	"strconv"
	"strings"

	"aoc/utils"
)

type Day7 struct {
	equations []string
}

func (d *Day7) Part1() (int, error) {
	tot := 0
	for _, eq := range d.equations {
		s := strings.Split(eq, ": ")
		res, err := strconv.Atoi(s[0])
		if err != nil {
			return 0, err
		}
		op, err := utils.FromLine(s[1], " ")
		if err != nil {
			return 0, err
		}

		if d.recurse(res, op[0], op[1], op[2:]) {
			tot += res
		}
	}

	return tot, nil
}

func (d *Day7) Part2() (int, error) {
	tot := 0
	for _, eq := range d.equations {
		s := strings.Split(eq, ": ")
		res, err := strconv.Atoi(s[0])
		if err != nil {
			return 0, err
		}
		op, err := utils.FromLine(s[1], " ")
		if err != nil {
			return 0, err
		}

		if d.recursePart2(res, op[0], op[1], op[2:]) {
			tot += res
		}
	}

	return tot, nil
}

func (d *Day7) Parse() error {
	content, err := os.ReadFile("./data/07_day.txt")
	if err != nil {
		return err
	}
	d.equations = strings.Split(strings.TrimSpace(string(content)), "\n")
	return nil
}

// [..., l, r, ops]
func (d *Day7) recurse(res, l, r int, ops []int) bool {
	if len(ops) == 0 {
		if l*r == res {
			return true
		}
		if l+r == res {
			return true
		}

		return false
	}

	if l >= res {
		return false
	}

	return d.recurse(res, l+r, ops[0], ops[1:]) || d.recurse(res, l*r, ops[0], ops[1:])
}

// [..., l, r, ops]
func (d *Day7) recursePart2(res, l, r int, ops []int) bool {
	if len(ops) == 0 {
		if l*r == res {
			return true
		}
		if l+r == res {
			return true
		}

		if d.concatenate(l, r) == res {
			return true
		}

		return false
	}

	if l >= res {
		return false
	}

	return (d.recursePart2(res, l+r, ops[0], ops[1:]) ||
		d.recursePart2(res, l*r, ops[0], ops[1:]) ||
		d.recursePart2(res, d.concatenate(l, r), ops[0], ops[1:]))
}

func (d *Day7) concatenate(a, b int) int {
	c := strconv.Itoa(a) + strconv.Itoa(b)
	out, _ := strconv.Atoi(c)
	return out
}
