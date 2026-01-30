package days

import (
	"testing"
)

func TestDay16_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"small example data", "8A004A801A8002F478", 16},
		{"large example data", "A0016C880162017C3686B18A3D4780", 31},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day16{}
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

func TestDay16_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"simple sum", "C200B40A82", 3},
		{"simple product", "04005AC33890", 54},
		{"simple minimum", "880086C3E88112", 7},
		{"simple maximum", "CE00C43D881120", 9},
		{"simple less than", "D8005AC2A8F0", 1},
		{"simple greater", "F600BC2D8F", 0},
		{"simple equal", "9C005AC2F8F0", 0},
		{"complex equal", "9C0141080250320F1802104A08", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := Day16{}
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
