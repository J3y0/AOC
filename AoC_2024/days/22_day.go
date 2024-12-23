package days

import (
	"os"
	"strings"

	"aoc/utils"
)

type Day22 struct {
	secrets []int
	lasts   map[int][]int
}

func (d *Day22) Part1() (int, error) {
	sum := 0
	for _, s := range d.secrets {
		sum += d.nextSecret(s, 2000)
	}
	return sum, nil
}

func (d *Day22) Part2() (int, error) {
	if len(d.lasts[d.secrets[0]]) == 0 {
		d.Part1() // Register all last digits of every changes (for each secret)
	}

	seq := make(map[[4]int]int)
	for _, s := range d.secrets {
		seen := make(map[[4]int]bool)
		for i := 0; i < len(d.lasts[s])-4; i++ {
			a := d.lasts[s][i : i+5]
			sequence := [4]int{a[1] - a[0], a[2] - a[1], a[3] - a[2], a[4] - a[3]}

			// If already bought from buyer, you can't anymore
			if seen[sequence] {
				continue
			}
			seen[sequence] = true

			// Keep track of bananas bought
			if _, ok := seq[sequence]; !ok {
				seq[sequence] = 0
			}
			seq[sequence] += a[4]
		}
	}

	// Find max bananas
	maxi := 0
	for _, v := range seq {
		if v > maxi {
			maxi = v
		}
	}
	return maxi, nil
}

func (d *Day22) Parse() error {
	content, err := os.ReadFile("./data/22_day.txt")
	if err != nil {
		return err
	}

	d.secrets, err = utils.FromLine(strings.TrimSpace(string(content)), "\n")
	d.lasts = make(map[int][]int)
	return err
}

func (d *Day22) nextSecret(secret, times int) int {
	mod := 16777216 - 1 // 0xffffff
	current := secret
	next := 0
	for range times {
		lastDigit := utils.PosMod(current, 10)
		if _, ok := d.lasts[secret]; !ok {
			d.lasts[secret] = make([]int, 0)
		}
		d.lasts[secret] = append(d.lasts[secret], lastDigit)

		tmp := current << 6  // Mul 64
		next = tmp ^ current // Mix
		next = next & mod    // Prune

		tmp = next >> 5   // Div 32
		next = tmp ^ next // Mix
		next = next & mod // Prune

		tmp = next << 11  // Mul 2048
		next = tmp ^ next // Mix
		next = next & mod // Prune

		current = next
	}

	return next
}
