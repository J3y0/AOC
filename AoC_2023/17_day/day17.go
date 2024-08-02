package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type IslandMap [][]uint8

func ReadMap(path string) (IslandMap, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	islandMap := make(IslandMap, 0)
	for s.Scan() {
		readLine := strings.TrimSpace(s.Text())
		row := make([]uint8, len(readLine))
		var nb int
		for i, char := range readLine {
			nb, err = strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			row[i] = uint8(nb)
		}

		islandMap = append(islandMap, row)
	}

	return islandMap, nil
}

type Direction [2]int

type StateTile struct {
	Pos         [2]int
	Direction   [2]int
	Consecutive int
}

func CountHeatLoss(islandMap IslandMap, start, end [2]int, minConsecutive, maxConsecutive int) int {
	// Queue
	pointsToExplore := []StateTile{
		{Pos: start, Direction: Direction{1, 0}, Consecutive: 1},
		{Pos: start, Direction: Direction{0, 1}, Consecutive: 1},
	}
	// Keep track of the shortest path for each vertex
	heatLossRecord := map[StateTile]int{{Pos: start, Direction: Direction{0, 0}, Consecutive: 0}: 0}

	var minHeatLoss = math.MaxInt
	for len(pointsToExplore) > 0 {
		// Pop first point in the queue
		curPoint := pointsToExplore[0]
		pointsToExplore = pointsToExplore[1:]

		// End if current is the end and the last segment has at least minConsecutive moves in same direction
		if curPoint.Pos == end && curPoint.Consecutive >= minConsecutive {
			minHeatLoss = min(minHeatLoss, heatLossRecord[curPoint])
		}

		// Add different neighbors to pointsToExplore (if possible)
		// First,  same direction if possible
		sameDirState := StateTile{
			Pos:         [2]int{curPoint.Pos[0] + curPoint.Direction[0], curPoint.Pos[1] + curPoint.Direction[1]},
			Direction:   curPoint.Direction,
			Consecutive: curPoint.Consecutive + 1,
		}
		if InsideMap(sameDirState.Pos, len(islandMap), len(islandMap[0])) && sameDirState.Consecutive <= maxConsecutive {
			alternativeHeatLoss := heatLossRecord[curPoint] + int(islandMap[sameDirState.Pos[0]][sameDirState.Pos[1]])
			curLoss, ok := heatLossRecord[sameDirState]
			if !ok || alternativeHeatLoss < curLoss {
				heatLossRecord[sameDirState] = alternativeHeatLoss
				pointsToExplore = append(pointsToExplore, sameDirState)
			}
		}

		// Directions on the side
		// Left direction from current direction's perspective
		onLeftOfCur := Direction{-curPoint.Direction[1], curPoint.Direction[0]}
		leftDirState := StateTile{
			Pos:         [2]int{curPoint.Pos[0] + onLeftOfCur[0], curPoint.Pos[1] + onLeftOfCur[1]},
			Direction:   onLeftOfCur,
			Consecutive: 1,
		}
		// For part2, we cannot consider those neighbors if our number of consecutive moves is not sufficient
		if InsideMap(leftDirState.Pos, len(islandMap), len(islandMap[0])) && curPoint.Consecutive >= minConsecutive {
			alternativeHeatLoss := heatLossRecord[curPoint] + int(islandMap[leftDirState.Pos[0]][leftDirState.Pos[1]])
			curLoss, ok := heatLossRecord[leftDirState]
			// Add if it does not exist or if we found a shorter path
			if !ok || alternativeHeatLoss < curLoss {
				heatLossRecord[leftDirState] = alternativeHeatLoss
				pointsToExplore = append(pointsToExplore, leftDirState)
			}
		}

		// Right direction from current direction's perspective
		onRightOfCur := Direction{curPoint.Direction[1], -curPoint.Direction[0]}
		rightDirState := StateTile{
			Pos:         [2]int{curPoint.Pos[0] + onRightOfCur[0], curPoint.Pos[1] + onRightOfCur[1]},
			Direction:   onRightOfCur,
			Consecutive: 1,
		}
		if InsideMap(rightDirState.Pos, len(islandMap), len(islandMap[0])) && curPoint.Consecutive >= minConsecutive {
			alternativeHeatLoss := heatLossRecord[curPoint] + int(islandMap[rightDirState.Pos[0]][rightDirState.Pos[1]])
			curLoss, ok := heatLossRecord[rightDirState]
			if !ok || alternativeHeatLoss < curLoss {
				heatLossRecord[rightDirState] = alternativeHeatLoss
				pointsToExplore = append(pointsToExplore, rightDirState)
			}
		}
	}

	return minHeatLoss
}

func InsideMap(pos [2]int, h, w int) bool {
	return pos[0] >= 0 && pos[0] < h && pos[1] >= 0 && pos[1] < w
}

func main() {
	islandMap, err := ReadMap("./data/day17.txt")
	if err != nil {
		panic(err)
	}

	start := [2]int{0, 0}
	end := [2]int{len(islandMap) - 1, len(islandMap[0]) - 1}

	tStart := time.Now()
	// ----- Part 1 -----
	part1 := CountHeatLoss(islandMap, start, end, 1, 3)
	fmt.Println("Part 1:", part1, time.Since(tStart))

	// ----- Part 2 -----
	part2 := CountHeatLoss(islandMap, start, end, 4, 10)
	fmt.Println("Part 2:", part2, time.Since(tStart))
}
