package main

import (
	"encoding/json"
	"io"
	"strings"
)

type Pairs struct {
	Packet1 any `json:"array1"`
	Packet2 any  `json:"array2"`
}

func ParseInput(r io.ReaderAt) (input []Pairs, err error) {
	// Read file
	var data string
    buf := make([]byte, 1024)
	endFile := false
	offset := 0
	for !endFile {
		n, err := r.ReadAt(buf, int64(offset))
		if err == io.EOF {
			endFile = true
		}

		data += string(buf[:n])
		offset += n
	}

	// Parse file
	pairs := strings.Split(data, "\n\n")
	for _, value := range pairs {
		var Array1 any
		var Array2 any
		elt := strings.Split(value, "\n")

		err = json.Unmarshal([]byte(elt[0]), &Array1)
		if err != nil {
			return
		}
		err = json.Unmarshal([]byte(elt[1]), &Array2)
		if err != nil {
			return
		}

		input = append(input, Pairs{Packet1: Array1, Packet2: Array2})
	}

	return
}

func SumIndexes(indexes []int) int {
    var sum int
    
    for _, index := range indexes {
        sum += index
    }
    return sum
}

// <0 -> Right order
// > 0 -> Wring order
// = 0 -> Same packets while = 0, whether we continue
// and find at some point a result != 0, whether they are
// the same packet
func IsValidPacket(packetLeft any, packetRight any) int {
    left, ok_left := packetLeft.([]any)
    right, ok_right := packetRight.([]any)

    switch {
    case !ok_left && !ok_right:
        // Float at both
        return int(packetLeft.(float64) - packetRight.(float64))
    case !ok_left:
        // Type assertion failed -> float at left
        left = []any{packetLeft}
    case !ok_right:
        // Type assertion failed -> float at right
        right = []any{packetRight}
    }
    
    for i:=0; i < len(left) && i < len(right);i++ {
        // If result is < 0 -> Right Order, we can stop searching
        // If result is > 0 -> Wrong order, we can stop searching
        if result := IsValidPacket(left[i], right[i]); result != 0 {
            return result
        }
    }
    return len(left) - len(right)
}

func Part1(pairs []Pairs) (sum int, err error) {
    var indexes []int

    for i, pair := range pairs{
        packetLeft := pair.Packet1
        packetRight := pair.Packet2
        if IsValidPacket(packetLeft, packetRight) < 0 {
            indexes = append(indexes, i + 1)   
        }

    }

    sum = SumIndexes(indexes)
    return
}

func Part2(pairs []Pairs) {

}
