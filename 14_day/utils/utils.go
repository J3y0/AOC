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

func FindMinAndMax(coords []Coords, minX, minY, maxX, maxY int) (int, int, int, int) {
	for i := range coords {
		x := coords[i].X
		y := coords[i].Y

		if x < minX {
			minX = x
		}
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
	return minX, minY, maxX, maxY
}
