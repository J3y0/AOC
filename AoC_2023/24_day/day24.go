package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Coord struct {
	X int
	Y int
	Z int
}

type Hailstone struct {
	Position Coord
	Velocity Coord
}

func ReadHailstones(path string) ([]Hailstone, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
	hailstones := make([]Hailstone, 0)
	for _, line := range lines {
		split := strings.Split(line, " @ ")
		pos := Coord{}
		vel := Coord{}
		_, err = fmt.Sscanf(split[0], "%d, %d, %d", &pos.X, &pos.Y, &pos.Z)
		if err != nil {
			return nil, err
		}
		_, err = fmt.Sscanf(split[1], "%d, %d, %d", &vel.X, &vel.Y, &vel.Z)
		if err != nil {
			return nil, err
		}

		hailstones = append(hailstones, Hailstone{Position: pos, Velocity: vel})
	}

	return hailstones, nil
}

func CountIntersect(hailstones []Hailstone) int {
	length := len(hailstones)
	var total int
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if Intersect(hailstones[i], hailstones[j]) {
				total++
			}
		}
	}

	return total
}

func Intersect(h1, h2 Hailstone) bool {
	/*
		Test area border for both X and Y
			- Example: between 7 and 27
			- Input  : between 200000000000000 and 400000000000000
	*/
	// Discard Z coordinate for part 1 (suppose both lines are coplanar)
	const MINI = 200000000000000
	const MAXI = 400000000000000

	// Line cartesian equation for first hailstone
	// v_ya*x - v_xa*y = (v_ya*x_a - v_xa*y_a) <=> Ax + By = C
	a1 := h1.Velocity.Y
	b1 := -h1.Velocity.X
	c1 := a1*h1.Position.X + b1*h1.Position.Y
	// Same for the second hailstone
	a2 := h2.Velocity.Y
	b2 := -h2.Velocity.X
	c2 := a2*h2.Position.X + b2*h2.Position.Y
	// Compute determinant
	det := a1*b2 - a2*b1
	if det == 0 {
		return false
	}

	// Intersection point (compelled to divide c coef first as they are huge)
	xInter := float64(b2)*float64(c1)/float64(det) - float64(b1)*float64(c2)/float64(det)
	yInter := float64(a1)*float64(c2)/float64(det) - float64(a2)*float64(c1)/float64(det)

	// Check if intersection point is within the area and if hailstones crossed in the future
	return xInter >= MINI && xInter <= MAXI && yInter >= MINI && yInter <= MAXI &&
		SameSign(xInter-float64(h1.Position.X), float64(h1.Velocity.X)) &&
		SameSign(yInter-float64(h1.Position.Y), float64(h1.Velocity.Y)) &&
		SameSign(xInter-float64(h2.Position.X), float64(h2.Velocity.X)) &&
		SameSign(yInter-float64(h2.Position.Y), float64(h2.Velocity.Y))
}

func SameSign(x, y float64) bool {
	return (x < 0 && y < 0) || (x > 0 && y > 0)
}

func main() {
	hailstones, err := ReadHailstones("./data/day24.txt")
	if err != nil {
		panic(err)
	}

	tStart := time.Now()
	// ------ Part 1 ------
	part1 := CountIntersect(hailstones)
	fmt.Println("Part 1:", part1, time.Since(tStart))
	// Part 2 solved in Python file
}
