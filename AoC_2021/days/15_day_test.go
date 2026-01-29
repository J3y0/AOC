package days

import (
	"testing"
)

func day15_ExampleData() string {
	return `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`
}

func TestDay15_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day15_ExampleData(), 40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day15{}
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

func TestDay15_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day15_ExampleData(), 315},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day15{}
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
