package main

import (
	"flag"
	"fmt"
	"os"

	"aoc/days"
)

func Usage() {
	fmt.Println("Usage: aoc -day <day-number> -part <part-number>")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func main() {
	var day, part int

	flag.IntVar(&day, "day", 1, "run solution for the day provided (1-25)")
	flag.IntVar(&part, "part", 1, "part of the day provided to run (0-2, 0 run both parts)")
	flag.Usage = Usage

	flag.Parse()

	if day < 1 || day > 25 {
		fmt.Fprintln(os.Stderr, "day should be between 1 and 25")
		os.Exit(1)
	}

	if part < 0 || part > 2 {
		fmt.Fprintln(os.Stderr, "part should be 1, 2 or 0 for both parts")
		os.Exit(1)
	}

	fmt.Printf("# --- Running day %d --- #\n", day)
	if err := days.SolutionToRun(day, part); err != nil {
		fmt.Fprintf(os.Stderr, "error %v\n", err)
	}
}
