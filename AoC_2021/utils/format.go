package utils

import (
	"fmt"
	"time"
)

func FormatAndPrintResult(day, part1, part2 int, time1, time2 time.Duration) {
	fmt.Printf("|------------ Day %d ------------|\n", day)
	if part1 != 0 {
		fmt.Printf("| Part 1: %d in %s\n", part1, time1)
	}

	if part2 != 0 {
		fmt.Printf("| Part 2: %d in %s\n", part2, time2)
	}
	fmt.Printf("|-------------------------------|\n")
}
