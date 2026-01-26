package utils

// All neighbors' unitary directions
var Neighbors = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

var OrthogonalNeighbors = [4][2]int{
	{-1, 0},
	{0, -1}, {0, 1},
	{1, 0},
}
