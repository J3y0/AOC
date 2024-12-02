package days

import "fmt"

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
	default:
	}

	err := toRun.Parse()
	if err != nil {
		return err
	}

	switch part {
	case 1:
		part1, err := toRun.Part1()
		if err != nil {
			return err
		}
		fmt.Println("Part 1:", part1)
	case 2:
		part2, err := toRun.Part2()
		if err != nil {
			return err
		}
		fmt.Println("Part 2:", part2)
	default:
	}

	return nil
}
