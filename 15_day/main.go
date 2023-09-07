package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"day15/sensor"
)

func main() {
	file, err := os.Open("./data/day15.txt")
	if err != nil {
		fmt.Printf("Error while opening file\n")
	}

	part1, err := solvePart1(file)
	if err != nil {
		fmt.Printf("Error while solving part 1\n")
	}

	fmt.Printf("Solution part1: %d\n", part1)
}

func solvePart1(file *os.File) (result int, err error) {
	var sensors []sensor.Sensor
	sensors, err = parseInput(file)
	if err != nil {
		return
	}

	var yCoord int = 2000000
	var lines []sensor.Line

	for _, s := range sensors {
		line := s.Coverage(yCoord)
		if line.Start != line.End {
			lines = append(lines, line)
		}
	}

	result += sensor.ComputeLength(lines)

	return
}

func parseInput(r io.ReaderAt) ([]sensor.Sensor, error) {
	// Read file
	var input string
	buf := make([]byte, 1024)
	endFile := false
	offset := 0
	for !endFile {
		n, errFile := r.ReadAt(buf, int64(offset))
		if errFile == io.EOF {
			endFile = true
		} else if errFile != nil {
			return nil, errFile
		}

		input += string(buf[:n])
		offset += n
	}

	// Parse coords of sensors and beacons
	var sensors []sensor.Sensor
	coordsRegex := regexp.MustCompile(`-?\d+`)
	coords := coordsRegex.FindAllString(input, -1)

	for i := 0; i < len(coords); i += 4 {
		sensorX, err := strconv.Atoi(coords[i])
		if err != nil {
			return nil, err
		}
		sensorY, err := strconv.Atoi(coords[i+1])
		if err != nil {
			return nil, err
		}
		sensorPos := sensor.Point{X: sensorX, Y: sensorY}

		beaconX, err := strconv.Atoi(coords[i+2])
		if err != nil {
			return nil, err
		}
		beaconY, err := strconv.Atoi(coords[i+3])
		if err != nil {
			return nil, err
		}
		beaconPos := sensor.Point{X: beaconX, Y: beaconY}

		sensors = append(sensors, sensor.Sensor{
			Position:              sensorPos,
			BeaconPosition:        beaconPos,
			ClosestBeaconDistance: sensor.Manhattan(sensorPos, beaconPos),
		})
	}

	return sensors, nil
}
