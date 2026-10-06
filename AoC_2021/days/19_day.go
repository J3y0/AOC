package days

import (
	"fmt"
	"main/utils"
	"strconv"
	"strings"
)

type scannerPoint struct {
	x, y, z int
}

func (p scannerPoint) add(q scannerPoint) scannerPoint {
	return scannerPoint{p.x + q.x, p.y + q.y, p.z + q.z}
}

func (p scannerPoint) sub(q scannerPoint) scannerPoint {
	return scannerPoint{p.x - q.x, p.y - q.y, p.z - q.z}
}

// Each orientation preserves handedness: reflections are not scanner rotations.
func (p scannerPoint) rotations() [24]scannerPoint {
	x, y, z := p.x, p.y, p.z
	return [24]scannerPoint{
		{x, y, z}, {x, -y, -z}, {x, z, -y}, {x, -z, y},
		{-x, y, -z}, {-x, -y, z}, {-x, z, y}, {-x, -z, -y},
		{y, x, -z}, {y, -x, z}, {y, z, x}, {y, -z, -x},
		{-y, x, z}, {-y, -x, -z}, {-y, z, -x}, {-y, -z, x},
		{z, x, y}, {z, -x, -y}, {z, y, -x}, {z, -y, x},
		{-z, x, -y}, {-z, -x, y}, {-z, y, x}, {-z, -y, -x},
	}
}

type Day19 struct {
	scanners  [][]scannerPoint
	beacons   map[scannerPoint]struct{}
	positions []scannerPoint
}

func (d *Day19) Parse(input string) error {
	var scanners [][]scannerPoint
	var seen map[scannerPoint]struct{}

	lines := utils.ParseLines(input)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "---") {
			id, ok := strings.CutPrefix(line, "--- scanner ")
			id, hasSuffix := strings.CutSuffix(id, " ---")
			n, err := strconv.Atoi(id)
			if !ok || !hasSuffix || err != nil || n < 0 {
				return fmt.Errorf("invalid scanner header %q", line)
			}
			if len(scanners) > 0 && len(scanners[len(scanners)-1]) == 0 {
				return fmt.Errorf("scanner has no beacons")
			}
			scanners = append(scanners, []scannerPoint{})
			seen = make(map[scannerPoint]struct{})
			continue
		}

		if len(scanners) == 0 {
			return fmt.Errorf("beacon without scanner header")
		}

		coords, err := utils.ParseLineToIntArray(line, ",")
		if err != nil {
			return fmt.Errorf("invalid beacon %q: %w", line, err)
		}

		if len(coords) != 3 {
			return fmt.Errorf("invalid beacon %q: expected three coordinates", line)
		}

		point := scannerPoint{coords[0], coords[1], coords[2]}
		if _, exists := seen[point]; exists {
			return fmt.Errorf("duplicate beacon %q", line)
		}

		seen[point] = struct{}{}
		idx := len(scanners) - 1
		scanners[idx] = append(scanners[idx], point)
	}

	if len(scanners) == 0 || len(scanners[len(scanners)-1]) == 0 {
		return fmt.Errorf("no scanner beacons")
	}

	d.scanners = scanners
	return nil
}

func (d *Day19) Part1() (int, error) {
	if err := d.buildMap(); err != nil {
		return 0, err
	}
	return len(d.beacons), nil
}

func (d *Day19) Part2() (int, error) {
	if err := d.buildMap(); err != nil {
		return 0, err
	}
	farthest := 0
	for i, p := range d.positions {
		for _, q := range d.positions[i+1:] {
			// manhattan distance
			delta := p.sub(q)
			distance := utils.AbsInt(delta.x) + utils.AbsInt(delta.y) + utils.AbsInt(delta.z)

			farthest = max(farthest, distance)
		}
	}
	return farthest, nil
}

func (d *Day19) buildMap() error {
	if d.beacons != nil {
		return nil
	}
	if len(d.scanners) == 0 {
		return fmt.Errorf("no scanners parsed")
	}

	// Scanner 0 defines the global coordinate system. Newly aligned scanners
	// become references so scanners connected through several overlaps can be found.
	aligned := make([][]scannerPoint, len(d.scanners))
	aligned[0] = d.scanners[0]
	positions := make([]scannerPoint, len(d.scanners))
	queue := []int{0}
	for head := 0; head < len(queue); head++ {
		reference := aligned[queue[head]]
		for i, scanner := range d.scanners {
			if aligned[i] != nil {
				continue
			}
			points, position, ok := alignScanner(reference, scanner)
			if ok {
				aligned[i] = points
				positions[i] = position
				queue = append(queue, i)
			}
		}
	}
	if len(queue) != len(d.scanners) {
		return fmt.Errorf("could not align %d scanners", len(d.scanners)-len(queue))
	}

	beacons := make(map[scannerPoint]struct{})
	for _, scanner := range aligned {
		for _, point := range scanner {
			beacons[point] = struct{}{}
		}
	}
	d.beacons = beacons
	d.positions = positions
	return nil
}

func alignScanner(reference, scanner []scannerPoint) ([]scannerPoint, scannerPoint, bool) {
	const requiredOverlap = 12
	if len(reference) < requiredOverlap || len(scanner) < requiredOverlap {
		return nil, scannerPoint{}, false
	}
	rotations := make([][24]scannerPoint, len(scanner))
	for i, point := range scanner {
		rotations[i] = point.rotations()
	}
	for orientation := range 24 {
		offsets := make(map[scannerPoint]int)
		for _, known := range reference {
			for _, rotated := range rotations {
				offset := known.sub(rotated[orientation])
				offsets[offset]++
				// Twelve pairs with the same offset identify twelve shared beacons.
				if offsets[offset] == requiredOverlap {
					aligned := make([]scannerPoint, len(scanner))
					for i, point := range rotations {
						aligned[i] = point[orientation].add(offset)
					}
					return aligned, offset, true
				}
			}
		}
	}
	return nil, scannerPoint{}, false
}
