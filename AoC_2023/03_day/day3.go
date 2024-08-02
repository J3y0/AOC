package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, _ := readFile("./data/day3.txt")

	part1 := sumMotorPart(content)
	fmt.Printf("Result for Part 1: %d\n", part1)

	part2 := sumGearRatio(content)
	fmt.Printf("Result for Part 2: %d\n", part2)
}

func readFile(path string) (content string, err error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while opening file: %v", err)
		return
	}
	// Read file
	buf := make([]byte, 1024)
	endFile := false
	offset := 0
	for !endFile {
		n, errFile := file.ReadAt(buf, int64(offset))
		if errFile == io.EOF {
			endFile = true
		} else if errFile != nil {
			fmt.Fprintf(os.Stderr, "Error received while reading file: %v", err)
			err = errFile
			return
		}

		content += string(buf[:n])
		offset += n
	}
	return
}

func sumMotorPart(content string) int {
	lines := strings.Split(content, "\n")

	var sum int
	var window []string
	for i, line := range lines {
		if i == 0 {
			window = lines[i : i+2]
		} else if i == len(lines)-1 {
			window = lines[i-1 : i+1]
		} else {
			window = lines[i-1 : i+2]
		}
		numbers_idx := numbersIndexes(line)
		for _, nb_range := range numbers_idx {
			start_idx := nb_range[0]
			end_idx := nb_range[1]
			if _, _, ok := isAdjacent(start_idx, end_idx, window, `[^\.0-9]`); ok {
				nb, _ := strconv.Atoi(line[start_idx:end_idx])
				sum += nb
			}
		}
	}

	return sum
}

func numbersIndexes(line string) [][]int {
	regex := regexp.MustCompile(`[0-9]+`) // Find numbers
	indexes := regex.FindAllSubmatchIndex([]byte(line), -1)
	return indexes
}

// Check if number starting at index start_idx and ending at end_idx
// has a adjacent symbols in window which contains adjacent lines
func isAdjacent(start_idx int, end_idx int, window []string, pattern_regex string) (int, int, bool) {
	symbol_regex := regexp.MustCompile(pattern_regex)

	var (
		before int = start_idx - 1
		after  int = end_idx + 1
	)
	for offset, line := range window {
		// Make sure you don't exceed line size in any way
		if before < 0 {
			before = 0
		}
		if after > len(line)-1 {
			after = len(line) - 1
		}
		// If match, there is a symbol adjacent to our number
		subline_idx := symbol_regex.FindIndex([]byte(line[before:after]))
		if subline_idx != nil {
			return before + subline_idx[0], offset, true
		}
	}

	return -1, -1, false
}

func gearRatio(nb1, nb2 int) int {
	return nb1 * nb2
}

type GearPosition struct {
	X int
	Y int
}

func sumGearRatio(content string) int {
	lines := strings.Split(content, "\n")
	var (
		sum           int
		window        []string
		window_offset int
	)
	gear_map := make(map[GearPosition][]int, 0)
	for i, line := range lines {
		// Define adjacent lines to current line as
		// window (one before and one after when correctly defined)
		if i == 0 {
			window = lines[i : i+2]
			window_offset = 0
		} else if i == len(lines)-1 {
			window = lines[i-1 : i+1]
			window_offset = -1
		} else {
			window = lines[i-1 : i+2]
			window_offset = -1
		}
		// Find numbers on line
		numbers_idx := numbersIndexes(line)
		for _, nb_range := range numbers_idx {
			start_idx := nb_range[0]
			end_idx := nb_range[1]

			if line_idx, match_offset, ok := isAdjacent(start_idx, end_idx, window, `\*`); ok {
				// Retrieve number
				nb, _ := strconv.Atoi(line[start_idx:end_idx])
				// Compute position of potential gear
				gear_pos := GearPosition{X: i + match_offset + window_offset, Y: line_idx}
				// Add adjacent number to map
				if numbers, ok := gear_map[gear_pos]; ok {
					numbers = append(numbers, nb)
					gear_map[gear_pos] = numbers
				} else {
					gear_map[gear_pos] = []int{nb}
				}
			}
		}
	}

	for _, adj_nb_to_gear := range gear_map {
		// If 2 adjacent numbers, compute gear ratio
		if len(adj_nb_to_gear) == 2 {
			sum += gearRatio(adj_nb_to_gear[0], adj_nb_to_gear[1])
		}
	}

	return sum
}
