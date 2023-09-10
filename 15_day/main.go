package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
    "time"

	"day15/sensor"
)

func main() {
	file, err := os.Open("./data/day15.txt")
	if err != nil {
		fmt.Printf("Error while opening file\n")
	}

    sensors, err := parseInput(file)
    if err != nil {
        fmt.Println("Error while parsing input file...")
    }
    
    tStart := time.Now()
	part1, err := solvePart1(sensors)
	if err != nil {
		fmt.Printf("Error while solving part 1\n")
	}
    tEnd := time.Since(tStart)

	fmt.Printf("Solution part1: %d\n", part1)
    fmt.Printf("It took: %s\n", tEnd)

    tStart = time.Now()
	part2, err := solvePart2(sensors)
	if err != nil {
		fmt.Printf("Error while solving part 2\n")
	}
    tEnd = time.Since(tStart)

	fmt.Printf("Solution part2: %d\n", part2)
    fmt.Printf("It took: %s\n", tEnd)
}

func solvePart1(sensors []sensor.Sensor) (result int, err error) {
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

func solvePart2(sensors []sensor.Sensor) (result int, err error) {
    var distress sensor.Point
    var allIntersections []sensor.Point
    // Find coord of the beacon emetting distress signal

    for _, s1 := range sensors {
        for _, s2 := range sensors {
            allIntersections = sensor.AddIntersectionPoints(s1, s2, allIntersections)
        }
    }
    
    for _, p := range allIntersections {
        valid := true
        for _, s := range sensors {
            if !valid || sensor.Manhattan(s.Position, p) <= s.ClosestBeaconDistance {
                valid = false
                break
            }
        }
        if valid {
            distress = p
            break
        }
    }

    fmt.Printf("Distress signal position: x=%d, y=%d\n", distress.X, distress.Y)
    // Compute the tuning frequency
    result = distress.X * 4000000 + distress.Y
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
