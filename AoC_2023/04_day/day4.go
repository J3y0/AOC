package main

import (
	"04_day/hash_table"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./data/day4.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while reading file: %v", err)
		os.Exit(1)
	}

	part1, err := totalPoints(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while solving part1: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Result for part 1: %d\n", part1)

	part2, err := totalCards(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error received while solving part2: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Result for part 2: %d\n", part2)
}

func pointsFor1Card(line string) (totalMatch int, err error) {
	numbers := strings.Split(line, " | ")
	winning_nbs, err := toArray(numbers[0])
	if err != nil {
		return 0, err
	}
	card_nbs, err := toArray(numbers[1])
	if err != nil {
		return 0, err
	}
	// Insert all winning nbs in Hash Table
	ht := hash_table.NewHashTable(len(winning_nbs))
	for _, nb := range winning_nbs {
		ht.Insert(nb, nb)
	}
	// Get each card_nbs in for loop in hash table.
	// If it exists this is a winning number
	// Else it is not
	for _, card_nb := range card_nbs {
		if _, ok := ht.Get(card_nb); ok {
			totalMatch += 1
		}
	}
	return
}

func totalPoints(data []byte) (int, error) {
	lines := strings.Split(string(data), "\n")

	var sum int
	for _, line := range lines {
		numbers := strings.Split(line, ": ")[1]
		totalMatch, err := pointsFor1Card(numbers)
		if err != nil {
			return 0, err
		}
		sum += computeScore(totalMatch)
	}

	return sum, nil
}

func computeScore(totalMatch int) int {
	if totalMatch == 0 {
		return 0
	}
	return 1 << (totalMatch - 1)
}

func toArray(numbers string) ([]int, error) {
	var result []int
	nb_regex := regexp.MustCompile(`\d+`)
	splitted := nb_regex.FindAllString(numbers, -1)
	for _, nb := range splitted {
		nb, err := strconv.Atoi(nb)
		if err != nil {
			return nil, err
		}
		result = append(result, nb)
	}
	return result, nil
}

func totalCards(data []byte) (int, error) {
	lines := strings.Split(string(data), "\n")

	// Init map for number of occurences for each card id
	idTotalCopy := make(map[int]int, len(lines))
	for j := 0; j < len(lines); j++ {
		idTotalCopy[j+1] = 1
	}

	for id, line := range lines {
		numbers := strings.Split(line, ": ")[1]
		totalPoints, err := pointsFor1Card(numbers)
		if err != nil {
			return 0, err
		}

		nbOccurences := idTotalCopy[id+1]
		for i := 1; i <= totalPoints; i++ {
			idTotalCopy[id+1+i] += nbOccurences
		}
	}

	var sum int
	for _, value := range idTotalCopy {
		sum += value
	}
	return sum, nil
}
