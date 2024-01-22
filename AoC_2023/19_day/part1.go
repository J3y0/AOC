package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Part map[rune]int

func (p Part) Sum() int {
	return p['x'] + p['m'] + p['a'] + p['s']
}

type Workflow struct {
	Rules    []string
	Fallback string
}

type Workflows map[string]Workflow

func ParseInput(path string) (Workflows, []Part, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	// First element contains the workflows
	// Second element contains the parts
	content := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n\n")

	// Parse workflows
	workflows := make(Workflows)
	for _, line := range strings.Split(content[0], "\n") {
		temp := strings.Split(line, "{")
		rules := strings.Split(temp[1][:len(temp[1])-1], ",")
		workflow := Workflow{Rules: rules[:len(rules)-1], Fallback: rules[len(rules)-1]}
		workflows[temp[0]] = workflow
	}

	// Parse Parts
	partsLines := strings.Split(content[1], "\n")
	parts := make([]Part, len(partsLines))
	for i, partLine := range partsLines {
		var x, m, a, s int
		_, err := fmt.Sscanf(partLine, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
		if err != nil {
			return nil, nil, err
		}
		parts[i] = Part{'x': x, 'm': m, 'a': a, 's': s}
	}

	return workflows, parts, nil
}

func PassCondition(instr string, part Part) (bool, error) {
	value, err := strconv.Atoi(instr[2:])
	if err != nil {
		return false, err
	}

	switch instr[1] {
	case '<':
		if part[rune(instr[0])] < value {
			return true, nil
		}
	case '>':
		if part[rune(instr[0])] > value {
			return true, nil
		}
	}

	return false, nil
}

func PartIsAccepted(workflows Workflows, part Part) (bool, error) {
	curModuleName := "in"
ModuleLoop:
	for curModuleName != "A" {
		if curModuleName == "R" {
			return false, nil
		}

		for _, rules := range workflows[curModuleName].Rules {
			step := strings.Split(rules, ":")
			// Retrieve condition
			passCondition, err := PassCondition(step[0], part)
			if err != nil {
				return false, err
			}
			// Check if pass condition
			if passCondition {
				// Name of next module
				curModuleName = step[1]
				continue ModuleLoop
			}
		}

		curModuleName = workflows[curModuleName].Fallback
	}

	return true, nil
}

func Part1(parts []Part, workflows Workflows) (int, error) {
	var sumAccepted int
	for _, part := range parts {
		isAccepted, err := PartIsAccepted(workflows, part)
		if err != nil {
			return 0, err
		}
		if isAccepted {
			sumAccepted += part.Sum()
		}
	}

	return sumAccepted, nil
}
