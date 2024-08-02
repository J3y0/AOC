package main

import (
	"09_day/parsing"
	"fmt"
	"math"
	"os"
)

func main() {
	data, err := os.ReadFile("./data/day9.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while reading file: %v", err)
		os.Exit(1)
	}

	linesNumbers, err := parsing.Parse(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while parsing: %v", err)
		os.Exit(1)
	}

	part1, part2 := SumPredicted(linesNumbers)
	fmt.Printf("Result for part 1: %d\n", part1)
	fmt.Printf("Result for part 2: %d\n", part2)
}

func SumPredicted(linesNumbers [][]int) (int, int) {
	var (
		sum1 int
		sum2 int
	)

	for _, lineNbs := range linesNumbers {
		nextPred, prevPred := PredictValue(lineNbs)
		sum1 += nextPred
		sum2 += prevPred
	}

	return sum1, sum2
}

func PredictValue(lineNbs []int) (int, int) {
	deltasMatrix := make([][]int, 0)
	deltasMatrix = append(deltasMatrix, lineNbs)

	// Compute deltas until we have only zeroes
	for !IsZeroArr(deltasMatrix[len(deltasMatrix)-1]) {
		var (
			newDeltasLine  []int
			lastDeltasLine []int = deltasMatrix[len(deltasMatrix)-1]
		)

		for i := 0; i < len(lastDeltasLine)-1; i++ {
			newDeltasLine = append(newDeltasLine, lastDeltasLine[i+1]-lastDeltasLine[i])
		}

		deltasMatrix = append(deltasMatrix, newDeltasLine)
	}

	var (
		nextPredicted int // Sum last values for next prediction
		prevPredicted int // Sum (-1 every two lines)*(first values) for previous prediction
	)

	for i, deltaLine := range deltasMatrix {
		nextPredicted += deltaLine[len(deltaLine)-1]
		prevPredicted += int(math.Pow(-1, float64(i%2))) * deltaLine[0]
	}

	return nextPredicted, prevPredicted
}

func IsZeroArr(arr []int) bool {
	for _, elt := range arr {
		if elt != 0 {
			return false
		}
	}
	return true
}
