package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Pos struct {
	X int
	Y int
	Z int
}

type Brick struct {
	Mini Pos
	Maxi Pos
}

// Overlap Check collision between 2 bricks on X and Y coordinates
func (b *Brick) Overlap(another *Brick) bool {
	return b.Mini.X <= another.Maxi.X && b.Maxi.X >= another.Mini.X &&
		b.Mini.Y <= another.Maxi.Y && b.Maxi.Y >= another.Mini.Y
}

func NewBrick(start, end string) (*Brick, error) {
	startPos := Pos{}
	endPos := Pos{}
	_, err := fmt.Sscanf(start, "%d,%d,%d", &startPos.X, &startPos.Y, &startPos.Z)
	if err != nil {
		return nil, err
	}
	_, err = fmt.Sscanf(end, "%d,%d,%d", &endPos.X, &endPos.Y, &endPos.Z)
	if err != nil {
		return nil, err
	}

	return &Brick{
		Mini: Pos{
			X: min(startPos.X, endPos.X),
			Y: min(startPos.Y, endPos.Y),
			Z: min(startPos.Z, endPos.Z),
		},
		Maxi: Pos{
			X: max(startPos.X, endPos.X),
			Y: max(startPos.Y, endPos.Y),
			Z: max(startPos.Z, endPos.Z),
		},
	}, nil
}

func ReadBricks(path string) (bricks []*Brick, err error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return
	}

	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
	var brick *Brick
	for _, line := range lines {
		split := strings.Split(line, "~")
		brick, err = NewBrick(split[0], split[1])
		if err != nil {
			return
		}
		bricks = append(bricks, brick)
	}
	return
}

func SettleSnapshot(bricks []*Brick) {
	// Sort by Z axis (first element is the closest to the ground)
	slices.SortStableFunc(bricks, func(a, b *Brick) int {
		return a.Mini.Z - b.Mini.Z
	})

	for i, brick := range bricks {
		var maxZ = 1
		for _, below := range bricks[:i] {
			if brick.Overlap(below) {
				maxZ = max(maxZ, below.Maxi.Z+1)
			}
		}
		brick.Maxi.Z = brick.Maxi.Z - brick.Mini.Z + maxZ
		brick.Mini.Z = maxZ
	}

	// Sort again to be sure nothing changed
	slices.SortStableFunc(bricks, func(a, b *Brick) int {
		return a.Mini.Z - b.Mini.Z
	})
}

func GenerateSupportMaps(bricks []*Brick) (map[int][]int, map[int][]int) {
	isSupportedMap := make(map[int][]int)
	supportMap := make(map[int][]int)
	for i, brick := range bricks {
		for j, below := range bricks[:i] {
			if brick.Overlap(below) && brick.Mini.Z == below.Maxi.Z+1 {
				if entry, ok := isSupportedMap[i]; ok {
					entry = append(entry, j)
					isSupportedMap[i] = entry
				} else {
					isSupportedMap[i] = make([]int, 0)
					isSupportedMap[i] = append(isSupportedMap[i], j)
				}

				if entry, ok := supportMap[j]; ok {
					entry = append(entry, i)
					supportMap[j] = entry
				} else {
					supportMap[j] = make([]int, 0)
					supportMap[j] = append(supportMap[j], i)
				}
			}
		}
	}

	return supportMap, isSupportedMap
}

func CountDesintegrated(bricks []*Brick, supportMap, isSupportedMap map[int][]int) int {
	var total int
	for i := range bricks {
		var safe = true
		for _, idx := range supportMap[i] {
			if len(isSupportedMap[idx]) < 2 {
				safe = false
				break
			}
		}
		if safe {
			total++
		}
	}
	return total
}

func ChainReaction(bricks []*Brick, supportMap, isSupportedMap map[int][]int) int {
	var sum int
	for i := range bricks {
		queue := make([]int, 0)
		alreadyFell := make([]int, 0)
		// Add bricks that fall to the queue
		for _, idx := range supportMap[i] {
			if len(isSupportedMap[idx]) == 1 {
				queue = append(queue, idx)
			}
		}

		alreadyFell = append(alreadyFell, i)
		for len(queue) > 0 {
			// Dequeue
			j := queue[0]
			queue = queue[1:]

			for _, idx := range supportMap[j] {
				// If only one support in isSupported - alreadyFell
				inter := Intersect(isSupportedMap[idx], alreadyFell)
				if len(isSupportedMap[idx])-len(inter) == 1 {
					queue = append(queue, idx)
				}
			}

			alreadyFell = append(alreadyFell, j)
		}
		sum += len(alreadyFell) - 1
	}
	return sum
}

func Intersect(a, b []int) []int {
	aCpy := slices.Clone(a)
	bCpy := slices.Clone(b)
	result := make([]int, 0)

	slices.Sort(aCpy)
	slices.Sort(bCpy)
	var i, j int
	for i < len(aCpy) && j < len(bCpy) {
		if aCpy[i] == bCpy[j] {
			result = append(result, aCpy[i])
			i++
			j++
		} else if aCpy[i] > bCpy[j] {
			j++
		} else {
			i++
		}
	}

	return result
}

func main() {
	bricks, err := ReadBricks("./data/day22.txt")
	if err != nil {
		panic(err)
	}

	tStart := time.Now()
	SettleSnapshot(bricks)
	supportMap, isSupportedMap := GenerateSupportMaps(bricks)
	// ------ Part 1 ------
	part1 := CountDesintegrated(bricks, supportMap, isSupportedMap)
	fmt.Println("Part 1:", part1, time.Since(tStart))

	// ------ Part 2 ------
	part2 := ChainReaction(bricks, supportMap, isSupportedMap)
	fmt.Println("Part 2:", part2, time.Since(tStart))
}
