package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Pulse uint8

const (
	LOW Pulse = iota
	HIGH
	NOPULSE
)

type Modules map[string]Module

func (m Modules) Clone() Modules {
	modules := make(Modules)
	for k, v := range m {
		modules[k] = v
	}

	return modules
}

type Module struct {
	Name    string
	Type    rune
	Memory  map[string]Pulse
	State   bool
	Outputs []string
}

func (m *Module) Propagate(from string, pulse Pulse) Pulse {
	// Flip  flop module
	if m.Type == '%' {
		if pulse == HIGH {
			return NOPULSE
		} else {
			var old = m.State
			m.State = !m.State
			if old {
				return LOW
			}
			return HIGH
		}
	} else if m.Type == '&' {
		// Conjunction module
		if entry, ok := m.Memory[from]; ok {
			entry = pulse
			m.Memory[from] = entry
		}
		if AllHighPulses(m.Memory) {
			return LOW
		}
		return HIGH
	}

	// Else Broadcaster module
	return pulse
}

func AllHighPulses(memory map[string]Pulse) bool {
	for _, v := range memory {
		if v != HIGH {
			return false
		}
	}
	return true
}

type QueueItem struct {
	NameMod       string
	PulseReceived Pulse
	From          string
}

func PushButton(modules Modules) (int, int) {
	var (
		nbHigh   int
		nbLow    = 1
		curPulse = LOW // From button pushed
		queue    = make([]QueueItem, 0)
	)
	queue = append(queue, QueueItem{NameMod: "broadcaster", PulseReceived: curPulse})
	for len(queue) > 0 {
		// Dequeue
		item := queue[0]
		queue = queue[1:]

		nameMod := item.NameMod
		if _, ok := modules[nameMod]; !ok {
			continue
		}
		curModule := modules[nameMod]
		curPulse = item.PulseReceived

		nextPulse := curModule.Propagate(item.From, curPulse)
		modules[nameMod] = curModule // To take into account changes into memory or state of modules
		if nextPulse == NOPULSE {
			continue
		} else if nextPulse == LOW {
			nbLow += len(curModule.Outputs)
		} else {
			nbHigh += len(curModule.Outputs)
		}

		for _, out := range curModule.Outputs {
			queue = append(queue, QueueItem{NameMod: out, PulseReceived: nextPulse, From: nameMod})
		}
	}

	return nbHigh, nbLow
}

func Part1(modules Modules) int {
	var sumHigh, sumLow int
	for i := 0; i < 1000; i++ {
		high, low := PushButton(modules)
		sumHigh += high
		sumLow += low
	}

	return sumLow * sumHigh
}

func Part2(modules Modules) int {
	// Find name of only module connected to "rx"
	var feed string
	for name, mod := range modules {
		for _, out := range mod.Outputs {
			if out == "rx" {
				feed = name
			}
		}
	}

	// Init map containing cycle length
	cycleLength := make(map[string]int)
	for name, mod := range modules {
		for _, out := range mod.Outputs {
			if out == feed {
				cycleLength[name] = 0
			}
		}
	}

	var nbPushed int
mainLoop:
	for {
		nbPushed++
		var (
			curPulse = LOW // From button pushed
			queue    = make([]QueueItem, 0)
		)
		queue = append(queue, QueueItem{NameMod: "broadcaster", PulseReceived: curPulse})
		for len(queue) > 0 {
			// Dequeue
			item := queue[0]
			queue = queue[1:]

			nameMod := item.NameMod
			if _, ok := modules[nameMod]; !ok {
				continue
			}
			curModule := modules[nameMod]
			curPulse = item.PulseReceived

			if cycle, ok := cycleLength[nameMod]; ok && curPulse == LOW {
				if cycle == 0 {
					cycle = nbPushed
					cycleLength[nameMod] = cycle
				}
				if AllSeen(cycleLength) {
					break mainLoop
				}
			}

			nextPulse := curModule.Propagate(item.From, curPulse)
			modules[nameMod] = curModule // To take into account changes into memory or state of modules
			if nextPulse == NOPULSE {
				continue
			}
			for _, out := range curModule.Outputs {
				queue = append(queue, QueueItem{NameMod: out, PulseReceived: nextPulse, From: nameMod})
			}
		}
	}

	return LCM(cycleLength)
}

func AllSeen(cycleLength map[string]int) bool {
	for _, v := range cycleLength {
		if v == 0 {
			return false
		}
	}

	return true
}

func ReadModules(path string) (Modules, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
	modules := make(map[string]Module)
	for _, line := range lines {
		split := strings.Split(line, " -> ")
		outputs := strings.Split(split[1], ", ")
		var (
			name    string
			typeMod rune
		)
		if split[0] == "broadcaster" {
			name = split[0]
		} else {
			typeMod = rune(split[0][0])
			name = split[0][1:]
		}

		modules[name] = Module{
			Name:    name,
			Type:    typeMod,
			Memory:  nil,
			State:   false,
			Outputs: outputs,
		}
	}

	// Memory update
	for name, mod := range modules {
		for _, out := range mod.Outputs {
			if entry, ok := modules[out]; ok && entry.Type == '&' {
				if entry.Memory == nil {
					entry.Memory = make(map[string]Pulse)
				}
				entry.Memory[name] = LOW

				modules[out] = entry
			}
		}
	}

	return modules, nil
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(arr map[string]int) int {
	var lcm = 1
	for _, length := range arr {
		lcm = (lcm * length) / GCD(lcm, length)
	}
	return lcm
}

func main() {
	modules, err := ReadModules("./data/day20.txt")
	if err != nil {
		panic(err)
	}

	tStart := time.Now()
	// ------ Part1 ------
	part1 := Part1(modules.Clone())
	fmt.Println("Part 1: ", part1, time.Since(tStart))

	// ------ Part2 ------
	part2 := Part2(modules)
	fmt.Println("Part 2: ", part2, time.Since(tStart))
}
