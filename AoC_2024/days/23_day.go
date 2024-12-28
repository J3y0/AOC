package days

import (
	"fmt"
	"hash/fnv"
	"os"
	"slices"
	"sort"
	"strings"

	"aoc/utils"
)

type Day23 struct {
	vertices map[string][]string
}

func (d *Day23) Part1() (int, error) {
	tot := 0
	seen := make(map[uint64]bool)
	for k, verticesList := range d.vertices {
		for i := range verticesList {
			for j := i + 1; j < len(verticesList); j++ {
				hKey := d.hash(k, verticesList[i], verticesList[j])
				if seen[hKey] {
					continue
				}
				seen[hKey] = true
				if !strings.HasPrefix(k, "t") && !strings.HasPrefix(verticesList[i], "t") && !strings.HasPrefix(verticesList[j], "t") {
					continue
				}

				if slices.Contains(d.vertices[verticesList[j]], verticesList[i]) && slices.Contains(d.vertices[verticesList[i]], verticesList[j]) {
					tot++
				}
			}
		}
	}
	return tot, nil
}

func (d *Day23) Part2() (int, error) {
	allVertices := make([]string, len(d.vertices))
	i := 0
	for k := range d.vertices {
		allVertices[i] = k
		i++
	}

	cliques := d.bronKerbosch([]string{}, allVertices, []string{})

	// Find max cliques among all cliques
	maxLen := 0
	var maxClique []string
	for _, c := range cliques {
		if len(c) > maxLen {
			maxLen = len(c)
			maxClique = c
		}
	}

	sort.Strings(maxClique)
	fmt.Println("Part2 answer (not a number):", strings.Join(maxClique, ","))

	return maxLen, nil
}

func (d *Day23) Parse() error {
	content, err := os.ReadFile("./data/23_day.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	d.vertices = make(map[string][]string)
	for _, l := range lines {
		splitted := strings.Split(l, "-")

		// Bidirectionnal
		if _, ok := d.vertices[splitted[0]]; !ok {
			d.vertices[splitted[0]] = make([]string, 0)
		}
		d.vertices[splitted[0]] = append(d.vertices[splitted[0]], splitted[1])

		if _, ok := d.vertices[splitted[1]]; !ok {
			d.vertices[splitted[1]] = make([]string, 0)
		}
		d.vertices[splitted[1]] = append(d.vertices[splitted[1]], splitted[0])
	}
	return nil
}

func (d *Day23) hash(a, b, c string) uint64 {
	h := fnv.New64a()

	h.Write([]byte(a))
	ah := h.Sum64()
	h.Reset()

	h.Write([]byte(b))
	bh := h.Sum64()
	h.Reset()

	h.Write([]byte(c))
	ch := h.Sum64()
	return ah ^ bh ^ ch
}

func (d *Day23) bronKerbosch(r, p, x []string) [][]string {
	var cliques [][]string
	if len(p) == 0 && len(x) == 0 {
		clique := make([]string, len(r))
		copy(clique, r) // Add a copy of R as a new maximal clique
		cliques = append(cliques, clique)
		return cliques
	}

	for i := 0; i < len(p); i++ {
		rbis := append(r, p[i])
		pbis := d.buildIntersect(p, d.vertices[p[i]]) // intersect with neighbors
		xbis := d.buildIntersect(x, d.vertices[p[i]])

		subcliques := d.bronKerbosch(rbis, pbis, xbis)
		cliques = append(cliques, subcliques...)

		x = append(x, p[i])
		p = utils.OmitIndex(p, i)
		i-- // index had been omitted
	}

	return cliques
}

func (d *Day23) buildIntersect(list []string, neighbors []string) []string {
	res := make([]string, 0)
	for _, n := range neighbors {
		if slices.Contains(list, n) {
			res = append(res, n)
		}
	}

	return res
}
