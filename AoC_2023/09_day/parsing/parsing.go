package parsing

import (
	"strconv"
	"strings"
)

func Parse(data []byte) (linesNumbers [][]int, err error) {
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		numbers := strings.Split(line, " ")

		var lineNumbers []int
		for _, nb := range numbers {
			var number int
			number, err = strconv.Atoi(nb)
			if err != nil {
				return
			}
			lineNumbers = append(lineNumbers, number)
		}

		linesNumbers = append(linesNumbers, lineNumbers)
	}

	return
}
