package parsing

import (
	"regexp"
	"strings"
)

type Junction struct {
	Left  string
	Right string
}

func NewJunction(left string, right string) Junction {
	return Junction{Left: left, Right: right}
}

func ParseFile(data []byte) (string, map[string]Junction) {
	lines := strings.Split(string(data), "\n")

	// Parse instructions in the first line
	instructions := lines[0]

	// Parse the rest
	regexJunction := regexp.MustCompile(`\w+, \w+`)

	mapJunction := make(map[string]Junction, 0)
	for i := 2; i < len(lines); i++ {
		splitted := strings.Split(lines[i], " = ")

		leftRight := strings.Split(regexJunction.FindString(splitted[1]), ", ")
		mapJunction[splitted[0]] = NewJunction(leftRight[0], leftRight[1])
	}

	return instructions, mapJunction
}

func FindStart(mapJunction map[string]Junction) []string {
	res := make([]string, 0)

	for key := range mapJunction {
		if strings.HasSuffix(key, "A") {
			res = append(res, key)
		}
	}

	return res
}
