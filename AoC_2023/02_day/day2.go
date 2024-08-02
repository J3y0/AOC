package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	Red int = iota
	Green
	Blue
)

func main() {
	file, err := os.Open("./data/day2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while opening file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	content, err := parseInput(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while parsing file: %v", err)
		os.Exit(1)
	}

	// Solve part 1
	part1 := possibleGames(content)
	fmt.Printf("Result for part 1: %d\n", part1)

	// Solve part 2
	part2 := solvePart2(content)
	fmt.Printf("Result for part 2: %d\n", part2)
}

func parseInput(r io.ReaderAt) (content string, err error) {
	// Read file
	buf := make([]byte, 1024)
	endFile := false
	offset := 0
	for !endFile {
		n, errFile := r.ReadAt(buf, int64(offset))
		if errFile == io.EOF {
			endFile = true
		} else if errFile != nil {
			err = errFile
			return
		}

		content += string(buf[:n])
		offset += n
	}

	return
}

func possible(line string) bool {
	var maxPerColor [3]int = [3]int{12, 13, 14} // Order: R, G, B
	sets := strings.Split(strings.Split(line, ": ")[1], "; ")

	for _, set := range sets {
		cubeItems := strings.Split(set, ", ")
		for _, item := range cubeItems {
			temp := strings.Split(item, " ")
			nb, _ := strconv.Atoi(temp[0])
			color := temp[1]
			if color == "red" && nb > maxPerColor[Red] {
				return false
			} else if color == "green" && nb > maxPerColor[Green] {
				return false
			} else if color == "blue" && nb > maxPerColor[Blue] {
				return false
			}
		}
	}

	return true
}

func possibleGames(content string) (sumIds int) {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		if possible(line) {
			sumIds += i + 1
		}
	}

	return
}

// Multiply minimum number of cubes for possible game together (only 3 colors)
func powerSet(numbersPerColor [3]int) int {
	return numbersPerColor[0] * numbersPerColor[1] * numbersPerColor[2]
}

func findMinPossible(line string) [3]int {
	var maxPerColor [3]int = [3]int{0, 0, 0} // Order: R, G, B
	sets := strings.Split(strings.Split(line, ": ")[1], "; ")

	for _, set := range sets {
		cubeItems := strings.Split(set, ", ")
		for _, item := range cubeItems {
			temp := strings.Split(item, " ")
			nb, _ := strconv.Atoi(temp[0])
			color := temp[1]
			if color == "red" && nb > maxPerColor[Red] {
				maxPerColor[Red] = nb
			} else if color == "green" && nb > maxPerColor[Green] {
				maxPerColor[Green] = nb
			} else if color == "blue" && nb > maxPerColor[Blue] {
				maxPerColor[Blue] = nb
			}
		}
	}

	return maxPerColor
}

func solvePart2(content string) (sumPowerSets int) {
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		numbersPerColor := findMinPossible(line)
		sumPowerSets += powerSet(numbersPerColor)
	}

	return
}
