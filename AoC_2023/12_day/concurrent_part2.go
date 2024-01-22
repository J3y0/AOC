package main

import (
	"fmt"
	"strings"
	"sync"
)

// Cache will store temp result of recurrent function so if one case is
// encountered again, just return the stored value and do not compute it again
var (
	lock  sync.Mutex
	cache = make(map[string]int)
)

func SumArrangementsConcurrently(lines []string) int {
	var (
		wg            sync.WaitGroup
		resultChannel = make(chan int, len(lines))
		taskChannel   = make(chan string, len(lines))
		poolSize      = 4
	)

	// Add and launch workers
	wg.Add(poolSize)
	for i := 0; i < poolSize; i++ {
		go Worker(taskChannel, resultChannel, &wg)
	}

	// Enqueue tasks
	for _, line := range lines {
		taskChannel <- line
	}
	close(taskChannel)

	go func() {
		wg.Wait()            // Wait for all goroutine to finish
		close(resultChannel) // Close the channel as no more data can be received
	}()

	var sum int
	for result := range resultChannel {
		sum += result
	}

	return sum
}

func Worker(taskChannel <-chan string, resultChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for line := range taskChannel {
		temp := strings.Split(line, " ")
		// Parse lines and repeat sequence of springs and of contiguous length 5 times
		springs := strings.Repeat(temp[0]+"?", 4) + temp[0]
		lengthContiguous, _ := ParseToNumbers(strings.Repeat(temp[1]+",", 4) + temp[1])

		arrangements := ConcurrentFindArrangementsWithCache(springs, lengthContiguous)
		resultChannel <- arrangements
	}
}

func ConcurrentFindArrangementsWithCache(springs string, lengthContiguous []int) int {
	// Base case
	if len(lengthContiguous) == 0 {
		if strings.Contains(springs, "#") {
			return 0 // Invalid
		} else {
			return 1 // Valid as no springs left or only . or ? that can be cast in .
		}
	}
	if springs == "" {
		if len(lengthContiguous) > 0 {
			return 0 // Invalid
		} else {
			return 1
		}
	}

	key := fmt.Sprintf("%s:%v", springs, lengthContiguous)
	lock.Lock()
	if value, ok := cache[key]; ok {
		lock.Unlock()
		return value
	}
	lock.Unlock()

	var arrangements int
	if springs[0] == '.' || (springs[0] == '?') {
		arrangements += ConcurrentFindArrangementsWithCache(springs[1:], lengthContiguous)
	}

	if springs[0] == '#' || springs[0] == '?' {
		// Check we have enough springs to respect next length of contiguous broken springs
		// Check the substring of size length is contiguous
		// Check the character after the substring is not another broken spring
		if len(springs) > lengthContiguous[0] && IsContiguous(springs[:lengthContiguous[0]]) && springs[lengthContiguous[0]] != '#' {
			arrangements += ConcurrentFindArrangementsWithCache(springs[lengthContiguous[0]+1:], lengthContiguous[1:])
		} else if len(springs) == lengthContiguous[0] && IsContiguous(springs[:lengthContiguous[0]]) {
			// If numbers of springs is exactly equal to the next length of contiguous broken springs
			// Then just check if this is contiguous and pass empty string to the next call
			// (no more springs, just trigger base case)
			arrangements += ConcurrentFindArrangementsWithCache("", lengthContiguous[1:])
		}
	}

	lock.Lock()
	cache[key] = arrangements
	lock.Unlock()

	return arrangements
}
