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

func SolutionToRun(day, part int) (err error) {
	var toRun Solution
	switch day {
	case 1:
		toRun = &Day1{}
	case 2:
		toRun = &Day2{}
	case 3:
		toRun = &Day3{}
	case 4:
		toRun = &Day4{}
	case 5:
		toRun = &Day5{}
	case 6:
		toRun = &Day6{}
	case 7:
		toRun = &Day7{}
	case 8:
		toRun = &Day8{}
	case 9:
		toRun = &Day9{}
	case 10:
		toRun = &Day10{}
	case 11:
		toRun = &Day11{}
	case 12:
		toRun = &Day12{}
	case 13:
		toRun = &Day13{}
	case 14:
		toRun = &Day14{}
	case 15:
		toRun = &Day15{}
	case 16:
		toRun = &Day16{}
	case 17:
		toRun = &Day17{}
	case 18:
		toRun = &Day18{}
	case 19:
		toRun = &Day19{}
	case 20:
		toRun = &Day20{}
	case 21:
		toRun = &Day21{}
	case 22:
		toRun = &Day22{}
	case 23:
		toRun = &Day23{}
	case 24:
		toRun = &Day24{}
	case 25:
		toRun = &Day25{}
	default:
	}

	err = toRun.Parse()
	if err != nil {
		return err
	}

	// Run part1, 2 or both
	switch part {
	case 1:
		err = runPart1(toRun)
		if err != nil {
			return err
		}
	case 2:
		err = runPart2(toRun)
		if err != nil {
			return err
		}
	default:
		// Run both part
		err = runPart1(toRun)
		if err != nil {
			return err
		}

		err = runPart2(toRun)
		if err != nil {
			return err
		}
	}

	return nil
}

func runPart1(toRun Solution) error {
	tStart := time.Now()
	part1, err := toRun.Part1()
	if err != nil {
		return err
	}
	tEnd := time.Now()

	fmt.Println("  # ---- Part 1 ---- #")
	fmt.Println("  | Part 1:", part1)
	fmt.Println("  | Time  :", tEnd.Sub(tStart))
	fmt.Println()
	return nil
}

func runPart2(toRun Solution) error {
	tStart := time.Now()
	part1, err := toRun.Part2()
	if err != nil {
		return err
	}
	tEnd := time.Now()

	fmt.Println("  # ---- Part 2 ---- #")
	fmt.Println("  | Part 2:", part1)
	fmt.Println("  | Time  :", tEnd.Sub(tStart))
	fmt.Println()
	return nil
}
