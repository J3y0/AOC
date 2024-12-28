package utils

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func PosMod(x, n int) int {
	res := x % n
	if res < 0 {
		return (res + n) % n
	}

	return res
}
