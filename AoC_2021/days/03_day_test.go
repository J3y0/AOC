package days

import "testing"

func day3_ExampleData() string {
	return `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
}

func TestDay3_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day3_ExampleData(), 198},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day3{}
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

func TestDay3_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day3_ExampleData(), 230},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day3{}
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
