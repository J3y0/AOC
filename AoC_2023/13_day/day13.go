package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "./data/day13.txt"

	blocks, err := ReadBlocks(path)
	if err != nil {
		panic(err)
	}

	part1 := ComputeResult(blocks, 1)
	fmt.Println("Result for part 1:", part1)

	part2 := ComputeResult(blocks, 2)
	fmt.Println("Result for part 2:", part2)
}

func ComputeResult(blocks [][]string, part int) int {
	var (
		allowedFix = part - 1
		sumLines   int
		sumCols    int
		columns    []string
	)
blockLoop:
	for _, lines := range blocks {
		// Check reflection on lines
		for idxLine := 0; idxLine < len(lines)-1; idxLine++ {
			if HasReflection(lines, idxLine, allowedFix) {
				sumLines += idxLine + 1
				continue blockLoop
			}
		}

		columns = TransposeLines(lines)
		// Check reflection on columns
		for idxCol := 0; idxCol < len(columns)-1; idxCol++ {
			if HasReflection(columns, idxCol, allowedFix) {
				sumCols += idxCol + 1
				continue blockLoop
			}
		}
	}

	return sumLines*100 + sumCols
}

func HasReflection(block []string, idx int, allowedFix int) bool {
	offset := 0
	var totalErr = 0
	for idx+1+offset < len(block) && idx-offset >= 0 {
		nb1, _ := strconv.ParseInt(block[idx-offset], 2, 64)
		nb2, _ := strconv.ParseInt(block[idx+offset+1], 2, 64)
		totalErr += bitCount(nb1 ^ nb2)
		if totalErr > allowedFix {
			return false
		}
		offset += 1
	}

	// Need that for part 2 and not just return true.
	// You have to consider reflection lines where you fix the smudge
	if allowedFix == totalErr {
		return true
	}
	return false
}

func bitCount(number int64) int {
	binary := strconv.FormatInt(number, 2)
	return strings.Count(binary, "1")
}

func TransposeLines(lines []string) []string {
	columns := make([]string, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		var col strings.Builder
		for _, line := range lines {
			col.WriteByte(line[i])
		}
		columns[i] = col.String()
	}
	return columns
}

func ReadBlocks(path string) (blocks [][]string, err error) {
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

	s := bufio.NewScanner(file)
	tempLines := make([]string, 0)
	for s.Scan() {
		readLine := strings.TrimSpace(s.Text())
		readLine = strings.ReplaceAll(readLine, ".", "1")
		readLine = strings.ReplaceAll(readLine, "#", "0")
		if readLine == "" {
			blocks = append(blocks, tempLines)
			tempLines = make([]string, 0)
			continue
		}
		tempLines = append(tempLines, readLine)
	}

	blocks = append(blocks, tempLines)
	return
}
