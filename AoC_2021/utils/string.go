package utils

import "unicode"

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) || !unicode.IsLower(r) {
			return false
		}
	}

	return true
}
