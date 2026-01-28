package days

import (
	"main/utils"
	"slices"
	"strings"
)

type cave struct {
	name    string
	small   bool
	visited bool
}

type Day12 struct {
	caves map[string]cave
	edges map[string][]string
}

func (d *Day12) Parse(input string) error {
	d.caves = make(map[string]cave)
	d.edges = make(map[string][]string)
	lines := utils.ParseLines(input)
	for _, l := range lines {
		split := strings.SplitN(l, "-", 2)
		left, right := split[0], split[1]

		// Add cave info
		if _, ok := d.caves[left]; !ok {
			d.caves[left] = cave{name: left, small: utils.IsLower(left), visited: false}
		}

		if _, ok := d.caves[right]; !ok {
			d.caves[right] = cave{name: right, small: utils.IsLower(right), visited: false}
		}

		// Add edges (in both directions)
		if _, ok := d.edges[left]; !ok {
			d.edges[left] = make([]string, 0)
		}
		if !slices.Contains(d.edges[left], right) {
			d.edges[left] = append(d.edges[left], right)
		}
		if !slices.Contains(d.edges[right], left) {
			d.edges[right] = append(d.edges[right], left)
		}
	}

	return nil
}

func (d *Day12) Part1() (int, error) {
	return dfs(d.caves, d.edges, map[string]int{"start": 1}, "start", 1, false), nil
}

func (d *Day12) Part2() (int, error) {
	return dfs(d.caves, d.edges, map[string]int{"start": 1}, "start", 2, false), nil
}

func dfs(caves map[string]cave, edges map[string][]string, visited map[string]int, cur string, part int, twiceHit bool) int {
	if cur == "end" {
		return 1
	}

	count := 0
	for _, e := range edges[cur] {
		if !canVisit(e, caves, visited, part, twiceHit) {
			continue
		}
		if c, ok := caves[e]; ok && c.small {
			visited[e]++
			count += dfs(caves, edges, visited, e, part, twiceHit || visited[e] == 2)
			visited[e]--
		} else {
			count += dfs(caves, edges, visited, e, part, twiceHit)
		}
	}

	return count
}

func canVisit(name string, caves map[string]cave, visited map[string]int, part int, twiceHit bool) bool {
	if name == "start" {
		return visited[name] == 0
	}

	c, ok := caves[name]
	if !ok {
		return false
	}

	if !c.small {
		return true
	}

	if part == 1 {
		return visited[name] < 1
	}

	// part 2
	if !twiceHit && visited[name] < 2 {
		return true
	}

	if twiceHit && visited[name] < 1 {
		return true
	}

	return false
}
