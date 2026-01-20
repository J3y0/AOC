package days

import "testing"

func day7_ExampleData() string {
	return `16,1,2,0,4,2,7,1,2,14`
}

func TestDay7_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day7_ExampleData(), 37},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day7{}
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

func TestDay7_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day7_ExampleData(), 168},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day7{}
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
