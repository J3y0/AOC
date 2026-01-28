package utils

func SumArray(arr []int) int {
	sum := 0
	for _, elt := range arr {
		sum += elt
	}

	return sum
}

// remove duplicates from an array
func Duplicates[Slice ~[]E, E comparable](arr Slice) Slice {
	found := make(map[E]bool)
	res := make(Slice, 0)
	for _, elt := range arr {
		if found[elt] {
			continue
		}
		found[elt] = true
		res = append(res, elt)
	}

	return res
}
