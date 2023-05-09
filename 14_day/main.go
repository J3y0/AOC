package main

import (
	"fmt"
	"os"

	"day14/cave"
)

func main() {
	file, err := os.Open("./data/day14.txt")
	if err != nil {
		fmt.Printf("Error while opening file\n")
	}

	part1, err := solvePart1(file)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1: %d\n", part1)

	part2, err := solvePart2(file)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Part2: %d\n", part2)
}

func solvePart1(file *os.File) (res int, err error) {
	c, err := cave.ParseInput(file)
	if err != nil {
		return
	}
	c.Pretty()

	var again bool = true
	fmt.Println("[~] The simulation is running ...")
	for again {
		again = cave.StepSimulation(c)
		if again {
			res++
		}
	}
	c.Pretty()
	return
}

func solvePart2(file *os.File) (res int, err error) {
	c, err := cave.ParseInput(file)

	// Add ground
	empty := make([]string, len(c.CaveMap[0]), len(c.CaveMap[0]))
	ground := make([]string, 0, len(c.CaveMap[0]))
	for i := 0; i < cap(ground); i++ {
		ground = append(ground, "#")
	}
	c.CaveMap = append(c.CaveMap, empty)
	c.CaveMap = append(c.CaveMap, ground)
	c.MaxX += 2

	if err != nil {
		return
	}
	c.Pretty()

	var again bool = true
	fmt.Println("[~] The simulation is running ...")
	for again {
		again = cave.StepSimulation2(c)
		res++
	}
	c.Pretty()
	return
}
