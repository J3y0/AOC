package days

import (
	"fmt"
	"main/utils"
	"slices"
	"strings"
)

type fold struct {
	dir rune
	val int
}

type Day13 struct {
	dots  [][2]int
	folds []fold
}

func (d *Day13) Parse(input string) error {
	split := strings.SplitN(input, "\n\n", 2)

	// Parse dots
	dot_lines := utils.ParseLines(split[0])
	d.dots = make([][2]int, len(dot_lines))
	for i, l := range dot_lines {
		var x, y int
		_, err := fmt.Sscanf(l, "%d,%d", &x, &y)
		if err != nil {
			return err
		}
		d.dots[i] = [2]int{x, y}
	}

	// Parse fold instructions
	fold_lines := utils.ParseLines(split[1])
	d.folds = make([]fold, len(fold_lines))
	for i, l := range fold_lines {
		var f fold
		_, err := fmt.Sscanf(l, "fold along %c=%d", &f.dir, &f.val)
		if err != nil {
			return err
		}
		d.folds[i] = f
	}

	return nil
}

func (d *Day13) Part1() (int, error) {
	dots, err := applyFold(d.dots, d.folds[0])
	if err != nil {
		return 0, nil
	}
	return len(dots), nil
}

func (d *Day13) Part2() (int, error) {
	var err error
	for _, inst := range d.folds {
		d.dots, err = applyFold(d.dots, inst)
		if err != nil {
			return 0, nil
		}
	}

	// answer is not a number: read the 8 capital letters displayed
	printFolds(d.dots)
	return 0, nil
}

func applyFold(dots [][2]int, instruction fold) ([][2]int, error) {
	res := make([][2]int, 0)
	switch instruction.dir {
	case 'x':
		for _, d := range dots {
			// not affected by fold
			if d[0] <= instruction.val {
				res = append(res, d)
				continue
			}

			delta := d[0] - instruction.val
			new_d := [2]int{instruction.val - delta, d[1]}
			res = append(res, new_d)
		}
	case 'y':
		for _, d := range dots {
			// not affected by fold
			if d[1] <= instruction.val {
				res = append(res, d)
				continue
			}

			delta := d[1] - instruction.val
			new_d := [2]int{d[0], instruction.val - delta}
			res = append(res, new_d)
		}
	default:
		return nil, fmt.Errorf("unknown fold direction: %c", instruction.dir)
	}
	res = utils.Duplicates(res)
	return res, nil
}

func printFolds(dots [][2]int) {
	maxX := 0
	maxY := 0
	for _, d := range dots {
		if d[0] > maxX {
			maxX = d[0]
		}
		if d[1] > maxY {
			maxY = d[1]
		}
	}
	var sb strings.Builder
	for j := range maxY + 1 {
		for i := range maxX + 1 {
			if slices.Contains(dots, [2]int{i, j}) {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}

	fmt.Println(sb.String())
}
