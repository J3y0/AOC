package days

import (
	"testing"
)

func day17_ExampleData() string {
	return `target area: x=20..30, y=-10..-5`
}

func TestDay17_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day17_ExampleData(), 45},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day17{}
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

func TestDay17_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day17_ExampleData(), 112},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day17{}
			if err := day.Parse(tt.input); err != nil {
				t.Fatalf("Parse() error = %v", err)
			}

			p2, err := day.Part2()
			if err != nil {
				t.Fatalf("Part2() error = %v", err)
			}

			if p2 != tt.want {
				t.Fatalf("got: %d, want: %d", p2, tt.want)
			}
		})
	}
}
