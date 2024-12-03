package days

import (
	"fmt"
	"time"
)

type Solution interface {
	Part1() (int, error)
	Part2() (int, error)
	Parse() error
}

func SolutionToRun(day, part int) error {
	var toRun Solution
	switch day {
	case 1:
		toRun = &Day1{}
	case 2:
		toRun = &Day2{}
	case 3:
		toRun = &Day3{}
	default:
	}

	err := toRun.Parse()
	if err != nil {
		return err
	}

	var tStart, tEnd time.Time
	switch part {
	case 1:
		tStart = time.Now()
		part1, err := toRun.Part1()
		if err != nil {
			return err
		}
		tEnd = time.Now()

		fmt.Println("Part 1:", part1)
		fmt.Println("Time  :", tEnd.Sub(tStart))
	case 2:
		tStart = time.Now()
		part2, err := toRun.Part2()
		if err != nil {
			return err
		}
		tEnd = time.Now()

		fmt.Println("Part 2:", part2)
		fmt.Println("Time  :", tEnd.Sub(tStart))
	default:
	}

	return nil
}
