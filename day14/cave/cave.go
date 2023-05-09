package cave

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"day14/utils"
)

type Cave struct {
	MinX     int
	MinY     int
	MaxX     int
	MaxY     int
	DropSand utils.Coords
	CaveMap  [][]string
}

func (c *Cave) Pretty() {
	for _, line := range c.CaveMap {
		toPrint := ""
		for _, elt := range line {
			if elt != "" {
				toPrint += elt
			} else {
				toPrint += "."
			}
		}
		fmt.Println(toPrint)
	}
	fmt.Println()
}

func (c *Cave) Fill(coords [][]utils.Coords) {
	for _, line := range coords {
		for j := 1; j < len(line); j++ {
			co1 := line[j-1]
			co2 := line[j]

			for iter_x := utils.Min(co1.X, co2.X); iter_x <= utils.Max(co1.X, co2.X); iter_x++ {
				// Append on all lines before adding one element ?
				c.CaveMap[iter_x][co1.Y-c.MinY] = "#"
			}
			for iter_y := utils.Min(co1.Y, co2.Y); iter_y <= utils.Max(co1.Y, co2.Y); iter_y++ {
				c.CaveMap[co1.X][iter_y-c.MinY] = "#"
			}
		}
	}
}

func ParseInput(r io.ReaderAt) (*Cave, error) {
	// Read file
	var input string
	buf := make([]byte, 1024)
	endFile := false
	offset := 0
	for !endFile {
		n, errFile := r.ReadAt(buf, int64(offset))
		if errFile == io.EOF {
			endFile = true
		}

		input += string(buf[:n])
		offset += n
	}

	// Parse coords
	find_coordinates := `\d+,\d+`
	regCoord := regexp.MustCompile(find_coordinates)
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	allCoords := make([][]utils.Coords, 0, len(lines))
	var minX int = 1000
	var minY int = 1000
	var maxX int = 0
	var maxY int = 0
	for _, line := range lines {
		coords := regCoord.FindAllString(line, -1)

		var lineCoords []utils.Coords
		for i := range coords {
			co := strings.Split(coords[i], ",")
			y, err := strconv.Atoi(co[0])
			if err != nil {
				return &Cave{}, err
			}
			x, err := strconv.Atoi(co[1])
			if err != nil {
				return &Cave{}, err
			}
			lineCoords = append(lineCoords, utils.Coords{X: x, Y: y})
		}
		minX, minY, maxX, maxY = utils.FindMinAndMax(lineCoords, minX, minY, maxX, maxY)
		allCoords = append(allCoords, lineCoords)
	}

	c := &Cave{
		MinX:     minX,
		MinY:     minY,
		MaxX:     maxX,
		MaxY:     maxY,
		DropSand: utils.Coords{X: 0, Y: 500},
		CaveMap:  make([][]string, maxX+1, maxX+1),
	}

	// Init cave with right dimensions now
	for i := range c.CaveMap {
		initLine := make([]string, maxY-minY+1)
		c.CaveMap[i] = initLine
	}
	c.CaveMap[c.DropSand.X][c.DropSand.Y-c.MinY] = "+"

	// Fill the cave with rocks
	c.Fill(allCoords)

	return c, nil
}
