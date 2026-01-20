package days

import "testing"

func day4_ExampleData() string {
	return `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`
}

func TestDay4_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day4_ExampleData(), 4512},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day4{}
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

func TestDay4_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day4_ExampleData(), 1924},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day4{}
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
