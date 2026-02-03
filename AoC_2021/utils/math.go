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

// mod n would produce numbers between [[0; n-1]]. This function shifts this interval by 1 and produce numbers between [[1; n]]
func ShiftedMod(x, n int) int {
	return (x-1)%n + 1
}
