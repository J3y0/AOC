package days

import (
	"testing"
)

func day12_SmallExampleData() string {
	return `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
}

func day12_MediumExampleData() string {
	return `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`
}

func day12_LargeExampleData() string {
	return `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`
}

func TestDay12_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"small example data", day12_SmallExampleData(), 10},
		{"medium example data", day12_MediumExampleData(), 19},
		{"large example data", day12_LargeExampleData(), 226},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day12{}
			if err := day.Parse(tt.input); err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			p1, err := day.Part1()
			if err != nil {
				t.Fatalf("Part1() error = %v", err)
			}

			if p1 != tt.want {
				t.Fatalf("got: %d, want: %d", p1, tt.want)
			}
		})
	}
}

func TestDay12_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"small example data", day12_SmallExampleData(), 36},
		{"medium example data", day12_MediumExampleData(), 103},
		{"large example data", day12_LargeExampleData(), 3509},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day12{}
			if err := day.Parse(tt.input); err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			p1, err := day.Part2()
			if err != nil {
				t.Fatalf("Part2() error = %v", err)
			}

			if p1 != tt.want {
				t.Fatalf("got: %d, want: %d", p1, tt.want)
			}
		})
	}
}
