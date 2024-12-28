package days

import (
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"aoc/utils"
)

type Day21 struct {
	keypad     [][]rune
	dirpad     [][]rune
	codes      []string
	cachePaths map[string][]string // paths saved for going from x to y
}

func (d *Day21) Part1() (int, error) {
	sum := 0
	cache := make(map[string]int)
	for _, c := range d.codes {
		numKeypadPaths := d.getSeqs("A" + c)
		minLen := math.MaxInt
		for _, nk := range numKeypadPaths {
			minLen = min(minLen, d.recurseRobot(nk, 2, &cache))
		}

		codeNb, err := strconv.Atoi(c[:3])
		if err != nil {
			return 0, err
		}
		sum += minLen * codeNb
	}
	return sum, nil
}

func (d *Day21) Part2() (int, error) {
	sum := 0
	cache := make(map[string]int)
	for _, c := range d.codes {
		numKeypadPaths := d.getSeqs("A" + c)
		minLen := math.MaxInt
		for _, nk := range numKeypadPaths {
			minLen = min(minLen, d.recurseRobot(nk, 25, &cache))
		}

		codeNb, err := strconv.Atoi(c[:3])
		if err != nil {
			return 0, err
		}
		sum += minLen * codeNb
	}
	return sum, nil
}

func (d *Day21) Parse() error {
	content, err := os.ReadFile("./data/21_day.txt")
	if err != nil {
		return err
	}

	d.codes = strings.Split(strings.TrimSpace(string(content)), "\n")

	d.keypad = [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{'.', '0', 'A'},
	}

	d.dirpad = [][]rune{
		{'.', '^', 'A'},
		{'<', 'v', '>'},
	}

	d.cachePaths = make(map[string][]string)
	// Precompute all possibilities
	d.computeInit(d.keypad)
	d.computeInit(d.dirpad)
	return nil
}

func (d *Day21) recurseRobot(seq string, depth int, cache *map[string]int) int {
	seq = "A" + seq
	if l, ok := (*cache)[seq+","+strconv.Itoa(depth)]; ok {
		return l
	}

	if depth == 0 {
		length := len(seq) - 1
		(*cache)[seq+","+strconv.Itoa(depth)] = length
		return length
	}

	if depth == 1 {
		// 1 robot -> equivalent of getSeqs but w/o cartesianProduct as we only need length
		length := 0
		for i := 1; i < len(seq); i++ {
			key := d.hashKey(rune(seq[i-1]), rune(seq[i]))
			length += len(d.cachePaths[key][0]) // Take first elt as all have same length
		}

		(*cache)[seq+","+strconv.Itoa(depth)] = length
		return length
	}

	minLen := 0
	for i := 1; i < len(seq); i++ {
		key := d.hashKey(rune(seq[i-1]), rune(seq[i]))

		lengths := make([]int, 0)
		for _, possib := range d.cachePaths[key] {
			l := d.recurseRobot(possib, depth-1, cache)
			lengths = append(lengths, l)
		}

		minLen += slices.Min(lengths)
	}

	(*cache)[seq+","+strconv.Itoa(depth)] = minLen
	return minLen
}

func (d *Day21) getSeqs(code string) (paths []string) {
	for i := 1; i < len(code); i++ {
		x := code[i-1]
		y := code[i]
		key := d.hashKey(rune(x), rune(y))

		paths = d.cartesianProduct(paths, d.cachePaths[key])
	}
	return
}

func (d *Day21) cartesianProduct(a, b []string) []string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	res := make([]string, 0)
	for _, elt := range a {
		for _, elt2 := range b {
			res = append(res, elt+elt2)
		}
	}
	return res
}

func (d *Day21) hashKey(start, end rune) string {
	return string(start) + "," + string(end)
}

type scoreState struct {
	pos   utils.Pos
	score int
	path  string
}

type neighborMove struct {
	pos  utils.Pos
	move string
}

func (d *Day21) computeInit(keypad [][]rune) {
	flatten := make([]utils.Pos, 0)
	for i, r := range keypad {
		for j := range r {
			if keypad[i][j] == '.' {
				continue
			}
			flatten = append(flatten, utils.Pos{X: i, Y: j})
		}
	}

	for i := 0; i < len(flatten); i++ {
		for j := 0; j < len(flatten); j++ {
			start := flatten[i]
			end := flatten[j]

			minScore := math.MaxInt
			queue := []scoreState{
				{
					pos:   start,
					score: 0,
					path:  "",
				},
			}
			scores := make(map[utils.Pos]int)
			scores[start] = 0
			for len(queue) > 0 {
				curState := queue[0]
				queue = queue[1:]

				// end
				if keypad[curState.pos.X][curState.pos.Y] == keypad[end.X][end.Y] {
					if curState.score <= minScore {
						minScore = curState.score
						key := d.hashKey(keypad[start.X][start.Y], keypad[end.X][end.Y])
						if !slices.Contains(d.cachePaths[key], curState.path+"A") {
							d.cachePaths[key] = append(d.cachePaths[key], curState.path+"A")
						}
					}
					continue
				}

				neighbors := []neighborMove{
					{pos: utils.Pos{X: curState.pos.X + 1, Y: curState.pos.Y}, move: "v"},
					{pos: utils.Pos{X: curState.pos.X - 1, Y: curState.pos.Y}, move: "^"},
					{pos: utils.Pos{X: curState.pos.X, Y: curState.pos.Y - 1}, move: "<"},
					{pos: utils.Pos{X: curState.pos.X, Y: curState.pos.Y + 1}, move: ">"},
				}

				for _, n := range neighbors {
					if utils.OutOfGrid(n.pos, len(keypad), len(keypad[0])) {
						continue
					}
					if keypad[n.pos.X][n.pos.Y] == '.' {
						continue
					}
					curScore, ok := scores[n.pos]
					if ok && curState.score+1 > curScore {
						continue
					}
					scores[n.pos] = curState.score + 1

					queue = append(queue, scoreState{
						pos:   n.pos,
						score: curState.score + 1,
						path:  curState.path + n.move,
					})
				}
			}
		}
	}
}
