package utils

type Coords struct {
	X int
	Y int
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func UpdateMinAndMax(coords []Coords, minY, maxX, maxY int) (int, int, int) {
	for i := range coords {
		x := coords[i].X
		y := coords[i].Y

		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	return minY, maxX, maxY
}
