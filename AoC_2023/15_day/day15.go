package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	Label       string
	FocalLength int
}

func main() {
	instructions, err := ParseInstructions("./data/day15.txt")
	if err != nil {
		panic(err)
	}

	// ------ Part 1 ------
	var (
		part1 int
	)
	for _, instr := range instructions {
		part1 += int(Hash(instr))
	}
	fmt.Println("Result for part 1:", part1)

	// ------ Part 2 ------
	var boxes = make(map[uint8][]Lens)
	for _, instr := range instructions {
		// = instruction
		if strings.Contains(instr, "=") {
			temp := strings.Split(instr, "=")
			label := temp[0]
			boxNb := Hash(label)
			focal, err := strconv.Atoi(temp[1])
			if err != nil {
				panic(err)
			}
			if arr, ok := boxes[boxNb]; ok {
				if idx, ok := Contain(arr, label); ok {
					// Replace len's focal length as it is already in the box
					boxes[boxNb][idx].FocalLength = focal
				} else {
					// Not in the box, then add the len into it
					boxes[boxNb] = append(boxes[boxNb], Lens{Label: label, FocalLength: focal})
				}
			} else {
				if boxes[boxNb] == nil {
					boxes[boxNb] = make([]Lens, 0)
				}
				// Else, add the len to the box
				boxes[boxNb] = append(boxes[boxNb], Lens{Label: label, FocalLength: focal})
			}
		} else { // - instruction
			temp := strings.Split(instr, "-")
			label := temp[0]
			boxNb := Hash(label)

			if arr, ok := boxes[boxNb]; ok {
				// If contain the len with label, remove it
				if idx, containBool := Contain(arr, label); containBool {
					// Remove len at idx
					boxes[boxNb] = append(arr[:idx], arr[idx+1:]...)
				}
			}
		}
	}

	var part2 int
	for boxNb, lens := range boxes {
		for i, lenObj := range lens {
			part2 += (int(boxNb) + 1) * (i + 1) * lenObj.FocalLength
		}
	}

	fmt.Println("Result for part 2:", part2)
}

func Hash(word string) uint8 {
	var result uint8
	for i := range word {
		result += word[i]
		result *= 17
	}
	return result
}

func ParseInstructions(path string) (instr []string, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	instr = strings.Split(strings.TrimSpace(string(data)), ",")
	return
}

func Contain(arr []Lens, label string) (int, bool) {
	for i, lenObj := range arr {
		if lenObj.Label == label {
			return i, true
		}
	}
	return 0, false
}
