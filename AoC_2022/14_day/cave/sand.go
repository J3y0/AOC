package cave

import (
	"fmt"

	"day14/utils"
)

type Direction int

const (
	left Direction = iota
	down
	right
	end
)

func StepSimulation(c *Cave) bool {
	currentSandPos := c.DropSand

	direction, can := canMove(currentSandPos, c)
	for can {
		switch direction {
		// Update the position
		case left:
			currentSandPos.X = currentSandPos.X + 1
			currentSandPos.Y = currentSandPos.Y - 1
		case right:
			currentSandPos.X = currentSandPos.X + 1
			currentSandPos.Y = currentSandPos.Y + 1
		case down:
			currentSandPos.X = currentSandPos.X + 1
		case end:
			// For part 1
			fmt.Println("[!] The simulation has stopped ! Sand now flow into abyssal void")
			return false
		default:
			// Do nothing
		}
		direction, can = canMove(currentSandPos, c)
	}
	c.CaveMap[currentSandPos.X][currentSandPos.Y-c.MinY] = "o"
	return true
}

func canMove(currentSandPos utils.Coords, c *Cave) (Direction, bool) {
	if currentSandPos.X >= len(c.CaveMap)-1 {
		return end, true
	}

	if c.CaveMap[currentSandPos.X+1][currentSandPos.Y-c.MinY] == "" {
		return down, true
	}

	if currentSandPos.Y <= c.MinY {
		return end, true
	}

	if c.CaveMap[currentSandPos.X+1][currentSandPos.Y-c.MinY-1] == "" {
		return left, true
	}

	if currentSandPos.Y-c.MinY >= len(c.CaveMap[0])-1 {
		return end, true
	}

	if c.CaveMap[currentSandPos.X+1][currentSandPos.Y-c.MinY+1] == "" {
		return right, true
	}

	return -1, false
}

func StepSimulation2(c *Cave) bool {
	currentSandPos := c.DropSand

	direction, can := canMove2(currentSandPos, c)
	for can {
		switch direction {
		// Update the position
		case left:
			currentSandPos.X = currentSandPos.X + 1
			currentSandPos.Y = currentSandPos.Y - 1
		case right:
			currentSandPos.X = currentSandPos.X + 1
			currentSandPos.Y = currentSandPos.Y + 1
		case down:
			currentSandPos.X = currentSandPos.X + 1
		default:
			// Do nothing
		}

		// Check if the next sand grain is on the edge of the array
		// If so, we need to extend the array
		if currentSandPos.Y-c.MinY >= len(c.CaveMap[0])-1 {
			c.ExtendMapToRight()
		}

		if currentSandPos.Y <= c.MinY {
			c.ExtendMapToLeft()
		}

		direction, can = canMove2(currentSandPos, c)
	}
	c.CaveMap[currentSandPos.X][currentSandPos.Y-c.MinY] = "o"

	// For part 2
	if currentSandPos.X == c.DropSand.X && currentSandPos.Y == c.DropSand.Y {
		fmt.Println("[!] The simulation has stopped ! Sand blocks the source")
		return false
	}
	return true
}

func canMove2(currentSandPos utils.Coords, c *Cave) (Direction, bool) {
	if c.CaveMap[currentSandPos.X+1][currentSandPos.Y-c.MinY] == "" {
		return down, true
	}

	if c.CaveMap[currentSandPos.X+1][currentSandPos.Y-c.MinY-1] == "" {
		return left, true
	}

	if c.CaveMap[currentSandPos.X+1][currentSandPos.Y-c.MinY+1] == "" {
		return right, true
	}

	return -1, false
}
