package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func IsContiguous(springs string) bool {
	for _, elt := range springs {
		if elt == '.' {
			return false
		}
	}
	return true
}

func ParseToNumbers(numbers string) (lengths []int, err error) {
	split := strings.Split(numbers, ",")

	for _, nb := range split {
		var nbInt int
		nbInt, err = strconv.Atoi(nb)
		if err != nil {
			return
		}
		lengths = append(lengths, nbInt)
	}
	return
}

func ReadLines(path string) (lines []string, err error) {
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
	for s.Scan() {
		readLine := s.Text()
		lines = append(lines, strings.TrimSpace(readLine))
	}
	return
}
