package days

import (
	"errors"
	"fmt"
	"main/utils"
	"slices"
	"strconv"
	"strings"
)

type registers map[string]int

type operator string

const (
	Inp operator = "inp"
	Add operator = "add"
	Mul operator = "mul"
	Eql operator = "eql"
	Div operator = "div"
	Mod operator = "mod"
)

const MonadLength int = 14

// Extracted constants for all 14 sub-programs
// A sub-program is the list of instructions delimited by `inp` instruction
var Ky = [MonadLength]int{7, 8, 10, 4, 4, 6, 11, 13, 1, 8, 4, 13, 4, 14}
var Kx = [MonadLength]int{12, 13, 13, -2, -10, 13, -14, -5, 15, 15, -14, 10, -14, -5}

// DivZ tells if the sub-program divides Z by 26
// It is important to notice that there are an equal number of div by 26 sub-programs than div by 1 sub-progams.
// It allows us to simlpify the search of valid model numbers: as the final z needs to be 0:
//   - all div by 1 sub-programs increases z by mul it by 26
//   - all div by 26 sub-programs reduces z by div 26
//
// If the final result is 0, then it must be an equal number of /26 and *26 as we start with z=0
// As the division needs another condition explicited by:
//   - w_i = Kx_i + z_{i-1} % 26   where w_i is the i-st digit of the model number (14 digits total)
//     and z_{i-1} depends on w_{i-1}, we can deduce conditions on numbers that constitutes the model number
var DivZ = [MonadLength]bool{false, false, false, true, true, false, true, true, false, false, true, false, true, true}

var allOperators = []operator{Inp, Add, Mod, Mul, Eql, Div}

func isValidOperator(op string) bool {
	return slices.Contains(allOperators, operator(op))
}

type instruction struct {
	op       operator
	operands []string
}

func (i instruction) String() string {
	return fmt.Sprintf("%s: %s", i.op, strings.Join(i.operands, ", "))
}

type Day24 struct {
	instructions []instruction
}

func (d *Day24) Parse(input string) error {
	lines := utils.ParseLines(input)
	d.instructions = make([]instruction, 0, len(lines))
	for _, l := range lines {
		ops := strings.SplitN(l, " ", 4)
		if !isValidOperator(ops[0]) {
			return fmt.Errorf("unknwon operator: %s", ops[0])
		}
		operator := operator(ops[0])
		operands := make([]string, 0, 2)
		if operator == Inp {
			operands = append(operands, ops[1])
		} else {
			operands = append(operands, ops[1])
			operands = append(operands, ops[2])
		}
		d.instructions = append(d.instructions, instruction{
			op:       operator,
			operands: operands,
		})
	}

	return nil
}

func (d *Day24) Part1() (int, error) {
	largest, ok := d.findLargestModelNumber(0, 0, []int{0})
	if !ok {
		return 0, errors.New("not found")
	}
	return largest, nil
}

func (d *Day24) Part2() (int, error) {
	lowest, ok := d.findLowestModelNumber(0, 0, []int{0})
	if !ok {
		return 0, errors.New("not found")
	}
	return lowest, nil
}

// helper function to understand what the program does and to see what values
// registers hold
func (d *Day24) runMonad(modelNumber int, r registers) registers {
	inpCount := 0
	for _, ins := range d.instructions {
		switch ins.op {
		case Inp:
			dst := ins.operands[0]
			inp := utils.Mod(modelNumber/utils.PowInt(10, MonadLength-inpCount-1), 10)
			inpCount++
			r[dst] = inp
		case Add:
			dst := ins.operands[0]
			src := ins.operands[1]
			conv, err := strconv.Atoi(src)
			if err != nil {
				// other register
				r[dst] += r[src]
			} else {
				r[dst] += conv
			}
		case Mul:
			dst := ins.operands[0]
			src := ins.operands[1]
			conv, err := strconv.Atoi(src)
			if err != nil {
				// other register
				r[dst] *= r[src]
			} else {
				r[dst] *= conv
			}
		case Mod:
			dst := ins.operands[0]
			src := ins.operands[1]
			conv, err := strconv.Atoi(src)
			if err != nil {
				// other register
				r[dst] = utils.Mod(r[dst], r[src])
			} else {
				r[dst] = utils.Mod(r[dst], conv)
			}
		case Div:
			dst := ins.operands[0]
			src := ins.operands[1]
			conv, err := strconv.Atoi(src)
			if err != nil {
				// other register
				r[dst] /= r[src]
			} else {
				r[dst] /= conv
			}
		case Eql:
			dst := ins.operands[0]
			src := ins.operands[1]
			conv, err := strconv.Atoi(src)
			eq := false
			if err != nil {
				// other register
				eq = r[dst] == r[src]
			} else {
				eq = r[dst] == conv
			}
			if eq {
				r[dst] = 1
			} else {
				r[dst] = 0
			}
		default:
		}
	}

	return r
}

// recursive with backtracking search
func (d *Day24) findLargestModelNumber(n, depth int, zStack []int) (int, bool) {
	if depth == MonadLength {
		// last z has to be zero to have a valid model number
		return n, zStack[len(zStack)-1] == 0
	}

	z := zStack[len(zStack)-1]
	for v := 9; v >= 1; v-- {
		// Compute next z
		// Functions found by reverse-engineer the algo and applied
		// simplifications when possible
		var nz int
		if DivZ[depth] {
			if v == Kx[depth]+utils.Mod(z, 26) {
				nz = z / 26
			} else {
				// we are being agressive in the search because we know that if we have a
				// sub-program that can reduce z value, we should ! Else it is
				// impossible to reach z=0 at the end
				//
				// If we needed to compute next z value instead, it would multiply it by 26 after
				// the div by 26, resulting in no reducing operation
				continue
			}
		} else {
			nz = v + Ky[depth] + 26*z
		}

		// push
		zStack = append(zStack, nz)

		nn := n + v*utils.PowInt(10, MonadLength-depth-1)
		if sol, ok := d.findLargestModelNumber(nn, depth+1, zStack); ok {
			return sol, true
		}
		// pop
		zStack = zStack[0 : len(zStack)-1]
	}

	return 0, false
}

// recursive with backtracking search
func (d *Day24) findLowestModelNumber(n, depth int, zStack []int) (int, bool) {
	if depth == MonadLength {
		// last z has to be zero to have a valid model number
		return n, zStack[len(zStack)-1] == 0
	}

	z := zStack[len(zStack)-1]
	for v := 1; v <= 9; v++ {
		// Compute next z
		// Functions found by reverse-engineer the algo and applied
		// simplifications when possible
		var nz int
		if DivZ[depth] {
			if v == Kx[depth]+utils.Mod(z, 26) {
				nz = z / 26
			} else {
				// we are being agressive in the search because we know that if we have a
				// sub-program that can reduce z value, we should ! Else it is
				// impossible to reach z=0 at the end
				//
				// If we needed to compute next z value instead, it would multiply it by 26 after
				// the div by 26, resulting in no reducing operation
				continue
			}
		} else {
			nz = v + Ky[depth] + 26*z
		}

		// push
		zStack = append(zStack, nz)

		nn := n + v*utils.PowInt(10, MonadLength-depth-1)
		if sol, ok := d.findLowestModelNumber(nn, depth+1, zStack); ok {
			return sol, true
		}
		// pop
		zStack = zStack[0 : len(zStack)-1]
	}

	return 0, false
}
