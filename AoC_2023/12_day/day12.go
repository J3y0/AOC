package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	path := "./data/day12.txt"

	lines, err := ReadLines(path)
	if err != nil {
		panic(err)
	}

	tStart := time.Now()
	part1 := SumArrangements(lines)
	t1 := time.Since(tStart)
	fmt.Println("Result for part 1:", part1)
	fmt.Println("Time for part 1:", t1)

	tStart = time.Now()
	part2 := SumArrangementsConcurrently(lines) // We can also use function for part 1 with cache (2-3 times slower)
	t2 := time.Since(tStart)
	fmt.Println("Result for part 2:", part2)
	fmt.Println("Time for part 2:", t2)
}

func SumArrangements(lines []string) int {
	var sum int
	for _, line := range lines {
		temp := strings.Split(line, " ")
		springs := temp[0]
		lengthContiguous, _ := ParseToNumbers(temp[1])

		sum += FindArrangements(springs, lengthContiguous)
	}

	return sum
}

func FindArrangements(springs string, lengthContiguous []int) int {
	// Base case
	if len(lengthContiguous) == 0 {
		if strings.Contains(springs, "#") {
			return 0 // Invalid
		} else {
			return 1 // Valid as no springs left or only . or ? that can be cast in .
		}
	}
	if springs == "" {
		if len(lengthContiguous) > 0 {
			return 0 // Invalid
		} else {
			return 1
		}

	}

	var arrangements int
	if springs[0] == '.' || (springs[0] == '?') {
		arrangements += FindArrangements(springs[1:], lengthContiguous)
	}

	if springs[0] == '#' || (springs[0] == '?') {
		// Check we have enough springs to respect next length of contiguous broken springs
		// Check the substring of size length is contiguous
		// Check the character after the substring is not another broken spring
		if len(springs) > lengthContiguous[0] && IsContiguous(springs[:lengthContiguous[0]]) && springs[lengthContiguous[0]] != '#' {
			arrangements += FindArrangements(springs[lengthContiguous[0]+1:], lengthContiguous[1:])
		} else if len(springs) == lengthContiguous[0] && IsContiguous(springs[:lengthContiguous[0]]) {
			// If numbers of springs is exactly equal to the next length of contiguous broken springs
			// Then just check if this is contiguous and pass empty string to the next call
			// (no more springs, just trigger base case)
			arrangements += FindArrangements("", lengthContiguous[1:])
		}
	}

	return arrangements
}
