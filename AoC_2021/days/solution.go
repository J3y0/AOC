package days

import (
	"main/utils"
	"time"
)

type NotInRangeError struct {
	message string
}

func (e *NotInRangeError) Error() string {
	return e.message
}

type Solution interface {
	Part1() (int, error)
	Part2() (int, error)
}

func RunSelectedSolution(day, part int) error {
	var solutionToRun Solution
	switch day {
	case 1:
		solutionToRun = &Day1{}
	case 2:
		solutionToRun = &Day2{}
	case 3:
		solutionToRun = &Day3{}
	case 4:
		solutionToRun = &Day4{}
	case 6:
		solutionToRun = &Day6{}
	case 7:
		solutionToRun = &Day7{}
	default:
		return &NotInRangeError{message: "day should be between 1 and 25"}
	}

	var (
		part1, part2 int
		time1, time2 time.Duration
		err          error
	)
	switch part {
	case 0:
		start := time.Now()
		part1, err = solutionToRun.Part1()
		if err != nil {
			return err
		}
		time1 = time.Since(start)

		start = time.Now()
		part2, err = solutionToRun.Part2()
		if err != nil {
			return err
		}
		time2 = time.Since(start)
	case 1:
		start := time.Now()
		part1, err = solutionToRun.Part1()
		if err != nil {
			return err
		}
		time1 = time.Since(start)
	case 2:
		start := time.Now()
		part2, err = solutionToRun.Part2()
		if err != nil {
			return err
		}
		time2 = time.Since(start)
	default:
		return &NotInRangeError{message: "part should be between 0 and 2"}
	}

	utils.FormatAndPrintResult(day, part1, part2, time1, time2)

	return nil
}
