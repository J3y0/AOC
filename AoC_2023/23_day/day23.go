package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"time"
)

type Pos struct {
	X int
	Y int
}

type HikingMap [][]rune

func ReadMap(path string) (HikingMap, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	hikingMap := make(HikingMap, 0)
	for s.Scan() {
		readLine := strings.TrimSpace(s.Text())
		hikingMap = append(hikingMap, []rune(readLine))
	}

	return hikingMap, nil
}

type Graph map[Pos]map[Pos]int

type QueueItem struct {
	Position Pos
	CurDist  int
}

// BuildGraph use edge contraction to reduce computation time as there are mainly straight paths
func BuildGraph(hikingMap HikingMap, start, end Pos, part int) Graph {
	// Find points that are an intersection, that is with at least 3 neighbors
	nodes := []Pos{start, end}
	for r, row := range hikingMap {
		for c := range row {
			if hikingMap[r][c] == '#' {
				continue
			}

			neighbors := []Pos{
				{X: r + 1, Y: c},
				{X: r, Y: c + 1},
				{X: r - 1, Y: c},
				{X: r, Y: c - 1},
			}
			var valid int
			for _, n := range neighbors {
				if OutsideMap(n, len(hikingMap), len(hikingMap[0])) || hikingMap[n.X][n.Y] == '#' {
					continue
				}
				valid++
			}
			// Intersection if degree of node >= 3
			if valid >= 3 {
				nodes = append(nodes, Pos{X: r, Y: c})
			}
		}
	}

	// Edge contraction
	graph := make(Graph)
	for _, point := range nodes {
		seen := make(map[Pos]bool)
		seen[point] = true
		queue := []QueueItem{{Position: point, CurDist: 0}}
		for len(queue) > 0 {
			// Dequeue
			curPos := queue[0].Position
			curDist := queue[0].CurDist
			queue = queue[1:]

			if curDist != 0 && slices.Contains(nodes, curPos) {
				if graph[point] == nil {
					graph[point] = make(map[Pos]int)
				}
				graph[point][curPos] = curDist
				continue
			}

			// Visit neighbors
			var neighbors []Pos
			if part == 1 {
				switch hikingMap[curPos.X][curPos.Y] {
				case '>':
					neighbors = []Pos{
						{X: curPos.X, Y: curPos.Y + 1},
					}
				case '<':
					neighbors = []Pos{
						{X: curPos.X, Y: curPos.Y - 1},
					}
				case 'v':
					neighbors = []Pos{
						{X: curPos.X + 1, Y: curPos.Y},
					}
				case '^':
					neighbors = []Pos{
						{X: curPos.X - 1, Y: curPos.Y},
					}
				case '.':
					neighbors = []Pos{
						{X: curPos.X + 1, Y: curPos.Y},
						{X: curPos.X, Y: curPos.Y + 1},
						{X: curPos.X - 1, Y: curPos.Y},
						{X: curPos.X, Y: curPos.Y - 1},
					}
				case '#':
					fmt.Println("Error, current pos is a wall")
				}
			} else if part == 2 {
				neighbors = []Pos{
					{X: curPos.X + 1, Y: curPos.Y},
					{X: curPos.X, Y: curPos.Y + 1},
					{X: curPos.X - 1, Y: curPos.Y},
					{X: curPos.X, Y: curPos.Y - 1},
				}
			}
			for _, n := range neighbors {
				visited, exist := seen[n]
				if OutsideMap(n, len(hikingMap), len(hikingMap[0])) || hikingMap[n.X][n.Y] == '#' || (exist && visited) {
					continue
				}
				queue = append(queue, QueueItem{Position: n, CurDist: curDist + 1})
				seen[n] = true
			}
		}
	}

	return graph
}

func DFS(graph *Graph, seen *map[Pos]bool, current, end Pos) int {
	if current == end {
		return 0
	}

	(*seen)[current] = true
	var length = -math.MaxInt
	for n, dist := range (*graph)[current] {
		if already, ok := (*seen)[n]; !ok || (ok && !already) {
			length = max(length, DFS(graph, seen, n, end)+dist)
		}
	}
	(*seen)[current] = false
	return length
}

func OutsideMap(pos Pos, h, w int) bool {
	return pos.X < 0 || pos.X >= h || pos.Y < 0 || pos.Y >= w
}

func main() {
	hikingMap, err := ReadMap("./data/day23.txt")
	if err != nil {
		panic(err)
	}
	start := Pos{
		X: 0,
		Y: 1,
	}
	end := Pos{
		X: len(hikingMap) - 1,
		Y: len(hikingMap[0]) - 2,
	}

	tStart := time.Now()
	// ------ Part 1 -------
	graph := BuildGraph(hikingMap, start, end, 1)
	seen := make(map[Pos]bool)
	part1 := DFS(&graph, &seen, start, end)
	fmt.Println("Part 1:", part1, time.Since(tStart))
	// ------ Part 2 -------
	graph = BuildGraph(hikingMap, start, end, 2)
	seen = make(map[Pos]bool)
	part2 := DFS(&graph, &seen, start, end)
	fmt.Println("Part 2:", part2, time.Since(tStart))
}
