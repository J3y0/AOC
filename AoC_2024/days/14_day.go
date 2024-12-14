package days

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"aoc/utils"
)

type robot struct {
	x, y, vx, vy int
}

func (r *robot) updatePos(seconds, maxRow, maxCol int) {
	r.x = utils.PosMod((r.x + r.vx*seconds), maxRow)
	r.y = utils.PosMod((r.y + r.vy*seconds), maxCol)
}

type Day14 struct {
	robots         []*robot
	maxRow, maxCol int
}

func (d *Day14) printRobots() {
	for i := range d.maxRow {
		for j := range d.maxCol {
			robot := false
			for _, r := range d.robots {
				if r.x == i && r.y == j {
					fmt.Print("r")
					robot = true
					break
				}
			}
			if !robot {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (d *Day14) Part1() (int, error) {
	seconds := 100
	robots := make([]*robot, 0)
	copy(robots, d.robots)
	for _, r := range robots {
		r.updatePos(seconds, d.maxRow, d.maxCol)
	}
	return d.countQuad(), nil
}

func (d *Day14) Part2() (int, error) {
	mini := math.MaxInt
	iterationMin := 0
	for s := 1; s < 10000; s++ {
		for _, r := range d.robots {
			r.updatePos(1, d.maxRow, d.maxCol)
		}

		safety := d.countQuad()
		if mini > safety {
			mini = safety
			iterationMin = s
		}

	}

	return iterationMin, nil
}

func (d *Day14) Parse() error {
	f, err := os.Open("./data/14_day.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	d.maxRow = 103
	d.maxCol = 101
	d.robots = make([]*robot, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y, vx, vy int
		_, err := fmt.Sscanf(s.Text(), "p=%d,%d v=%d,%d", &y, &x, &vy, &vx)
		if err != nil {
			return err
		}
		d.robots = append(d.robots, &robot{x, y, vx, vy})
	}

	return nil
}

func (d *Day14) countQuad() int {
	mx := d.maxRow / 2
	my := d.maxCol / 2
	totQuad := [4]int{0}

	for _, r := range d.robots {
		if r.x < mx && r.y < my {
			totQuad[0]++
		} else if r.x < mx && r.y > my {
			totQuad[1]++
		} else if r.x > mx && r.y < my {
			totQuad[2]++
		} else if r.x > mx && r.y > my {
			totQuad[3]++
		}
	}

	tot := 1
	for _, elt := range totQuad {
		tot *= elt
	}

	return tot
}
