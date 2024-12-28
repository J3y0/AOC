package days

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Day24 struct {
	computed   map[string]int
	operations [][]string
}

func (d *Day24) Part1() (int, error) {
	queue := d.operations
	for len(queue) > 0 {
		op := queue[0]
		queue = queue[1:]

		// If not computed yet, continue
		if _, ok := d.computed[op[0]]; !ok {
			queue = append(queue, op)
			continue
		}
		if _, ok := d.computed[op[2]]; !ok {
			queue = append(queue, op)
			continue
		}

		// Switch operation
		switch op[1] {
		case "OR":
			d.computed[op[3]] = d.computed[op[0]] | d.computed[op[2]]
		case "XOR":
			d.computed[op[3]] = d.computed[op[0]] ^ d.computed[op[2]]
		case "AND":
			d.computed[op[3]] = d.computed[op[0]] & d.computed[op[2]]
		default:
			return 0, fmt.Errorf("Unknown operation\n")
		}
	}

	return d.getWire("z"), nil
}

func (d *Day24) Part2() (int, error) {
	badZ := make([][]string, 0)
	badNZ := make([][]string, 0)
	for _, op := range d.operations {
		if strings.HasPrefix(op[3], "z") && op[3] != "z45" && op[1] != "XOR" {
			badZ = append(badZ, op)
		}

		if !strings.HasPrefix(op[3], "z") && !strings.HasPrefix(op[0], "x") && !strings.HasPrefix(op[0], "y") && op[1] == "XOR" {
			badNZ = append(badNZ, op)
		}
	}

	for _, res := range badNZ {
		zToSwap := d.findNextZ(res[3])
		for _, fz := range badZ {
			if fz[3] == zToSwap {
				// Swap
				resIdx := slices.IndexFunc(d.operations, func(a []string) bool {
					return slices.Equal(a, res)
				})

				fzIdx := slices.IndexFunc(d.operations, func(a []string) bool {
					return slices.Equal(a, fz)
				})

				d.operations[resIdx][3], d.operations[fzIdx][3] = d.operations[fzIdx][3], d.operations[resIdx][3]
			}
		}
	}

	z, err := d.Part1()
	if err != nil {
		return 0, nil
	}
	x, y := d.getWire("x"), d.getWire("y")
	diff := strconv.FormatInt(int64((x+y)^z), 2)
	totZero := strings.Count(diff, "0")
	// Spot last errors in operations (AND and XOR) implying x and y with idx = totZero
	badCarry := make([][]string, 0)
	for _, op := range d.operations {
		if op[0] == fmt.Sprintf("x%02d", totZero) || op[0] == fmt.Sprintf("y%02d", totZero) {
			badCarry = append(badCarry, op)
		}
	}

	var bad []string
	for _, elt := range slices.Concat(badZ, badNZ, badCarry) {
		bad = append(bad, elt[3])
	}
	sort.Strings(bad)
	fmt.Println("Part2: (Not a number)", strings.Join(bad, ","))

	return 0, nil
}

func (d *Day24) findNextZ(res string) string {
	var next []string
	for _, op := range d.operations {
		if op[0] == res || op[2] == res {
			next = op
		}
	}

	if strings.HasPrefix(next[3], "z") {
		idx, _ := strconv.Atoi(next[3][1:])
		return fmt.Sprintf("z%02d", idx-1)
	}

	return d.findNextZ(next[3])
}

func (d *Day24) getWire(name string) int {
	keys := []string{}
	for k := range d.computed {
		if strings.HasPrefix(k, name) {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	sum := 0
	for i, z := range keys {
		sum += d.computed[z] * int(math.Pow(2, float64(i)))
	}

	return sum
}

func (d *Day24) Parse() error {
	content, err := os.ReadFile("./data/24_day.txt")
	if err != nil {
		return err
	}

	splitted := strings.Split(strings.TrimSpace(string(content)), "\n\n")
	d.computed = make(map[string]int)
	inits := strings.Split(splitted[0], "\n")
	for _, init := range inits {
		s := strings.Split(init, ": ")
		d.computed[s[0]], _ = strconv.Atoi(s[1])
	}

	ops := strings.Split(splitted[1], "\n")
	d.operations = make([][]string, len(ops))
	for i, op := range ops {
		d.operations[i] = d.parseOp(op)
	}
	return nil
}

func (d *Day24) parseOp(op string) []string {
	var res []string
	out := strings.Split(op, " -> ")
	res = strings.Split(out[0], " ")
	res = append(res, out[1])
	return res
}
