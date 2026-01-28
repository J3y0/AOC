package days

import (
	"fmt"
	"os"
	"time"
)

type NotImplemented struct {
	message string
}

func (e *NotImplemented) Error() string {
	return e.message
}

type Solution interface {
	Parse(string) error
	Part1() (int, error)
	Part2() (int, error)
}

func RunSolution(day int) error {
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
	case 5:
		solutionToRun = &Day5{}
	case 6:
		solutionToRun = &Day6{}
	case 7:
		solutionToRun = &Day7{}
	case 8:
		solutionToRun = &Day8{}
	case 9:
		solutionToRun = &Day9{}
	case 10:
		solutionToRun = &Day10{}
	case 11:
		solutionToRun = &Day11{}
	case 12:
		solutionToRun = &Day12{}
	case 13:
		solutionToRun = &Day13{}
	default:
		return &NotImplemented{message: "day not implemented yet\n"}
	}

	filename := fmt.Sprintf("./input/%02d_day.txt", day)
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if err = solutionToRun.Parse(string(content)); err != nil {
		return err
	}

	var (
		part1, part2 int
		time1, time2 time.Duration
	)
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

	fmt.Printf("|------------ Day %02d ------------|\n", day)
	fmt.Printf("| Part 1: %d in %d ns\n", part1, time1.Nanoseconds())
	fmt.Printf("| Part 2: %d in %d ns\n", part2, time2.Nanoseconds())
	fmt.Println("|--------------------------------|")

	return nil
}
