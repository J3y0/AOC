package main

import (
	"encoding/json"
	"io"
	"strings"
    "sort"
)

type Pairs struct {
	Packet1 any `json:"array1"`
	Packet2 any `json:"array2"`
}

type IndexError struct {}

func (ie *IndexError) Error() string {
    return "Error while finding index"
}

func ParseInput(r io.ReaderAt) (input []Pairs, err error) {
	// Read file
	var data string
	buf := make([]byte, 1024)
	endFile := false
	offset := 0
	for !endFile {
		n, errFile := r.ReadAt(buf, int64(offset))
		if errFile == io.EOF {
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

func sumIndexes(indexes []int) int {
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
func ValidPairs(packetLeft any, packetRight any) int {
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

	for i := 0; i < len(left) && i < len(right); i++ {
		// If result is < 0 -> Right Order, we can stop searching
		// If result is > 0 -> Wrong order, we can stop searching
		if result := ValidPairs(left[i], right[i]); result != 0 {
			return result
		}
	}
	return len(left) - len(right)
}

func Part1(pairs []Pairs) (sum int, err error) {
	var indexes []int

	for i, pair := range pairs {
		packetLeft := pair.Packet1
		packetRight := pair.Packet2
		if ValidPairs(packetLeft, packetRight) < 0 {
			indexes = append(indexes, i+1)
		}

	}

	sum = sumIndexes(indexes)
	return
}

func findDividerIndex(packets []any, elt string) (index int, err error) {
    var packetString []byte
    for i, packet := range packets {
        packetString, err = json.Marshal(packet)
        if err != nil {
            return
        }

        if string(packetString) == elt {
            index = i + 1
            return
        }
    }
    return
}

func Part2(pairs []Pairs) (decoderKey int, err error) {
    dividerPackets := [2]string{`[[2]]`, `[[6]]`}
    dividerPacketsIndex := [2]int{0, 0}
    
    var packets []any
    var temp any
    for i := 0;i < 2; i++ {
        err = json.Unmarshal([]byte(dividerPackets[i]), &temp)
        if err != nil {
            return
        }
        packets = append(packets, temp)
    }

    for _, pair := range pairs {
        packets = append(packets, pair.Packet1)
        packets = append(packets, pair.Packet2)
    }

    // Sort the array with all packets
    sort.SliceStable(packets, func(i, j int) bool {
        return ValidPairs(packets[i], packets[j]) < 0
    })

    var err_index error
    dividerPacketsIndex[0], err_index = findDividerIndex(packets, dividerPackets[0])
    if err_index != nil {
        err = &IndexError{}
        return
    }
    dividerPacketsIndex[1], err_index = findDividerIndex(packets, dividerPackets[1])
    if err_index != nil {
        err = &IndexError{}
        return
    }
    decoderKey = dividerPacketsIndex[0] * dividerPacketsIndex[1]
    return
}
