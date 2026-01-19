package days

import (
	"main/utils"
	"strconv"
	"strings"
)

type Day6 struct {
	fishes []int
}

func (d *Day6) Parse(input string) error {
	input = strings.Trim(input, "\n")
	fishes := strings.Split(input, ",")
	parsedFishes := make([]int, nbChildReproduce+1)
	for _, fish := range fishes {
		intFish, err := strconv.Atoi(fish)
		parsedFishes[intFish] += 1
		if err != nil {
			return err
		}
	}

	d.fishes = parsedFishes

	return nil
}

const nbFatherReproduce = 6
const nbChildReproduce = 8

func (d *Day6) Part1() (int, error) {
	return countOffspring(d.fishes, 80), nil
}

func (d *Day6) Part2() (int, error) {
	return countOffspring(d.fishes, 256), nil
}

func countOffspring(fishes []int, steps int) int {
	for range steps {
		newborns := fishes[0]
		fishes = append(fishes[1:], newborns)
		fishes[nbFatherReproduce] += newborns
	}

	return utils.SumArray(fishes)
}
