package days

import (
	"main/utils"
	"os"
	"strconv"
	"strings"
)

type Day6 struct {
	fishes []int
}

const nbFatherReproduce = 6
const nbChildReproduce = 8

func (d *Day6) Part1() (int, error) {
	fishes, err := parseFishes("./input/06_day.txt")
	if err != nil {
		return 0, err
	}
	d.fishes = fishes

	return countOffspring(d.fishes, 80), nil
}

func (d *Day6) Part2() (int, error) {
	if len(d.fishes) == 0 {
		fishes, err := parseFishes("./input/06_day.txt")
		if err != nil {
			return 0, err
		}
		d.fishes = fishes
	}

	return countOffspring(d.fishes, 256), nil
}

func parseFishes(path string) ([]int, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	fishes := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), ",")
	parsedFishes := make([]int, nbChildReproduce+1)
	for _, fish := range fishes {
		intFish, err := strconv.Atoi(fish)
		parsedFishes[intFish] += 1
		if err != nil {
			return nil, err
		}
	}

	return parsedFishes, nil
}

func countOffspring(fishes []int, steps int) int {
	for range steps {
		newborns := fishes[0]
		fishes = append(fishes[1:], newborns)
		fishes[nbFatherReproduce] += newborns
	}

	return utils.SumArray(fishes)
}
