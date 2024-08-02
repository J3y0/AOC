package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"slices"
	"strings"
)

type Rocks struct {
	X int
	Y int
}

// Direction
const (
	NORTH = 1
	SOUTH = -1
	WEST  = 1
	EAST  = -1
)

func main() {
	rockMap, rocks, err := ParseRocks("./data/day14.txt")
	if err != nil {
		panic(err)
	}

	// ------ Part1: Roll rocks to North -----
	RollNorthSouth(rockMap, rocks, NORTH)
	part1 := TotalLoad(rocks, len(rockMap))
	fmt.Println("Result for part 1:", part1)

	// ------ Part 2: Cycle then compute load -----
	part2 := Cycle(1000000000, rockMap, rocks)
	fmt.Println("Result for part 2:", part2)
}

func Cycle(N int, rockMap []string, rocks []Rocks) int {
	var (
		state    uint64
		seen     = make(map[uint64]int) // Used as a "cache"
		cycleLen = N
	)
	for i := 0; i < N; i++ {
		// Define properly our state
		h := fnv.New64a()
		h.Write([]byte(fmt.Sprintf("%v", rocks)))
		state = h.Sum64()
		// If state already seen, then we hit a cycle
		if firstSeenOff, ok := seen[state]; ok {
			cycleLen = i - firstSeenOff // Compute cycle length
			break
		} else {
			seen[state] = i // At which index we have seen that state
		}

		RollNorthSouth(rockMap, rocks, NORTH)
		RollWestEast(rockMap, rocks, WEST)
		RollNorthSouth(rockMap, rocks, SOUTH)
		RollWestEast(rockMap, rocks, EAST)
	}

	firstSeenOff := seen[state]
	for i := 0; i < (N-firstSeenOff)%cycleLen; i++ {
		RollNorthSouth(rockMap, rocks, NORTH)
		RollWestEast(rockMap, rocks, WEST)
		RollNorthSouth(rockMap, rocks, SOUTH)
		RollWestEast(rockMap, rocks, EAST)
	}

	return TotalLoad(rocks, len(rockMap))
}

func TotalLoad(rocks []Rocks, nbLines int) int {
	var totalLoad int

	for _, rock := range rocks {
		totalLoad += nbLines - rock.X
	}

	return totalLoad
}

func RollNorthSouth(rockMap []string, rocks []Rocks, direction int) {
	var (
		// The map keeps track of the idx where the rock will be blocked
		lastRockPosPerCol map[int]int
		initValue         int
	)
	if direction == NORTH {
		initValue = 0
		lastRockPosPerCol = InitRockIdxMap(initValue, len(rockMap[0]))
	} else {
		initValue = len(rockMap) - 1
		lastRockPosPerCol = InitRockIdxMap(initValue, len(rockMap[0]))
	}
	// Sort rocks according to X coordinates (so I don't start rolling a rock which is behind a rock)
	slices.SortStableFunc(rocks, func(a, b Rocks) int {
		return direction * (a.X - b.X)
	})

	for i := range rocks {
		rock := &rocks[i]
		colIdx := rock.Y

		for j := rock.X; j >= 0 && j < len(rockMap); j -= direction {
			// Reach map limit or other rock
			if j == lastRockPosPerCol[colIdx] {
				rock.X = j
				lastRockPosPerCol[colIdx] += direction
				break
			}

			// Wall
			if rockMap[j-direction][colIdx] == '#' {
				lastRockPosPerCol[colIdx] = j + direction
				rock.X = j
				break
			}
		}
	}
}

func RollWestEast(rockMap []string, rocks []Rocks, direction int) {
	var (
		// The map keeps track of the idx where the rock will be blocked
		lastRockPosPerLine map[int]int
		initValue          int
	)
	if direction == WEST {
		initValue = 0
		lastRockPosPerLine = InitRockIdxMap(initValue, len(rockMap))
	} else {
		initValue = len(rockMap[0]) - 1
		lastRockPosPerLine = InitRockIdxMap(initValue, len(rockMap))
	}
	// Sort rocks according to X coordinates (so I don't start rolling a rock which is behind a rock)
	slices.SortStableFunc(rocks, func(a, b Rocks) int {
		return direction * (a.Y - b.Y)
	})

	for i := range rocks {
		rock := &rocks[i]
		lineIdx := rock.X

		for j := rock.Y; j >= 0 && j < len(rockMap[0]); j -= direction {
			// Reach map limit or other rock
			if j == lastRockPosPerLine[lineIdx] {
				rock.Y = j
				lastRockPosPerLine[lineIdx] += direction
				break
			}

			// Wall
			if rockMap[lineIdx][j-direction] == '#' {
				lastRockPosPerLine[lineIdx] = j + direction
				rock.Y = j
				break
			}
		}
	}
}

func InitRockIdxMap(initValue int, nbCol int) map[int]int {
	result := make(map[int]int)
	for i := 0; i < nbCol; i++ {
		result[i] = initValue
	}
	return result
}

func ParseRocks(path string) (lines []string, rocks []Rocks, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	i := 0
	s := bufio.NewScanner(file)
	for s.Scan() {
		readLine := strings.TrimSpace(s.Text())
		lines = append(lines, readLine)

		// Add rocks
		for j := range readLine {
			if readLine[j] == 'O' {
				rocks = append(rocks, Rocks{
					X: i,
					Y: j,
				})
			}
		}
		i++
	}

	return
}
