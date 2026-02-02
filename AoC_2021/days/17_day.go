package days

import (
	"fmt"
	"main/utils"
	"strings"
)

type Range struct {
	start, end int
}

func (r *Range) Contains(v int) bool {
	return r.start <= v && v <= r.end
}

type Day17 struct {
	xr, yr Range
}

func (d *Day17) Parse(input string) error {
	input = strings.TrimSpace(input)
	var xr, yr Range
	_, err := fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &xr.start, &xr.end, &yr.start, &yr.end)
	d.xr = xr
	d.yr = yr
	return err
}

/*
 * To reach the maximum height value, you don't need to care for x coordinate. We can assume there is always one matching the y number of steps
 * (since target area is pretty wide).
 *
 * We can determine that:
 * 		y(t) = vy*t - sum{i=0, t-1}(i)
 * 		y(t) = vy*t - t(t - 1)/2
 * 		y(t) = -t^2/2 + (vy + 1/2)t
 * This function reach a maximum of vy(vy+1)/2 at t=vy
 *
 * Now, to determine the correct vy, we know that we reach y=0 at speed vy=-vy_start. We assume that we want to reach the lowest bound
 * of the target area in one step, meaning we want vy=-(lowest_bound + 1). +1 because the next step will add another -1 that will compensate.
 *
 * Finally, we note yb the lowest bound and get ymax for vy_start=|yb+1|
 */
func (d *Day17) Part1() (int, error) {
	vy_start := utils.AbsInt(d.yr.start + 1)
	return (vy_start * (vy_start + 1)) / 2, nil
}

func (d *Day17) Part2() (int, error) {
	yrange := max(utils.AbsInt(d.yr.start), utils.AbsInt(d.yr.end))
	count := 0
	for vx := 1; vx <= d.xr.end; vx++ {
		for vy := -yrange; vy <= yrange; vy++ {
			if valid(d.xr, d.yr, vx, vy, yrange) {
				count += 1
			}
		}
	}
	return count, nil
}

func valid(xr, yr Range, vx, vy, r int) bool {
	x := vx
	y := vy
	for i := 1; i <= 2*r; i++ {
		if inTarget(xr, yr, x, y) {
			return true
		}
		if vx-i >= 0 {
			x += vx - i
		}
		y += vy - i
	}
	return false
}

func inTarget(xr, yr Range, x, y int) bool {
	return xr.Contains(x) && yr.Contains(y)
}
