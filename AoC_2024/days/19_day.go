package days

import (
	"os"
	"strings"
)

type Day19 struct {
	patterns []string
	designs  []string
}

func (d *Day19) Part1() (int, error) {
	// Build regex: ^(?:pattern1|pattern2|...|patternN)+$
	// r := strings.Builder{}
	// r.WriteString("^(?:")
	// for i, pat := range d.patterns {
	// 	if i == 0 {
	// 		r.WriteString(pat)
	// 		continue
	// 	}
	// 	r.WriteString("|" + pat)
	// }
	// r.WriteString(")+$")
	//
	// regxp := regexp.MustCompile(r.String())
	// count := 0
	// for _, design := range d.designs {
	// 	if regxp.MatchString(design) {
	// 		count++
	// 	}
	// }
	tot := 0
	for _, des := range d.designs {
		if d.countWays(des) > 0 {
			tot++
		}
	}
	return tot, nil
}

func (d *Day19) Part2() (int, error) {
	tot := 0
	for _, des := range d.designs {
		tot += d.countWays(des)
	}
	return tot, nil
}

func (d *Day19) Parse() error {
	content, err := os.ReadFile("./data/19_day.txt")
	if err != nil {
		return err
	}

	splitted := strings.Split(strings.TrimSpace(string(content)), "\n\n")
	d.patterns = strings.Split(splitted[0], ", ")
	d.designs = strings.Split(splitted[1], "\n")
	return nil
}

func (d *Day19) countWays(design string) int {
	n := len(design)
	dp := make([]int, n+1)
	dp[0] = 1 // Base case: 1 way to make empty string
	for i := 1; i <= n; i++ {
		for _, pat := range d.patterns {
			patternLen := len(pat)
			if patternLen <= i && design[i-patternLen:i] == pat {
				dp[i] += dp[i-patternLen]
			}
		}
	}

	return dp[n]
}
