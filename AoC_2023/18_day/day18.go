package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Area(lines []string, part int) int {
	var (
		currentPos = [2]int{0, 0}
		area       int
		perimeter  int
		steps      int
	)
	for _, line := range lines {
		split := strings.Split(line, " ")
		var dir string
		if part == 1 {
			dir = split[0]
			steps, _ = strconv.Atoi(split[1])
		} else if part == 2 {
			hexValue := split[2]
			hexValue = hexValue[2 : len(hexValue)-1]
			dir = string(hexValue[len(hexValue)-1])
			convertedInt64, _ := strconv.ParseInt(hexValue[:len(hexValue)-1], 16, strconv.IntSize)
			steps = int(convertedInt64)
		}

		var newPos = currentPos
		switch dir {
		case "R", "0":
			newPos[1] += steps
		case "L", "2":
			newPos[1] -= steps
		case "D", "1":
			newPos[0] += steps
		case "U", "3":
			newPos[0] -= steps
		default:
			panic("Unknown direction")
		}

		perimeter += steps
		area += newPos[1]*currentPos[0] - newPos[0]*currentPos[1] // Shoelace
		currentPos = newPos
	}

	/*
		More explanations on the result.
		We can use Shoelace formula and combine it with Pick's Theorem.
		 - Shoelace formula: A = 1/2*Abs(Sum{x_i*y_{i+1} - x_{i+1}*y_i})
		 - Pick's Theorem: A = i + b/2 - 1 with i the number of inner points and b the number of points forming the polygon's perimeter

		Now, our points are aligned with the center of a 'box' but Shoelace needs to be applied on a border which is including
		all our boxes. So applying directly Shoelace won't work.

		As the area is defined as the number of boxes in this problem, we can consider the total area is:
			A_tot = i + b, the sum of the inner boxes and the boxes forming the perimeter
		It's easy to determine b as this is the sum of all the move operation of the input file.

		As for i, we will use Pick's formula: i = A - b/2 + 1

		Now, in that case, A can be computed with Shoelace as this is the true area of the polygon. (polygon which is correctly defined
		with the coordinates obtained above). Indeed, this area contains the same number of inner points. (each point corresponds to a box
		because each point is the center of the box)

		Finally, we have the relation: A_tot = b + i = A + b/2 + 1 = (b + Abs(Sum{x_i*y_{i+1} - x_{i+1}*y_i}))/2 + 1
	*/
	return (Abs(area)+perimeter)/2 + 1
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	lines := make([]string, 0)
	for s.Scan() {
		readLine := strings.TrimSpace(s.Text())
		lines = append(lines, readLine)
	}

	return lines, nil
}

func main() {
	lines, err := ReadLines("./data/day18.txt")
	if err != nil {
		panic(err)
	}

	tStart := time.Now()
	part1 := Area(lines, 1)
	fmt.Println("Part 1:", part1, time.Since(tStart))

	part2 := Area(lines, 2)
	fmt.Println("Part 2:", part2, time.Since(tStart))
}
