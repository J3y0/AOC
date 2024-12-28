package utils

import "strings"

func Reverse(str string) string {
	var res strings.Builder
	for i := len(str) - 1; i >= 0; i -= 1 {
		res.WriteByte(str[i])
	}

	return res.String()
}
