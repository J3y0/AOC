package parsing

import (
	"regexp"
	"strconv"
	"strings"

	utils "05_day/utils"
)

func ParseSeedRanges(seeds string) (parsedSeedRanges []*utils.Range, err error) {
	seedRegex := regexp.MustCompile(`\d+`)
	seedStrings := seedRegex.FindAllString(seeds, -1)
	for idx := 0; idx < len(seedStrings); idx += 2 {
		var seedStart int
		seedStart, err = strconv.Atoi(seedStrings[idx])
		if err != nil {
			return
		}

		var seedSize int
		seedSize, err = strconv.Atoi(seedStrings[idx+1])
		if err != nil {
			return
		}

		parsedSeedRanges = append(parsedSeedRanges, utils.RangeFromSize(seedStart, seedSize))
	}
	return
}

func ParseSeeds(seeds string) (parsedSeeds []int, err error) {
	seedRegex := regexp.MustCompile(`\d+`)
	seedStrings := seedRegex.FindAllString(seeds, -1)
	for _, elt := range seedStrings {
		var seedNb int
		seedNb, err = strconv.Atoi(elt)
		if err != nil {
			return
		}
		parsedSeeds = append(parsedSeeds, seedNb)
	}
	return
}

func ParseTransformations(allMaps []string) (map[string]*utils.TransformationRanges, error) {
	parsedMaps := make(map[string]*utils.TransformationRanges, 0)
	for _, individualMap := range allMaps {
		lines := strings.Split(individualMap, "\n")
		name := strings.Split(strings.Split(lines[0], " ")[0], "-")
		srcRangeName := name[0]
		destRangeName := name[2]

		var currentTransformationItem utils.TransformationItem
		currentTransformationItem.DestRangeName = destRangeName
		currentTransformationItem.SrcRangeName = srcRangeName

		arrayMapItems := make([]utils.TransformationItem, 0)
		for _, line := range lines[1:] {
			// lines are 3 numbers sep by a space ' '
			info := strings.Split(line, " ")
			destRangeNb, _ := strconv.Atoi(info[0])
			srcRangeNb, _ := strconv.Atoi(info[1])
			rangeNb, _ := strconv.Atoi(info[2])

			currentTransformationItem.Range.RangeStart = srcRangeNb
			currentTransformationItem.Range.RangeEnd = srcRangeNb + rangeNb
			currentTransformationItem.Shift = destRangeNb - srcRangeNb

			arrayMapItems = append(arrayMapItems, currentTransformationItem)
		}
		currentMap := &utils.TransformationRanges{Ranges: arrayMapItems}
		parsedMaps[srcRangeName] = currentMap
	}

	return parsedMaps, nil
}
