package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const HEADER_SIZE = 6

type typeID = int8

const (
	SUM typeID = iota
	PRODUCT
	MINIMUM
	MAXIMUM
	LITERAL
	GREATER
	LESS
	EQUAL
)

type packetHeader struct {
	version int8
	typeID  int8
}

type packet struct {
	header    packetHeader
	isLiteral bool

	// literal packet field
	literalValue int64

	// operator packet fields
	lengthID   bool
	length     int // either nb of subpackets either total length
	subpackets []*packet
}

func (p *packet) compute() (int, error) {
	var result int
	switch p.header.typeID {
	case SUM:
		for _, sub := range p.subpackets {
			if sub.header.typeID == LITERAL {
				result += int(sub.literalValue)
				continue
			}

			v, err := sub.compute()
			if err != nil {
				return 0, nil
			}
			result += v
		}
	case PRODUCT:
		result = 1
		for _, sub := range p.subpackets {
			v, err := sub.compute()
			if err != nil {
				return 0, nil
			}
			result *= v
		}
	case MINIMUM:
		result = math.MaxInt
		for _, sub := range p.subpackets {
			v, err := sub.compute()
			if err != nil {
				return 0, nil
			}
			if v < result {
				result = v
			}
		}
	case MAXIMUM:
		for _, sub := range p.subpackets {
			v, err := sub.compute()
			if err != nil {
				return 0, nil
			}
			if v > result {
				result = v
			}
		}
	case LITERAL:
		return int(p.literalValue), nil
	case GREATER:
		// exactly 2 subpackets
		first, err := p.subpackets[0].compute()
		if err != nil {
			return 0, nil
		}
		second, err := p.subpackets[1].compute()
		if err != nil {
			return 0, nil
		}
		if first > second {
			result = 1
		}
	case LESS:
		// exactly 2 subpackets
		first, err := p.subpackets[0].compute()
		if err != nil {
			return 0, nil
		}
		second, err := p.subpackets[1].compute()
		if err != nil {
			return 0, nil
		}
		if first < second {
			result = 1
		}
	case EQUAL:
		// exactly 2 subpackets
		first, err := p.subpackets[0].compute()
		if err != nil {
			return 0, nil
		}
		second, err := p.subpackets[1].compute()
		if err != nil {
			return 0, nil
		}
		if first == second {
			result = 1
		}
	default:
		return 0, fmt.Errorf("unknown type ID detected: %d", p.header.typeID)
	}

	return result, nil
}

type Day16 struct {
	raw string
}

func (d *Day16) Parse(input string) error {
	input = strings.TrimSpace(input)
	var sb strings.Builder
	sb.Grow(4 * len(input))
	for _, r := range input {
		parsed, err := strconv.ParseUint(string(r), 16, 4)
		if err != nil {
			return err
		}
		bin := fmt.Sprintf("%04b", parsed)
		sb.WriteString(bin)
	}
	d.raw = sb.String()
	return nil
}

func (d *Day16) Part1() (int, error) {
	packet, _, err := parsePacket(d.raw)
	if err != nil {
		return 0, err
	}
	return sumVersion(packet), nil
}

func (d *Day16) Part2() (int, error) {
	packet, _, err := parsePacket(d.raw)
	if err != nil {
		return 0, err
	}
	return packet.compute()
}

func sumVersion(p *packet) int {
	sum := int(p.header.version)
	for _, sub := range p.subpackets {
		sum += sumVersion(sub)
	}
	return sum
}

func parsePacket(s string) (*packet, int, error) {
	version, err := strconv.ParseInt(s[:3], 2, 8)
	if err != nil {
		return nil, 0, err
	}
	typeID, err := strconv.ParseInt(s[3:6], 2, 8)
	if err != nil {
		return nil, 0, err
	}

	if int8(typeID) == LITERAL {
		val, offset, err := parseLiteralValue(s)
		if err != nil {
			return nil, 0, err
		}

		return &packet{
			header:       packetHeader{int8(version), int8(typeID)},
			isLiteral:    true,
			literalValue: val,
		}, offset, nil
	}

	// operator packet
	lengthID, err := strconv.ParseBool(string(s[HEADER_SIZE]))
	if err != nil {
		return nil, 0, err
	}

	if lengthID {
		// number of subpackets
		nbSubPackets, err := strconv.ParseInt(s[HEADER_SIZE+1:HEADER_SIZE+1+11], 2, 32)
		if err != nil {
			return nil, 0, err
		}

		sub := make([]*packet, nbSubPackets)
		offset := HEADER_SIZE + 12
		for i := range nbSubPackets {
			subPacket, nbParsed, err := parsePacket(s[offset:])
			if err != nil {
				return nil, 0, err
			}
			offset += nbParsed
			sub[i] = subPacket
		}

		return &packet{
			header:     packetHeader{int8(version), int8(typeID)},
			isLiteral:  false,
			lengthID:   lengthID,
			length:     int(nbSubPackets),
			subpackets: sub,
		}, offset, nil
	} else {
		// data total length
		totalLength, err := strconv.ParseInt(s[HEADER_SIZE+1:HEADER_SIZE+1+15], 2, 32)
		if err != nil {
			return nil, 0, err
		}

		sub := make([]*packet, 0)
		offset := HEADER_SIZE + 16
		for offset < int(totalLength)+HEADER_SIZE+16 {
			subPacket, nbParsed, err := parsePacket(s[offset:])
			if err != nil {
				return nil, 0, err
			}
			offset += nbParsed
			sub = append(sub, subPacket)
		}

		return &packet{
			header:     packetHeader{int8(version), int8(typeID)},
			isLiteral:  false,
			lengthID:   lengthID,
			length:     int(totalLength),
			subpackets: sub,
		}, offset, nil
	}
}

// nbParsed is the total of bytes parsed by the function
func parseLiteralValue(s string) (val int64, nbParsed int, err error) {
	var (
		sb    strings.Builder
		still = true
	)
	nbParsed = HEADER_SIZE
	for still {
		still, err = strconv.ParseBool(string(s[nbParsed]))
		if err != nil {
			return
		}
		nbParsed++
		sb.WriteString(s[nbParsed : nbParsed+4])
		nbParsed += 4
	}
	val, err = strconv.ParseInt(sb.String(), 2, 64)
	return
}
