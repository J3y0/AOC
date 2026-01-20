package days

import "testing"

func day6_ExampleData() string {
	return `3,4,3,1,2`
}

func TestDay6_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day6_ExampleData(), 5934},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day6{}
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

func TestDay6_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day6_ExampleData(), 26984457539},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day6{}
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
