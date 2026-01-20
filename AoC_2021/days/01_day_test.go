package days

import "testing"

func day1_ExampleData() string {
	return `199
200
208
210
200
207
240
269
260
263`
}

func TestDay1_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day1_ExampleData(), 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day1{}
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

func TestDay1_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day1_ExampleData(), 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day1{}
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
