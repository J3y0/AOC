package days

import (
	"os"
	"strconv"
	"strings"
)

type Day7 struct {
	equations []string
}

func (d *Day7) Part1() (int, error) {
	var tot int64
	for _, eq := range d.equations {
		s := strings.Split(eq, ": ")
		res, err := strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			return 0, err
		}

		nbs := strings.Split(s[1], " ")
		ops := make([]int64, len(nbs))
		for i, elt := range nbs {
			ops[i], _ = strconv.ParseInt(elt, 10, 64)
		}

		if d.recurse(res, ops) {
			tot += res
		}
	}

	return int(tot), nil
}

func (d *Day7) Part2() (int, error) {
	var tot int64
	for _, eq := range d.equations {
		s := strings.Split(eq, ": ")
		res, err := strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			return 0, err
		}

		nbs := strings.Split(s[1], " ")
		ops := make([]int64, len(nbs))
		for i, elt := range nbs {
			ops[i], _ = strconv.ParseInt(elt, 10, 64)
		}

		if d.recursePart2(res, ops) {
			tot += res
		}
	}

	return int(tot), nil
}

func (d *Day7) Parse() error {
	content, err := os.ReadFile("./data/07_day.txt")
	if err != nil {
		return err
	}
	d.equations = strings.Split(strings.TrimSpace(string(content)), "\n")
	return nil
}

func (d *Day7) recurse(res int64, ops []int64) bool {
	n := len(ops)
	if n == 1 {
		return res == ops[0]
	}
	if res%ops[n-1] == 0 && d.recurse(res/ops[n-1], ops[:n-1]) {
		return true
	}

	if res > ops[n-1] && d.recurse(res-ops[n-1], ops[:n-1]) {
		return true
	}
	return false
}

func (d *Day7) recursePart2(res int64, ops []int64) bool {
	n := len(ops)
	if n == 1 {
		return res == ops[0]
	}
	if res%ops[n-1] == 0 && d.recursePart2(res/ops[n-1], ops[:n-1]) {
		return true
	}

	if res > ops[n-1] && d.recursePart2(res-ops[n-1], ops[:n-1]) {
		return true
	}

	strRes := strconv.FormatInt(res, 10)
	strLast := strconv.FormatInt(ops[n-1], 10)
	if len(strRes) > len(strLast) && strings.HasSuffix(strRes, strLast) {
		short, _ := strconv.ParseInt(strRes[:len(strRes)-len(strLast)], 10, 64)
		if d.recursePart2(short, ops[:n-1]) {
			return true
		}
	}
	return false
}
