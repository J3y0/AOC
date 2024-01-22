package main

import (
	"bufio"
	"os"
	"strings"
)

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
