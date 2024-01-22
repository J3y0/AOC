package main

import (
	"07_day/parsing"
	"fmt"
	"os"
	"slices"
	"time"
)

func main() {
	data, err := os.ReadFile("./data/day7.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while reading file: %v", err)
	}

	handsPart1, _ := parsing.ParseInput(data, 1)

	tStart := time.Now()
	part1 := totalWinning(handsPart1)
	t1 := time.Since(tStart)

	fmt.Printf("Result for part 1: %d\n", part1)
	fmt.Printf("Time for part 1: %s\n\n", t1)

	handsPart2, _ := parsing.ParseInput(data, 2)

	tStart = time.Now()
	part2 := totalWinning(handsPart2)
	t2 := time.Since(tStart)

	fmt.Printf("Result for part 2: %d\n", part2)
	fmt.Printf("Time for part 2: %s\n", t2)
}

func CmpHand(hand1, hand2 parsing.Hand) int {
	if hand1.Type == hand2.Type {
		for i := range hand1.HandValues {
			if hand1.HandValues[i] != hand2.HandValues[i] {
				return hand1.HandValues[i] - hand2.HandValues[i]
			}
		}
	}

	return hand1.Type - hand2.Type
}

func totalWinning(hands []parsing.Hand) int {
	slices.SortFunc(hands, CmpHand)

	var res int
	for rank, hand := range hands {
		res += (rank + 1) * hand.Bid
	}

	return res
}
