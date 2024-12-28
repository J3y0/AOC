package utils

import (
	"strconv"
	"strings"
)

func FromLine(line, sep string) ([]int, error) {
	splitted := strings.Split(line, sep)
	res := make([]int, len(splitted))

	for i, nb := range splitted {
		conv, err := strconv.Atoi(nb)
		if err != nil {
			return nil, err
		}

		res[i] = conv
	}

	return res, nil
}

func OmitIndex[S ~[]T, T any](arr S, idx int) S {
	result := make(S, 0)
	for i, elt := range arr {
		if i == idx {
			continue
		}
		result = append(result, elt)
	}

	return result
}

func CopyArr(arr [][]rune) [][]rune {
	res := make([][]rune, len(arr))
	for i, elt := range arr {
		line := make([]rune, len(elt))
		for j := range elt {
			line[j] = elt[j]
		}

		res[i] = line
	}

	return res
}
