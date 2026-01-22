package days

import (
	"testing"
)

func day9_ExampleData() string {
	return `2199943210
3987894921
9856789892
8767896789
9899965678`
}

func TestDay9_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day9_ExampleData(), 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day9{}
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

func TestDay9_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day9_ExampleData(), 1134},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day9{}
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
