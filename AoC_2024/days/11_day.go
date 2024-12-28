package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day11 struct {
	initStones map[int]int
}

func (d *Day11) Part1() (int, error) {
	return d.blinkTimes(25), nil
}

func (d *Day11) Part2() (int, error) {
	return d.blinkTimes(75), nil
}

func (d *Day11) Parse() error {
	content, err := os.ReadFile("./data/11_day.txt")
	if err != nil {
		return err
	}

	values := strings.Split(strings.TrimSpace(string(content)), " ")

	d.initStones = make(map[int]int)
	for _, s := range values {
		parsed, err := strconv.Atoi(s)
		if err != nil {
			return err
		}

		if _, ok := d.initStones[parsed]; ok {
			d.initStones[parsed] += 1
		} else {
			d.initStones[parsed] = 1
		}
	}

	return nil
}

func (d *Day11) applyRules(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	str := fmt.Sprintf("%d", stone)
	if len(str)%2 == 0 {
		half, _ := strconv.Atoi(str[:len(str)/2])
		secHalf, _ := strconv.Atoi(str[len(str)/2:])

		return []int{half, secHalf}
	}

	return []int{2024 * stone}
}

func (d *Day11) blinkTimes(times int) int {
	stoneMap := d.initStones
	for range times {
		nextMap := make(map[int]int)
		for stone, count := range stoneMap {
			created := d.applyRules(stone)
			// Add new stones to nextMap
			for _, s := range created {
				if _, ok := nextMap[s]; ok {
					nextMap[s] += count
				} else {
					nextMap[s] = count
				}
			}
		}
		stoneMap = nextMap
	}

	// Count total
	tot := 0
	for _, v := range stoneMap {
		tot += v
	}
	return tot
}
