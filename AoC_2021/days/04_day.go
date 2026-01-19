package days

import (
	"fmt"
	"main/utils"
	"strings"
)

const BINGOSIZE = 5

type BingoGrid struct {
	grid [BINGOSIZE][BINGOSIZE]BingoTile
	won  bool
}

type BingoTile struct {
	number int
	drawn  bool
}

type Day4 struct {
	numbers   []int
	bingoList []BingoGrid
}

func (d *Day4) Parse(input string) error {
	var (
		numbers   []int
		bingoList []BingoGrid
	)

	input = strings.Trim(input, "\n")
	content := strings.Split(input, "\n\n")
	numbers, err := utils.ParseLineToIntArray(content[0], ",")
	if err != nil {
		return err
	}

	for i := 1; i < len(content); i++ {
		bingoGrid := BingoGrid{won: false}

		var gridNumbersLine [BINGOSIZE]int
		for j, line := range strings.Split(content[i], "\n") {
			_, err := fmt.Sscanf(line, "%d %d %d %d %d", &gridNumbersLine[0], &gridNumbersLine[1], &gridNumbersLine[2], &gridNumbersLine[3], &gridNumbersLine[4])
			if err != nil {
				return err
			}

			tilesLine := [BINGOSIZE]BingoTile{}
			for k, number := range gridNumbersLine {
				tilesLine[k].number = number
				tilesLine[k].drawn = false
			}

			bingoGrid.grid[j] = tilesLine
		}
		bingoList = append(bingoList, bingoGrid)
	}

	d.bingoList = bingoList
	d.numbers = numbers

	return nil
}

func (d *Day4) Part1() (int, error) {
	part1, numbersLeft := drawNumbersUntilFirst(d.bingoList, d.numbers)
	d.numbers = numbersLeft
	return part1, nil
}

func (d *Day4) Part2() (int, error) {
	return drawNumbersUntilLast(d.bingoList, d.numbers), nil
}

func drawNumbersUntilLast(bingoList []BingoGrid, numbers []int) int {
	var (
		winningNumber int
		lastGrid      BingoGrid
		totalGrid     = len(bingoList)
	)
	nbWinningGrids := 1
OuterLoop:
	for len(numbers) > 0 {
		number := numbers[0]
		numbers = numbers[1:]
		for k, bingo := range bingoList {
			if bingo.won {
				continue
			}
			for i := range BINGOSIZE {
				for j := range BINGOSIZE {
					if bingo.grid[i][j].number == number {
						bingo.grid[i][j].drawn = true
					}
				}
			}
			bingoList[k].grid = bingo.grid

			if !bingo.won && (hasColCompleted(bingo) || hasLineCompleted(bingo)) {
				bingoList[k].won = true
				nbWinningGrids++
				if nbWinningGrids == totalGrid {
					winningNumber = number
					lastGrid = bingo
					break OuterLoop
				}
			}
		}
	}
	return computeFinalScore(lastGrid, winningNumber)
}

func drawNumbersUntilFirst(bingoList []BingoGrid, numbers []int) (int, []int) {
	var (
		winnerGrid    BingoGrid
		winningNumber int
	)
OuterLoop:
	for len(numbers) > 0 {
		number := numbers[0]
		numbers = numbers[1:]
		for k, bingo := range bingoList {
			for i := range BINGOSIZE {
				for j := range BINGOSIZE {
					if bingo.grid[i][j].number == number {
						bingo.grid[i][j].drawn = true
					}
				}
			}
			bingoList[k].grid = bingo.grid

			if hasColCompleted(bingo) || hasLineCompleted(bingo) {
				bingoList[k].won = true
				winnerGrid = bingo
				winningNumber = number
				break OuterLoop
			}
		}
	}

	// Add again the number as we didn't go through all grids
	numbers = append([]int{winningNumber}, numbers...)

	return computeFinalScore(winnerGrid, winningNumber), numbers
}

func computeFinalScore(winnerGrid BingoGrid, winningNumber int) int {
	var total int
	for i := range BINGOSIZE {
		for j := range BINGOSIZE {
			if !winnerGrid.grid[i][j].drawn {
				total += winnerGrid.grid[i][j].number
			}
		}
	}

	return total * winningNumber
}

func hasColCompleted(bingoGrid BingoGrid) bool {
	for i := range BINGOSIZE {
		completed := true
		for j := range BINGOSIZE {
			if !bingoGrid.grid[j][i].drawn {
				completed = false
				break
			}
		}

		if completed {
			return true
		}
	}
	return false
}

func hasLineCompleted(bingoGrid BingoGrid) bool {
	for i := range BINGOSIZE {
		completed := true
		for j := range BINGOSIZE {
			if !bingoGrid.grid[i][j].drawn {
				completed = false
				break
			}
		}

		if completed {
			return true
		}
	}
	return false
}
