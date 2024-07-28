package utils

func SumArray(arr []int) int {
	sum := 0
	for _, elt := range arr {
		sum += elt
	}

	return sum
}
