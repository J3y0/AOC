package main

import (
	"06_day/parsing"
	"fmt"
	"os"
	"time"
)

func main() {
	data, err := os.ReadFile("./data/day6.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while reading file: %v\n", err)
		os.Exit(1)
	}

	races, err := parsing.ParsePart1(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while parsing input: %v\n", err)
		os.Exit(1)
	}

	tStart := time.Now()

	part1 := WaysBeatRecord(races)

	t1 := time.Since(tStart)

	fmt.Printf("Result for part1: %d\n", part1)
	fmt.Printf("Time for part1: %s\n\n", t1)

	race, err := parsing.ParsePart2(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while parsing input: %v\n", err)
		os.Exit(1)
	}

	tStart = time.Now()
	part2 := WaysBeatRecord(race)
	t2 := time.Since(tStart)

	fmt.Printf("Result for part2: %d\n", part2)
	fmt.Printf("Time for part2: %s\n", t2)
}

func WaysBeatRecord(races []parsing.TimeDistance) int {
	var res int = 1
	for i := 0; i < len(races); i++ {
		dist := races[i].Distance
		totalWays := 0
		time := races[i].Time

		for k := 1; k < time/2+1; k++ { // Symmetrical pb, only need to search for first half
			// k*(time-k) > dist ? if yes, valid range is: k->(n-k) so there is n-2k+1 ways
			if k*(time-k) > dist {
				totalWays = time - 2*k + 1
				break
			}
		}

		res *= totalWays
	}

	return res
}
