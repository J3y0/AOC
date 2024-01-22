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
	file, err := os.Open("./data/day1.txt")
	if err != nil {
		fmt.Printf("Error while opening file\n")
	}
	defer file.Close()

	var instructions string
	instructions, err = readFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Received following error while reading file: %v", err)
		os.Exit(1)
	}

	var part1 int = 0
	part1, err = solvePart1(instructions)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Received following error while solving part 1: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Result for part 1: %d\n", part1)

	var part2 int = 0
	part2, err = solvePart2(instructions)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Received following error while solving part 2: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Result for part 2: %d\n", part2)
}

func readFile(r io.ReaderAt) (content string, err error) {
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

func solvePart1(instructions string) (calibrationValue int, err error) {
	lines := strings.Split(instructions, "\n")
	regNb := regexp.MustCompile(`[0-9]`)

	var toAdd int
	for _, line := range lines {
		allDigits := regNb.FindAllString(line, -1)
		if allDigits != nil {
			toAdd, err = strconv.Atoi(allDigits[0] + allDigits[len(allDigits)-1])
			if err != nil {
				return
			}
			calibrationValue += toAdd
		}
	}

	return
}

func solvePart2(instructions string) (int, error) {
	convertionMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	var newInstr string = instructions
	for key, value := range convertionMap {
		newInstr = strings.ReplaceAll(newInstr, key, value)
	}

	return solvePart1(newInstr)
}
