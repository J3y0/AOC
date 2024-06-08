package utils

import (
	"fmt"
	"time"
)

func FormatAndPrintResultWithTime(part1, part2 int, time1, time2 time.Time) {
	fmt.Printf("Part1: %d in %s\n", part1, time1)
	fmt.Printf("Part2: %d in %s\n", part2, time2)
}

func FormatAndPrintResultWithoutTime(part1, part2 int) {
	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
