package utils

type Pos struct {
	X, Y int
}

func OutOfGrid(p Pos, maxRow, maxCol int) bool {
	return p.X < 0 || p.Y < 0 || p.X >= maxRow || p.Y >= maxCol
}
