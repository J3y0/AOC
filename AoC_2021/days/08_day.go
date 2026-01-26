package days

import (
	"main/utils"
	"strings"
)

type Entry struct {
	signals []string
	outputs []string
}

// Digits binary representation: 0b(a-b-c-d-e-f-g) - 1 if used for representation, 0 otherwise
const (
	g_seg = 1 << iota
	f_seg
	e_seg
	d_seg
	c_seg
	b_seg
	a_seg
)

const (
	d0 = a_seg + b_seg + c_seg + e_seg + f_seg + g_seg         // 0b1110111
	d1 = c_seg + f_seg                                         // 0b0010010
	d2 = a_seg + c_seg + d_seg + e_seg + g_seg                 // 0b1011101
	d3 = a_seg + c_seg + d_seg + f_seg + g_seg                 // 0b1011011
	d4 = b_seg + c_seg + d_seg + f_seg                         // 0b0111010
	d5 = a_seg + b_seg + d_seg + f_seg + g_seg                 // 0b1101011
	d6 = a_seg + b_seg + d_seg + e_seg + f_seg + g_seg         // 0b1101111
	d7 = a_seg + c_seg + f_seg                                 // 0b1010010
	d8 = a_seg + b_seg + c_seg + d_seg + e_seg + f_seg + g_seg // 0b1111111
	d9 = a_seg + b_seg + c_seg + d_seg + f_seg + g_seg         // 0b1111011
)

type Day8 struct {
	entries []Entry
}

func (d *Day8) Parse(input string) error {
	lines := utils.ParseLines(input)
	entries := make([]Entry, len(lines))
	for i, line := range lines {
		entry := strings.Split(line, " | ")
		signals := entry[0]
		outputs := entry[1]
		entries[i] = Entry{
			signals: strings.Split(signals, " "),
			outputs: strings.Split(outputs, " "),
		}
	}
	d.entries = entries

	return nil
}

func (d *Day8) Part1() (int, error) {
	count := 0
	for _, entry := range d.entries {
		for _, out := range entry.outputs {
			nbSegments := len(out)
			// check number is: 1 || 4 || 7 || 8
			if nbSegments == 2 || nbSegments == 4 || nbSegments == 3 || nbSegments == 7 {
				count++
			}
		}
	}

	return count, nil
}

/*
* Use numbers: 1, 4, 7, 8, 9, 3, 6
* To determine segment:
*   - a: diff 1 and 7
*   - e: diff 8 and 9 (9 can be determined as it shares all 4's segments and has a 6 segments)
*   - g: diff between 9 and 4 + seg a
*	- d: 3 (can be determined) which is composed of segs. a, g, d and 1's segments
*   - b: nb 9 and remove 1's segments, a, d, g
*   - f: nb 6 and remove a, d, g, b, e
*   - c: 1 and remove f
 */
func (d *Day8) Part2() (int, error) {
	tot := 0
	for _, entry := range d.entries {
		// { real: encoded }
		segmap := make(map[rune]rune)
		// { encoded: real } encoded is the value mixed and real is the segment represented by encoded
		rev_segmap := make(map[rune]rune)

		// known numbers
		four_nb := findPerLength(4, entry.signals)[0]
		eight_nb := findPerLength(7, entry.signals)[0]
		one_nb := findPerLength(2, entry.signals)[0]
		seven_nb := findPerLength(3, entry.signals)[0]

		found_a_seg := diffSegments(seven_nb, one_nb)
		segmap['a'] = found_a_seg
		rev_segmap[found_a_seg] = 'a'

		// find number 9: len of 6 and shares all 4's segments
		var nine_nb string
		for _, s := range findPerLength(6, entry.signals) {
			if containSegments(s, four_nb) {
				nine_nb = s
				break
			}
		}

		found_e_seg := diffSegments(eight_nb, nine_nb)
		segmap['e'] = found_e_seg
		rev_segmap[found_e_seg] = 'e'

		found_g_seg := diffSegments(nine_nb, four_nb+string(segmap['a']))
		segmap['g'] = found_g_seg
		rev_segmap[found_g_seg] = 'g'

		// find number 3: len of 5 and shares all 1's segments
		var three_nb string
		for _, s := range findPerLength(5, entry.signals) {
			if containSegments(s, one_nb) {
				three_nb = s
				break
			}
		}
		found_d_seg := diffSegments(three_nb, one_nb+string(segmap['a'])+string(segmap['g']))
		segmap['d'] = found_d_seg
		rev_segmap[found_d_seg] = 'd'

		found_b_seg := diffSegments(nine_nb, one_nb+string(segmap['a'])+string(segmap['g'])+string(segmap['d']))
		segmap['b'] = found_b_seg
		rev_segmap[found_b_seg] = 'b'

		// find number 6: len of 6 and does not share all 1's segments
		var six_nb string
		for _, s := range findPerLength(6, entry.signals) {
			if !containSegments(s, one_nb) {
				six_nb = s
				break
			}
		}

		found_f_seg := diffSegments(six_nb, string(segmap['a'])+string(segmap['b'])+string(segmap['d'])+string(segmap['g'])+string(segmap['e']))
		segmap['f'] = found_f_seg
		rev_segmap[found_f_seg] = 'f'

		found_c_seg := diffSegments(one_nb, string(segmap['f']))
		segmap['c'] = found_c_seg
		rev_segmap[found_c_seg] = 'c'

		outNumber := 0
		for i, out := range entry.outputs {
			outNumber += utils.PowInt(10, len(entry.outputs)-i-1) * retrieveNumber(out, rev_segmap)
		}

		tot += outNumber
	}
	return tot, nil
}

// diff more - less
//
//	return the segment which is in more and not in less
func diffSegments(more, less string) rune {
	for _, r := range more {
		if !strings.Contains(less, string(r)) {
			return r
		}
	}

	return rune(0)
}

func containSegments(s, sub string) bool {
	for _, c := range sub {
		if !strings.Contains(s, string(c)) {
			return false
		}
	}

	return true
}

func findPerLength(length int, signals []string) []string {
	res := make([]string, 0)
	for _, s := range signals {
		if length == len(s) {
			res = append(res, s)
		}
	}

	return res
}

func retrieveNumber(str string, rev_segmap map[rune]rune) int {
	res := 0
	for _, r := range str {
		decoded := rev_segmap[r]
		var toAdd int
		switch decoded {
		case 'a':
			toAdd = a_seg
		case 'b':
			toAdd = b_seg
		case 'c':
			toAdd = c_seg
		case 'd':
			toAdd = d_seg
		case 'e':
			toAdd = e_seg
		case 'f':
			toAdd = f_seg
		case 'g':
			toAdd = g_seg
		}
		res += toAdd
	}
	return matchNumber(res)
}

func matchNumber(nb int) int {
	switch nb {
	case d0:
		return 0
	case d1:
		return 1
	case d2:
		return 2
	case d3:
		return 3
	case d4:
		return 4
	case d5:
		return 5
	case d6:
		return 6
	case d7:
		return 7
	case d8:
		return 8
	case d9:
		return 9
	}
	return -1
}
