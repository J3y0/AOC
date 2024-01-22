package parsing

import "strings"

type Position struct {
	X      int
	Y      int
	Symbol string
}

func ParseMap(data []byte) (mapIsland [][]string, startPosition Position) {
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		newLine := make([]string, 0)

		for j := range line {
			newLine = append(newLine, string(line[j]))

			if string(line[j]) == "S" {
				startPosition.X = i
				startPosition.Y = j
				startPosition.Symbol = "S"
			}
		}
		mapIsland = append(mapIsland, newLine)
	}

	return
}
