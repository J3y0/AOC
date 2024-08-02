package utils

import "math"

func PowInt(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
