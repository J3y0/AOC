package days

import "testing"

func day2_ExampleData() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}

func TestDay2_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day2_ExampleData(), 150},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day2{}
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

func TestDay2_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day2_ExampleData(), 900},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day2{}
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
