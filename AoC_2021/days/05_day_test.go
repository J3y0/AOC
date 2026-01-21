package days

import "testing"

func day5_ExampleData() string {
	return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
}

func TestDay5_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day5_ExampleData(), 5},
		{"basic horizontal overlap", "1,2 -> 1,5\n1,4 -> 1,8", 2},
		{"basic horizontal overlap exchanged", "1,4 -> 1,8\n1,2 -> 1,5", 2},
		{"basic vertical overlap", "2,1 -> 5,1\n8,1 -> 4,1", 2},
		{"basic vertical overlap exchanged", "4,1 -> 8,1\n2,1 -> 5,1", 2},
		{"basic one contained in another", "4,1 -> 6,1\n2,1 -> 8,1", 3},
		{"basic one contained in another exchanged", "2,1 -> 8,1\n4,1 -> 6,1", 3},
		{"basic no overlap", "2,4 -> 3,4\n4,1 -> 6,1", 0},
		{"basic intersect", "2,19 -> 10,19\n10,1 -> 10,20", 1},
		{"basic no intersect", "2,19 -> 10,19\n10,1 -> 10,10", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day5{}
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

func TestDay5_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example data", day5_ExampleData(), 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day5{}
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
