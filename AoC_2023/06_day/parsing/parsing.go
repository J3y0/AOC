package parsing

import (
	"regexp"
	"strconv"
	"strings"
)

type TimeDistance struct {
	Time     int
	Distance int
}

func ParsePart1(data []byte) (timeDist []TimeDistance, err error) {
	lines := strings.Split(string(data), "\n")

	regex := regexp.MustCompile(`\d+`)
	times := regex.FindAllString(lines[0], -1)
	distances := regex.FindAllString(lines[1], -1)

	var time int
	var distance int
	for i := 0; i < len(times); i++ {
		time, err = strconv.Atoi(times[i])
		if err != nil {
			return
		}
		distance, err = strconv.Atoi(distances[i])
		if err != nil {
			return
		}

		timeDist = append(timeDist, TimeDistance{Time: time, Distance: distance})
	}

	return
}

func ParsePart2(data []byte) (timeDist []TimeDistance, err error) {
	noSpace := strings.ReplaceAll(string(data), " ", "")
	lines := strings.Split(noSpace, "\n")

	regex := regexp.MustCompile(`\d+`)
	timeStr := regex.FindString(lines[0])
	distanceStr := regex.FindString(lines[1])

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return
	}
	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		return
	}

	timeDist = append(timeDist, TimeDistance{Time: time, Distance: distance})
	return
}
