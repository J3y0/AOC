package days

import (
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type equation struct {
	m     [2][2]int
	prize [2]int
}

type Day13 struct {
	equations []equation
}

func (d *Day13) Part1() (int, error) {
	tokens := 0
	for _, eq := range d.equations {
		// [a b] = inv(m) * prize
		a := eq.m[0][0]
		b := eq.m[0][1]
		c := eq.m[1][0]
		d := eq.m[1][1]
		det := float64(a*d - b*c)

		if det == 0 {
			panic("Det is 0")
		}
		sol1 := float64(eq.prize[0]*d+eq.prize[1]*(-b)) / det
		sol2 := float64(eq.prize[0]*(-c)+eq.prize[1]*a) / det

		// sol must be an int
		if math.Floor(sol1) != sol1 || math.Floor(sol2) != sol2 {
			continue
		}

		if sol1 > 100 || sol2 > 100 || sol1 < 0 || sol2 < 0 {
			continue
		}

		tokens += 3*int(sol1) + int(sol2)
	}
	return tokens, nil
}

func (d *Day13) Part2() (int, error) {
	var tokens int64
	for _, eq := range d.equations {
		// [a b] = inv(m) * prize
		a := eq.m[0][0]
		b := eq.m[0][1]
		c := eq.m[1][0]
		d := eq.m[1][1]
		det := float64(a*d - b*c)
		px := 10000000000000 + eq.prize[0]
		py := 10000000000000 + eq.prize[1]

		if det == 0 {
			panic("Det is 0")
		}
		sol1 := float64(px*d+py*(-b)) / det
		sol2 := float64(px*(-c)+py*a) / det

		// sol must be an int
		if float64(int64(sol1)) != sol1 || float64(int64(sol2)) != sol2 {
			continue
		}

		tokens += 3*int64(sol1) + int64(sol2)
	}
	return int(tokens), nil
}

func (d *Day13) Parse() error {
	content, err := os.ReadFile("./data/13_day.txt")
	if err != nil {
		return err
	}

	r := regexp.MustCompile(".*: X(?:\\+|=)([0-9]+), Y(?:\\+|=)([0-9]+)")
	eqs := strings.Split(strings.TrimSpace(string(content)), "\n\n")
	d.equations = make([]equation, len(eqs))
	for i, eq := range eqs {
		mat := [2][2]int{{0}}
		prize := [2]int{0}
		matches := r.FindAllStringSubmatch(eq, -1)
		for j, m := range matches {
			x, _ := strconv.Atoi(m[1])
			y, _ := strconv.Atoi(m[2])

			if j < 2 {
				mat[0][j] = x
				mat[1][j] = y
			} else {
				prize[0] = x
				prize[1] = y
			}
		}

		d.equations[i] = equation{m: mat, prize: prize}
	}

	return nil
}
