package main

import (
	"fmt"
	"time"
)

func main() {
	workflows, parts, err := ParseInput("./data/day19.txt")
	if err != nil {
		panic(err)
	}

	tStart := time.Now()
	// ----- Part 1 -----
	part1, err := Part1(parts, workflows)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1, time.Since(tStart))

	// ----- Part 2 -----
	partRange := map[rune]Range{
		'x': {Start: 1, End: 4000},
		'm': {Start: 1, End: 4000},
		'a': {Start: 1, End: 4000},
		's': {Start: 1, End: 4000},
	}
	part2 := Part2(partRange, &workflows, "in")
	fmt.Println("Part 2:", part2, time.Since(tStart))
}
