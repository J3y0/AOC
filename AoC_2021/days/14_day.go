package days

import (
	"fmt"
	"main/utils"
	"math"
	"strings"
)

type Day14 struct {
	template   string
	insertions map[string]byte
}

func (d *Day14) Parse(input string) error {
	split := strings.Split(input, "\n\n")
	d.template = split[0]
	d.insertions = make(map[string]byte)
	for _, l := range utils.ParseLines(split[1]) {
		var (
			pair     string
			toInsert byte
		)
		_, err := fmt.Sscanf(l, "%s -> %c", &pair, &toInsert)
		if err != nil {
			return err
		}

		d.insertions[pair] = toInsert
	}
	return nil
}

func (d *Day14) Part1() (int, error) {
	return polymerStep(10, d.template, d.insertions), nil
}

func (d *Day14) Part2() (int, error) {
	return polymerStep(40, d.template, d.insertions), nil
}

func polymerStep(step int, start string, insertions map[string]byte) int {
	pairMap := make(map[string]int)
	letterMap := make(map[byte]int)
	// init
	for i := range len(start) - 1 {
		pair := start[i : i+2]
		pairMap[pair] += 1
		letterMap[start[i]] += 1
	}
	letterMap[start[len(start)-1]] += 1

	for range step {
		nextPairMap := make(map[string]int)
		for pair, count := range pairMap {
			toInsert := insertions[pair]
			letterMap[toInsert] += count
			// insert new pairs
			newPair1 := string([]byte{pair[0], toInsert})
			newPair2 := string([]byte{toInsert, pair[1]})
			nextPairMap[newPair1] += count
			nextPairMap[newPair2] += count
		}
		pairMap = nextPairMap
	}

	mini := math.MaxInt
	maxi := 0
	for _, v := range letterMap {
		if v < mini {
			mini = v
		} else if v > maxi {
			maxi = v
		}
	}

	return maxi - mini
}
