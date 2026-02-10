package days

import (
	"fmt"
	"main/utils"
)

type cuboid struct {
	isOn       bool
	xr, yr, zr utils.Range
}

func (c *cuboid) volume() int {
	vol := (c.xr.End - c.xr.Start + 1) * (c.yr.End - c.yr.Start + 1) * (c.zr.End - c.zr.Start + 1)
	if c.isOn {
		return vol
	}
	return -vol
}

// c2 is considered to be the latest step applied. For instance if c1 is on and c2 is off, then the intersection
// would be considered off.
func intersect(c1, c2 cuboid) (cuboid, bool) {
	interXr, okX := utils.Intersect(c1.xr, c2.xr)
	interYr, okY := utils.Intersect(c1.yr, c2.yr)
	interZr, okZ := utils.Intersect(c1.zr, c2.zr)
	if !okX || !okY || !okZ {
		return cuboid{}, false
	}

	var interState bool
	if c1.isOn && c2.isOn {
		// remove once as counted twice
		interState = false
	} else if !c1.isOn && !c2.isOn {
		// add once as removed twice
		interState = true
	} else {
		// else the second cuboid state is final
		interState = c2.isOn
	}

	return cuboid{
		isOn: interState,
		xr:   interXr,
		yr:   interYr,
		zr:   interZr,
	}, true
}

type Day22 struct {
	cuboids []cuboid
}

func (d *Day22) Parse(input string) error {
	lines := utils.ParseLines(input)
	d.cuboids = make([]cuboid, len(lines))
	for i, l := range lines {
		var (
			onoff      string
			xr, yr, zr utils.Range
		)
		_, err := fmt.Sscanf(l, "%s x=%d..%d,y=%d..%d,z=%d..%d", &onoff, &xr.Start, &xr.End, &yr.Start, &yr.End, &zr.Start, &zr.End)
		if err != nil {
			return err
		}

		d.cuboids[i] = cuboid{
			isOn: onoff == "on",
			xr:   xr,
			yr:   yr,
			zr:   zr,
		}
	}
	return nil
}

func (d *Day22) Part1() (int, error) {
	cubesSet := make([]cuboid, 0)
	limitRange := utils.Range{Start: -50, End: 50}
	for _, cube := range d.cuboids {
		// Clip cube such that it does not exceed -50, 50 range for all coordinates
		initxr, okx := utils.Intersect(limitRange, cube.xr)
		inityr, oky := utils.Intersect(limitRange, cube.yr)
		initzr, okz := utils.Intersect(limitRange, cube.zr)
		if !okx || !oky || !okz {
			continue
		}
		initCube := cuboid{
			isOn: cube.isOn,
			xr:   initxr,
			yr:   inityr,
			zr:   initzr,
		}
		var toAdd []cuboid
		for _, cs := range cubesSet {
			if interCube, hasInter := intersect(cs, initCube); hasInter {
				toAdd = append(toAdd, interCube)
			}
		}
		if initCube.isOn {
			toAdd = append(toAdd, initCube)
		}
		cubesSet = append(cubesSet, toAdd...)
	}

	tot := 0
	for _, c := range cubesSet {
		tot += c.volume()
	}
	return tot, nil
}

/*
 * The reasoning is based on set theory in 2d and with the formula:
 *   (2 sets) AvB = A + B - A^B
 *   (3 sets) AvBvC = (A + B - A^B) + C - A^C - B^C + A^B^C
 *
 * While this works for computing the total cardinality, we only want the "on"
 * cardinality. This will only impact the signs in the operation.
 *
 * An "off" cuboid should be removed, hence the sign "-"
 * An "on" cuboid should be added, hence the sign "+"
 *
 * For a single cuboid, if it is "on", we add it. If it is "off", a question arises:
 * does it intersects with any other "on" cuboids ?
 *   If no, do not take it into account: it is isolated;
 *   If yes, only substract the intersection of both.
 *
 * Also, the resulting cuboid of the intersection has its state determined by the following table;
 *   on / on -> off (because you count cuboids twice, thus you need to remove the intersection once)
 *   off / off -> on (same reflexion, you remove cuboids twice so you need to add the intersection once)
 *   on / off -> off
 *   off / on -> on
 */
func (d *Day22) Part2() (int, error) {
	// list containing cuboid and intersections. It holds all the terms of
	// the set's formula AvB = A + B - A^B
	//
	// Terms representing single cuboid that are turned off are ignored (count for 0 cuboid turned on)
	//
	// For example, for 3 sets A, B, C (C is an off step): cubesSet = [A, B, AvB, AvC, AvBvC]
	cubesSet := make([]cuboid, 0)
	for _, cube := range d.cuboids {
		var toAdd []cuboid
		for _, cs := range cubesSet {
			if interCube, hasInter := intersect(cs, cube); hasInter {
				toAdd = append(toAdd, interCube)
			}
		}

		// Add the cuboid itself if and only if the cuboid is turned on. Else it accounts for nothing
		if cube.isOn {
			toAdd = append(toAdd, cube)
		}

		cubesSet = append(cubesSet, toAdd...)
	}

	tot := 0
	for _, c := range cubesSet {
		tot += c.volume()
	}
	return tot, nil
}
