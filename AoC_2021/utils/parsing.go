package utils

import (
	"os"
	"strconv"
	"strings"
)

func ParseLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	return lines, nil
}

func ParseLineToIntArray(line string, sep string) ([]int, error) {
	numbersStr := strings.Split(line, sep)
	numbers := make([]int, len(numbersStr))
	for i, nbStr := range numbersStr {
		nb, err := strconv.Atoi(nbStr)
		if err != nil {
			return nil, err
		}
		numbers[i] = nb
	}

	return numbers, nil
}
