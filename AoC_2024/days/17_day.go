package days

import (
	"aoc/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day17 struct {
	registers    map[rune]int
	instructions []int
	stdout       []string
}

func (d *Day17) Part1() (int, error) {
	var i int
	for d.registers['I'] < len(d.instructions) {
		i = d.registers['I']
		d.registers['I'] = d.opcode(d.instructions[i], d.instructions[i+1])
	}
	fmt.Println("Part1 answer (not a number):", strings.Join(d.stdout, ","))
	return 0, nil
}

func (d *Day17) Part2() (int, error) {
	return int(d.backtrack(d.instructions, 0)), nil
}

func (d *Day17) Parse() error {
	content, err := os.ReadFile("./data/17_day.txt")
	if err != nil {
		return err
	}

	split := strings.Split(strings.TrimSpace(string(content)), "\n\n")
	instructions := strings.Split(split[1], ": ")[1]
	d.instructions, err = utils.FromLine(instructions, ",")
	if err != nil {
		return err
	}

	var (
		reg rune
		val int
	)
	d.registers = make(map[rune]int)
	for _, l := range strings.Split(split[0], "\n") {
		_, err := fmt.Sscanf(l, "Register %c: %d", &reg, &val)
		if err != nil {
			return err
		}
		d.registers[reg] = val
	}
	// Add Instr Pointer
	d.registers['I'] = 0
	return nil
}

// Return new IP value after opcode
func (d *Day17) opcode(opcode, operand int) int {
	switch opcode {
	case 0:
		// adv instruction - combo oprd
		res := float64(d.registers['A']) / math.Pow(2, float64(d.comboOprd(operand)))
		d.registers['A'] = int(math.Trunc(res))
	case 1:
		// bxl instruction - literal oprd
		d.registers['B'] ^= operand
	case 2:
		// bst instruction - combo oprd
		d.registers['B'] = utils.PosMod(d.comboOprd(operand), 8)
	case 3:
		// jnz instruction - literal oprd
		if d.registers['A'] != 0 {
			return operand
		}
	case 4:
		// bxc instruction - operand ignored
		d.registers['B'] ^= d.registers['C']
	case 5:
		d.stdout = append(d.stdout, strconv.Itoa(utils.PosMod(d.comboOprd(operand), 8)))
	case 6:
		// bdv instruction - combo oprd
		res := float64(d.registers['A']) / math.Pow(2, float64(d.comboOprd(operand)))
		d.registers['B'] = int(math.Trunc(res))
	case 7:
		// cdv instruction - combo oprd
		res := float64(d.registers['A']) / math.Pow(2, float64(d.comboOprd(operand)))
		d.registers['C'] = int(math.Trunc(res))
	default:
		panic("unknown opcode")
	}
	return d.registers['I'] + 2
}

func (d *Day17) comboOprd(operand int) int {
	if operand >= 0 && operand <= 3 {
		return operand
	}

	var toReturn int
	switch {
	case 0 <= operand && operand <= 3:
		toReturn = operand
	case operand == 4:
		toReturn = d.registers['A']
	case operand == 5:
		toReturn = d.registers['B']
	case operand == 6:
		toReturn = d.registers['C']
	default:
		toReturn = 0
	}

	return toReturn
}

func (d *Day17) backtrack(program []int, res int64) int64 {
	if len(program) == 0 {
		return res
	}
	for val := range 8 {
		nextA := res<<3 | int64(val)
		b := utils.PosMod(int(nextA), 8)
		b = b ^ 3
		c := int(nextA) >> b
		b = b ^ 4
		b = b ^ c
		if utils.PosMod(b, 8) == program[len(program)-1] {
			sub := d.backtrack(program[:len(program)-1], nextA)
			if sub == -1 {
				continue
			}
			return sub
		}
	}
	return -1
}
