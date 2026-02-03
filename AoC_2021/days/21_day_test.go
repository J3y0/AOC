package days

import (
	"testing"
)

func day21_ExampleData() string {
	return `Player 1 starting position: 4
Player 2 starting position: 8`
}

func TestDay21_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day21_ExampleData(), 739785},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day21{}
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

func TestDay21_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day21_ExampleData(), 444356092776315},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day21{}
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
