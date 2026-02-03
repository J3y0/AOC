package days

import (
	"fmt"
	"main/utils"
)

type Day21 struct {
	start1 int
	start2 int
}

func (d *Day21) Parse(input string) error {
	lines := utils.ParseLines(input)
	_, err := fmt.Sscanf(lines[0], "Player 1 starting position: %d", &d.start1)
	if err != nil {
		return err
	}

	_, err = fmt.Sscanf(lines[1], "Player 2 starting position: %d", &d.start2)
	return err
}

func (d *Day21) Part1() (int, error) {
	p1, p2, score1, score2 := d.start1, d.start2, 0, 0
	dice := 0
	for {
		p1, score1, dice = deterministicTurn(p1, score1, dice)
		if score1 >= 1000 {
			return score2 * dice, nil
		}

		p2, score2, dice = deterministicTurn(p2, score2, dice)
		if score2 >= 1000 {
			return score1 * dice, nil
		}
	}
}

func (d *Day21) Part2() (int, error) {
	// { sum: freq }
	diceFreq := make(map[int]int)
	for i := range 3 {
		for j := range 3 {
			for k := range 3 {
				diceFreq[i+j+k+3]++
			}
		}
	}

	cache := make(map[[5]int][2]int)
	p1, p2 := countUniverses(diceFreq, d.start1, d.start2, 0, 0, 1, cache)
	return max(p1, p2), nil
}

func deterministicTurn(pos, score, dice int) (npos, nscore, ndice int) {
	// (dice)+(dice+1)+(dice+2) = 3*dice + 6
	rollsSum := 3*utils.ShiftedMod(dice, 100) + 6
	npos = utils.ShiftedMod(pos+rollsSum, 10)
	nscore = score + npos
	ndice = dice + 3
	return
}

// p1Turn = 1 if it is p1 turn to play, 0 otherwise
func countUniverses(diceFreq map[int]int, pos1, pos2, s1, s2, p1Turn int, cache map[[5]int][2]int) (int, int) {
	if s1 >= 21 {
		return 1, 0
	}
	if s2 >= 21 {
		return 0, 1
	}

	key := [5]int{pos1, pos2, s1, s2, p1Turn}
	if entry, ok := cache[key]; ok {
		return entry[0], entry[1]
	}

	p1Win, p2Win := 0, 0
	for sum, freq := range diceFreq {
		var cur1Win, cur2Win int
		if p1Turn == 1 {
			npos1 := utils.ShiftedMod(pos1+sum, 10)
			cur1Win, cur2Win = countUniverses(diceFreq, npos1, pos2, s1+npos1, s2, 0, cache)
		} else {
			npos2 := utils.ShiftedMod(pos2+sum, 10)
			cur1Win, cur2Win = countUniverses(diceFreq, pos1, npos2, s1, s2+npos2, 1, cache)
		}
		p1Win += cur1Win * freq
		p2Win += cur2Win * freq
	}

	cache[key] = [2]int{p1Win, p2Win}
	return p1Win, p2Win
}
