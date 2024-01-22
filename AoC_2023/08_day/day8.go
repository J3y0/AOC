package main

import (
	"08_day/parsing"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	data, err := os.ReadFile("./data/day8.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while reading file: %v\n", err)
		os.Exit(1)
	}

	instructions, mapJunction := parsing.ParseFile(data)

	// -------- Part 1 --------
	tStart := time.Now()
	part1 := CountSteps(instructions, mapJunction)
	t1 := time.Since(tStart)
	fmt.Printf("Result for part 1: %d\n", part1)
	fmt.Printf("Time for part 1: %s\n\n", t1)

	// -------- Part 2 --------
	tStart = time.Now()
	part2 := CountMultipleCycles(instructions, mapJunction)
	t2 := time.Since(tStart)
	fmt.Printf("Result for part 2: %d\n", part2)
	fmt.Printf("Time for part 2: %s\n", t2)
}

func CountSteps(instructions string, mapJunction map[string]parsing.Junction) int {
	var (
		stepCounter int    = 0
		currentStep string = "AAA"
		length      int    = len(instructions)
	)

loop:
	for currentStep != "ZZZ" {
		switch string(instructions[stepCounter%length]) {
		case "L":
			currentStep = mapJunction[currentStep].Left
		case "R":
			currentStep = mapJunction[currentStep].Right
		default:
			fmt.Printf("[!] Unknown instructions")
			break loop
		}

		stepCounter++
	}

	return stepCounter
}

func TaskFindCycleLength(instructions string, mapJunction map[string]parsing.Junction, start string, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		stepCounter int    = 0
		currentStep string = start
		length      int    = len(instructions)
	)

loop:
	for !strings.HasSuffix(currentStep, "Z") {
		switch string(instructions[stepCounter%length]) {
		case "L":
			currentStep = mapJunction[currentStep].Left
		case "R":
			currentStep = mapJunction[currentStep].Right
		default:
			fmt.Printf("[!] Unknown instructions")
			break loop
		}

		stepCounter++
	}

	resultChan <- stepCounter
}

func CountMultipleCycles(instructions string, mapJunction map[string]parsing.Junction) int {
	var (
		res        int = 1
		resultChan     = make(chan int)
		wg         sync.WaitGroup
	)

	start := parsing.FindStart(mapJunction)
	for _, key := range start {
		wg.Add(1)
		go TaskFindCycleLength(instructions, mapJunction, key, resultChan, &wg)
	}

	/* This is done in another goroutine to avoid deadlock
	With the code below:
	---
	wg.Wait()
	close(resultChan)
	---
	There is a deadlock as main thread is waiting for others to finish but
	other threads wait for the channel to receive their data
	*/
	go func() {
		wg.Wait()         // Wait for all goroutine to finish
		close(resultChan) // Close the channel as no more data can be received
	}()

	// Compute LCM over results
	for result := range resultChan {
		res = (res * result) / GCD(res, result)
	}

	return res
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
