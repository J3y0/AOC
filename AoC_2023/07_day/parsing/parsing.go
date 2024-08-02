package parsing

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Card value (for 10, Heads and As)
const (
	Ten = iota + 10
	Jack
	Queen
	King
	As
)

// Add Joker
const Joker int = 1

// Type of hands
const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Bid        int
	HandValues [5]int // Array containing cards' value of the hand
	Type       int
}

func NewHandFromStrAndBid(hand string, bid int, part int) Hand {
	cardValues := GetCardValues(hand, part)
	handType := GetHandType(cardValues)
	return Hand{
		Bid:        bid,
		Type:       handType,
		HandValues: cardValues,
	}
}

func ParseInput(data []byte, part int) ([]Hand, error) {
	var hands []Hand

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		temp := strings.Split(line, " ")
		bid, err := strconv.Atoi(temp[1])
		if err != nil {
			return nil, err
		}
		hands = append(hands, NewHandFromStrAndBid(temp[0], bid, part))
	}

	return hands, nil
}

func GetCardValues(hand string, part int) [5]int {
	cards := [5]int{0}

	valueRegex := regexp.MustCompile(`[2-9]`)
	for i, card := range hand {
		if valueRegex.Match([]byte{hand[i]}) {
			nb, _ := strconv.Atoi(string(card))
			cards[i] = nb
			continue
		}

		switch string(card) {
		case "T":
			cards[i] = Ten
		case "J":
			if part == 1 {
				cards[i] = Jack
			} else if part == 2 {
				cards[i] = Joker
			}
		case "Q":
			cards[i] = Queen
		case "K":
			cards[i] = King
		case "A":
			cards[i] = As
		default:
			fmt.Printf("[!] Char %c not recognized\n", card)
		}
	}
	return cards
}

func GetHandType(cards [5]int) int {
	occurenceCards := make(map[int]int, 0)

	var countJoker int
	for _, card := range cards {
		if card == Joker {
			countJoker++
			continue
		}
		occurenceCards[card] += 1
	}

	var length int = len(occurenceCards)

	// Case with full jokers as I don't add them to the map
	if length == 0 {
		return FiveOfAKind
	}

	var bestType int = HighCard
	for _, val := range occurenceCards {
		valWithJoker := val + countJoker
		if valWithJoker == 5 {
			bestType = FiveOfAKind
			continue
		}

		if valWithJoker == 4 && FourOfAKind > bestType {
			bestType = FourOfAKind
			continue
		}

		if valWithJoker == 3 && length == 2 && FullHouse > bestType {
			bestType = FullHouse
			continue
		}

		if valWithJoker == 3 && length == 3 && ThreeOfAKind > bestType {
			bestType = ThreeOfAKind
			continue
		}

		if valWithJoker == 2 && length == 3 && TwoPair > bestType {
			bestType = TwoPair
			continue
		}

		if valWithJoker == 2 && length == 4 && OnePair > bestType {
			bestType = OnePair
			continue
		}
	}

	return bestType
}
