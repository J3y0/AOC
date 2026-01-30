package utils

import (
	"strconv"
	"strings"
)

func ParseLines(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, "\n")
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
