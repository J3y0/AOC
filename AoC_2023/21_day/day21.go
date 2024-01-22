package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Garden [][]rune

type Position [2]int

func ReadGarden(path string) (garden Garden, start Position, err error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return
	}

	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
	garden = make(Garden, len(lines))
	for i, line := range lines {
		//garden = append(garden, []rune(line))
		gardenLine := make([]rune, len(line))
		for j, c := range line {
			if c == 'S' {
				start = [2]int{i, j}
			}
			gardenLine[j] = c
		}
		garden[i] = gardenLine
	}

	return
}

type QueueItem struct {
	Pos            Position
	RemainingSteps int
}

func CountMap(start Position, garden Garden, steps int) int {
	var queue = make([]QueueItem, 0)
	var reachable = make([]Position, 0)
	alreadySeen := make(map[Position]bool)

	alreadySeen[start] = true
	queue = append(queue, QueueItem{Pos: start, RemainingSteps: steps})

	for len(queue) > 0 {
		// Dequeue
		item := queue[0]
		queue = queue[1:]
		remaining := item.RemainingSteps
		pos := item.Pos

		if remaining%2 == 0 {
			reachable = append(reachable, pos)
		}

		if remaining == 0 {
			continue
		}

		neighbors := Neighbors(pos, garden)
		for _, neighbor := range neighbors {
			if _, ok := alreadySeen[neighbor]; !ok {
				alreadySeen[neighbor] = true
				queue = append(queue, QueueItem{Pos: neighbor, RemainingSteps: remaining - 1})
			}
		}
	}

	return LenWithoutDuplicates(reachable)
}

func LenWithoutDuplicates(arr []Position) int {
	var length int
	alreadySeen := make(map[Position]bool)
	for _, pos := range arr {
		if _, ok := alreadySeen[pos]; !ok {
			alreadySeen[pos] = true
			length++
		}
	}
	return length
}

func Neighbors(pos Position, garden Garden) []Position {
	neighbors := make([]Position, 0)
	var (
		nextPos Position
		w       = len(garden[0])
		h       = len(garden)
	)

	// Left
	nextPos = [2]int{pos[0], pos[1] - 1}
	if !IsOutsideGarden(nextPos, w, h) && garden[nextPos[0]][nextPos[1]] != '#' {
		neighbors = append(neighbors, nextPos)
	}

	// Right
	nextPos = [2]int{pos[0], pos[1] + 1}
	if !IsOutsideGarden(nextPos, w, h) && garden[nextPos[0]][nextPos[1]] != '#' {
		neighbors = append(neighbors, nextPos)
	}

	// Up
	nextPos = [2]int{pos[0] - 1, pos[1]}
	if !IsOutsideGarden(nextPos, w, h) && garden[nextPos[0]][nextPos[1]] != '#' {
		neighbors = append(neighbors, nextPos)
	}

	// Down
	nextPos = [2]int{pos[0] + 1, pos[1]}
	if !IsOutsideGarden(nextPos, w, h) && garden[nextPos[0]][nextPos[1]] != '#' {
		neighbors = append(neighbors, nextPos)
	}

	return neighbors
}

func IsOutsideGarden(pos Position, w, h int) bool {
	return pos[0] < 0 || pos[0] >= h || pos[1] < 0 || pos[1] >= w
}

func Part2(start Position, garden Garden, steps int) int {
	size := len(garden)
	// The start position is at the center of the garden: size//2 x size//2

	// Also there is no rocks on the start horizontal and vertical line. It means the furthest point
	// that can be reached is following the horizontal or vertical line (creating a diamond shape)

	// Then, we can compute the number of map in each direction
	nbGardens := steps / size
	// steps = size*nbGardens + size//2. As we start at the center, it means we reach the side of the latest garden
	var nbOddGardens, nbEvenGardens int
	if nbGardens%2 == 0 {
		nbEvenGardens = nbGardens * nbGardens
		nbOddGardens = (nbGardens + 1) * (nbGardens + 1)
	} else {
		nbOddGardens = nbGardens * nbGardens
		nbEvenGardens = (nbGardens + 1) * (nbGardens + 1)
	}
	plotsOddGarden := CountMap(start, garden, 2*size+1) // I just need to encapsulate the odd garden: nb of steps should be large enough and odd
	plotsEvenGarden := CountMap(start, garden, 2*size)  // Same but with an even number of steps (will take the complementary of the result above)

	// We need to add the corners now (the diamond borders go through gardens, leaving corners that need to be considered too
	cornersOdd := plotsOddGarden - CountMap(start, garden, 65)   // Set of 4 corners
	cornersEven := plotsEvenGarden - CountMap(start, garden, 64) // Set of 4 corners
	return plotsEvenGarden*nbEvenGardens + plotsOddGarden*nbOddGardens - (nbGardens+1)*cornersOdd + nbGardens*cornersEven
}

func main() {
	garden, start, err := ReadGarden("./data/day21.txt")
	if err != nil {
		panic(err)
	}

	// ------ Part 1 ------
	tStart := time.Now()
	part1 := CountMap(start, garden, 64)
	fmt.Println("Part 1:", part1, time.Since(tStart))

	// ------ Part 2 ------
	part2 := Part2(start, garden, 26501365)
	fmt.Println("Part 2:", part2, time.Since(tStart))
}
