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
