package days

import (
	"main/utils"
	"strings"
)

type Day20 struct {
	enhancement []rune
	image       [][]rune
}

func (d *Day20) Parse(input string) error {
	input = strings.TrimSpace(input)
	split := strings.SplitN(input, "\n\n", 2)

	d.enhancement = []rune(split[0])

	lines := utils.ParseLines(split[1])
	d.image = make([][]rune, len(lines))
	for i, l := range lines {
		d.image[i] = []rune(l)
	}

	return nil
}

func (d *Day20) Part1() (int, error) {
	return expand(d.image, d.enhancement, 2)
}

func (d *Day20) Part2() (int, error) {
	return expand(d.image, d.enhancement, 50)
}

func expand(image [][]rune, enhancement []rune, steps int) (int, error) {
	// initial infinite background
	infinite := '.'
	for _ = range steps {
		image, infinite = generateOutput(image, enhancement, infinite)
	}

	lit := 0
	for _, r := range image {
		for _, c := range r {
			if c == '#' {
				lit++
			}
		}
	}
	return lit, nil
}

func generateOutput(image [][]rune, enhancement []rune, infinite rune) ([][]rune, rune) {
	/*
	 * # is the input image and . are the points added to the output image
	 * .....
	 * .###.
	 * .###.
	 * .###.
	 * .....
	 */
	h, w := len(image), len(image[0])
	output := make([][]rune, h+2)
	for i := -1; i < h+1; i++ {
		row := make([]rune, w+2)
		for j := -1; j < w+1; j++ {
			idx := getIndex(image, i, j, infinite)
			row[j+1] = enhancement[idx]
		}

		output[i+1] = row
	}

	var nextInfinite rune
	if infinite == '.' {
		nextInfinite = enhancement[0b000000000]
	} else {
		nextInfinite = enhancement[0b111111111]
	}

	return output, nextInfinite
}

func getIndex(image [][]rune, i, j int, infinite rune) int {
	idx := 0
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			idx <<= 1
			if pixelAt(image, i+di, j+dj, infinite) == '#' {
				idx |= 1
			}
		}
	}
	return idx
}

func pixelAt(image [][]rune, i, j int, infinite rune) rune {
	if i < 0 || j < 0 || i >= len(image) || j >= len(image[0]) {
		return infinite
	}
	return image[i][j]
}
