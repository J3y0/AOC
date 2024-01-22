package main

import (
	"10_day/parsing"
	"fmt"
	"os"
	"slices"
	"time"
)

func main() {
	data, err := os.ReadFile("./data/day10.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while reading file: %v\n", err)
		os.Exit(1)
	}

	mapIsland, startPos := parsing.ParseMap(data)

	loop := FindLoop(mapIsland, startPos)

	part1 := len(loop) / 2
	fmt.Println("Result for part 1:", part1)

	tStart := time.Now()
	part2 := ComputeInside(mapIsland, loop)
	t2 := time.Since(tStart)
	fmt.Println("Result for part 2:", part2)
	fmt.Println("Time for part 2:", t2)
}

func FindLoop(mapIsland [][]string, startPos parsing.Position) []parsing.Position {
	var (
		loopPositions []parsing.Position
		prevPos       parsing.Position = startPos
		curPos        parsing.Position = FindStart(mapIsland, startPos)
	)

	loopPositions = append(loopPositions, startPos)

	for curPos.Symbol != "S" {
		loopPositions = append(loopPositions, curPos)
		nextPos := FindNextPos(prevPos, curPos)
		nextPos.Symbol = mapIsland[nextPos.X][nextPos.Y]

		prevPos = curPos
		curPos = nextPos
	}

	return loopPositions
}

// Find next position based on previous and current postion
func FindNextPos(prevPos parsing.Position, curPos parsing.Position) parsing.Position {
	var nextPos parsing.Position

	switch curPos.Symbol {
	case "|":
		// Column doesn't change
		nextPos.Y = curPos.Y
		if curPos.X+1 == prevPos.X {
			// Comes from bottom
			nextPos.X = curPos.X - 1
		} else {
			// Comes from top
			nextPos.X = curPos.X + 1
		}
	case "-":
		// Line doesn't change
		nextPos.X = curPos.X
		if curPos.Y+1 == prevPos.Y {
			// Comes from right
			nextPos.Y = curPos.Y - 1
		} else {
			// Comes from left
			nextPos.Y = curPos.Y + 1
		}
	case "L":
		if curPos.X-1 == prevPos.X {
			// Comes from top
			nextPos.X = curPos.X
			nextPos.Y = curPos.Y + 1 // Go to the right
		} else {
			// Comes from right
			nextPos.X = curPos.X - 1 // Go to the top
			nextPos.Y = curPos.Y
		}
	case "J":
		if curPos.X-1 == prevPos.X {
			// Comes from top
			nextPos.X = curPos.X
			nextPos.Y = curPos.Y - 1 // Go to the left
		} else {
			// Comes from left
			nextPos.X = curPos.X - 1 // Go to the top
			nextPos.Y = curPos.Y
		}
	case "7":
		if curPos.X+1 == prevPos.X {
			// Comes from bot
			nextPos.X = curPos.X
			nextPos.Y = curPos.Y - 1 // Go to the left
		} else {
			// Comes from left
			nextPos.X = curPos.X + 1 // Go to the bottom
			nextPos.Y = curPos.Y
		}
	case "F":
		if curPos.X+1 == prevPos.X {
			// Comes from bottom
			nextPos.X = curPos.X
			nextPos.Y = curPos.Y + 1 // Go to the right
		} else {
			// Comes from right
			nextPos.X = curPos.X + 1 // Go to the botton
			nextPos.Y = curPos.Y
		}
	default:
		fmt.Println("[!] Not a valid symbol")
	}

	return nextPos
}

// Find direction right after the Start as symbol S doesn't give information on where to go
func FindStart(mapIsland [][]string, startPos parsing.Position) parsing.Position {
	if startPos.Y-1 > 0 && (mapIsland[startPos.X][startPos.Y-1] == "-" ||
		mapIsland[startPos.X][startPos.Y-1] == "F" ||
		mapIsland[startPos.X][startPos.Y-1] == "L") { // Test going left
		return parsing.Position{
			X:      startPos.X,
			Y:      startPos.Y - 1,
			Symbol: mapIsland[startPos.X][startPos.Y-1],
		}
	} else if startPos.Y+1 < len(mapIsland[0]) && (mapIsland[startPos.X][startPos.Y+1] == "-" ||
		mapIsland[startPos.X][startPos.Y+1] == "7" ||
		mapIsland[startPos.X][startPos.Y+1] == "J") { // Test going right
		return parsing.Position{
			X:      startPos.X,
			Y:      startPos.Y + 1,
			Symbol: mapIsland[startPos.X][startPos.Y+1],
		}
	} else if startPos.X-1 > 0 && (mapIsland[startPos.X-1][startPos.Y] == "|" ||
		mapIsland[startPos.X-1][startPos.Y] == "7" ||
		mapIsland[startPos.X-1][startPos.Y] == "F") { // Test going top
		return parsing.Position{
			X:      startPos.X - 1,
			Y:      startPos.Y,
			Symbol: mapIsland[startPos.X-1][startPos.Y],
		}
	}
	// Not found then go bot
	return parsing.Position{
		X:      startPos.X + 1,
		Y:      startPos.Y,
		Symbol: mapIsland[startPos.X+1][startPos.Y],
	}
}

func ComputeInside(mapIsland [][]string, loop []parsing.Position) int {
	/*
		Go through entire map from left to right
		If one point is inside, you will encounter |,J or L character
		an odd number of times (Important note, we count | char
		belonging to the loop)
	*/
	var counterInside int
	for i, line := range mapIsland {
		var (
			allColumnIdx    []int = ColumnInLoop(loop, i)
			curCol          int
			verticalCounter int = 0
		)

		for j := range line {
			if len(allColumnIdx) > 0 && curCol < len(allColumnIdx) && j > allColumnIdx[curCol] {
				verticalCounter++
				curCol++
			}

			// If pos (i, j) belongs to loop continue
			if InLoop(loop, i, j) {
				continue
			}

			// Point is inside
			if verticalCounter%2 == 1 {
				counterInside++
			}
		}
	}

	return counterInside
}

// Return column indexes of every |, J or L belonging to the loop for line i.
// Returned column indexes are sorted in ascending order
func ColumnInLoop(loop []parsing.Position, index int) []int {
	var columnIdx []int

	for _, pos := range loop {
		// Count | (obvious) and L and J because if I go down with L, and then a 7, the point is inside
		// I need to count only one time a 'vertical' bar when both symbols go to the bottom or the top (same for J
		// I don't count its 'pairmate' F that goes in the same direction as it)
		// Don't forget S which should count if it replaces a vertical part
		if pos.X == index && (pos.Symbol == "|" || pos.Symbol == "L" || pos.Symbol == "J" || pos.Symbol == "S") {
			columnIdx = append(columnIdx, pos.Y)
		}
	}

	slices.Sort(columnIdx)
	return columnIdx
}

func InLoop(loop []parsing.Position, x int, y int) bool {
	for _, pos := range loop {
		if pos.X == x && pos.Y == y {
			return true
		}
	}
	return false
}
