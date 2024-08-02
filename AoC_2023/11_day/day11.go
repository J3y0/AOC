package main

import (
	"fmt"
	"time"
)

func main() {
	path := "./data/day11.txt"
	// ------- Part 1 -------
	tStart := time.Now()
	galaxies1, err := ParseFile(path, 1)
	if err != nil {
		panic(err)
	}

	part1 := SumDistances(galaxies1)
	fmt.Println("Result for part 1:", part1)

	t1 := time.Since(tStart)
	fmt.Println("Time for part 1:", t1)

	// ------- Part 2 -------
	tStart = time.Now()
	galaxies2, err := ParseFile(path, 2)
	if err != nil {
		panic(err)
	}

	part2 := SumDistances(galaxies2)
	fmt.Println("Result for part 2:", part2)

	t2 := time.Since(tStart)
	fmt.Println("Time for part 2:", t2)
}

func SumDistances(galaxies []GalaxyPos) int {
	var sum int
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += ComputeDistance(galaxies[i], galaxies[j])
		}
	}

	return sum
}

func ComputeDistance(gal1, gal2 GalaxyPos) int {
	return Abs(gal2.X-gal1.X) + Abs(gal2.Y-gal1.Y)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
