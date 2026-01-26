package days

import (
	"fmt"
	"main/utils"
	"slices"
)

type token = rune

const (
	PAR_OPEN      token = '('
	PAR_CLOSE     token = ')'
	SQ_BRA_OPEN   token = '['
	SQ_BRA_CLOSE  token = ']'
	CUR_BRA_OPEN  token = '{'
	CUR_BRA_CLOSE token = '}'
	ANG_BRA_OPEN  token = '<'
	ANG_BRA_CLOSE token = '>'
)

type Day10 struct {
	lines []string
}

func (d *Day10) Parse(input string) error {
	d.lines = utils.ParseLines(input)
	return nil
}

func (d *Day10) Part1() (int, error) {
	tot := 0
	for _, l := range d.lines {
		_, illegal, corrupted := checkLine(l)
		if corrupted {
			tot += syntaxErrorScore(illegal)
		}
	}

	return tot, nil
}

func (d *Day10) Part2() (int, error) {
	scores := make([]int, 0)
	for _, l := range d.lines {
		unclosed, _, corrupted := checkLine(l)
		if !corrupted {
			completed := make([]token, len(unclosed))
			for i, token := range unclosed {
				idx := len(completed) - i - 1
				switch token {
				case PAR_OPEN:
					completed[idx] = PAR_CLOSE
				case SQ_BRA_OPEN:
					completed[idx] = SQ_BRA_CLOSE
				case CUR_BRA_OPEN:
					completed[idx] = CUR_BRA_CLOSE
				case ANG_BRA_OPEN:
					completed[idx] = ANG_BRA_CLOSE
				default:
					return 0, fmt.Errorf("unexpected token encountered: '%v'", token)
				}
			}

			scores = append(scores, autocompleteScore(completed))
		}
	}
	slices.Sort(scores)
	return scores[len(scores)/2], nil
}

func checkLine(line string) ([]token, token, bool) {
	opened := make([]token, 0)
	for _, r := range line {
		if r == PAR_OPEN || r == SQ_BRA_OPEN || r == ANG_BRA_OPEN || r == CUR_BRA_OPEN {
			opened = append(opened, r)
			continue
		}

		if len(opened) == 0 {
			break
		}

		last_opened := opened[len(opened)-1]

		var toCheck token
		switch last_opened {
		case PAR_OPEN:
			toCheck = PAR_CLOSE
		case SQ_BRA_OPEN:
			toCheck = SQ_BRA_CLOSE
		case CUR_BRA_OPEN:
			toCheck = CUR_BRA_CLOSE
		case ANG_BRA_OPEN:
			toCheck = ANG_BRA_CLOSE
		}

		if r == toCheck {
			opened = opened[:len(opened)-1]
		} else {
			return nil, r, true
		}
	}

	return opened, rune(0), false
}

func syntaxErrorScore(illegal token) int {
	switch illegal {
	case PAR_CLOSE:
		return 3
	case SQ_BRA_CLOSE:
		return 57
	case CUR_BRA_CLOSE:
		return 1197
	case ANG_BRA_CLOSE:
		return 25137
	default:
		return 0
	}
}

func autocompleteScore(seq []token) int {
	scores := map[token]int{
		PAR_CLOSE:     1,
		SQ_BRA_CLOSE:  2,
		CUR_BRA_CLOSE: 3,
		ANG_BRA_CLOSE: 4,
	}

	tot := 0
	for _, r := range seq {
		tot *= 5
		tot += scores[r]
	}

	return tot
}
