package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
)

var logger log.Logger = *log.Default()

func handleInput(c echo.Context) error {
	logger.Printf("HandleInput called")
	var input []Pairs

	fileReader, err := os.Open("./data/example.txt")
	if err != nil {
		return err
	}
	input, err = ParseInput(fileReader)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, input)
}

func handlePart1(c echo.Context) error {
	logger.Printf("HandlePart1 called")
	var input []Pairs

	fileReader, err := os.Open("./data/day13.txt")
	if err != nil {
		return err
	}
	input, err = ParseInput(fileReader)
	if err != nil {
		return err
	}
	sum, err := Part1(input)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Part1 result: "+strconv.Itoa(sum)+"\n")
}

func handlePart2(c echo.Context) error {
	logger.Printf("HandlePart2 called")
	var input []Pairs

	fileReader, err := os.Open("./data/day13.txt")
	if err != nil {
		return err
	}
	input, err = ParseInput(fileReader)
	if err != nil {
		return err
	}
    decoderKey, err := Part2(input)
    if err != nil {
        return err
    }
    return c.String(http.StatusOK, "Part2 result:" + strconv.Itoa(decoderKey) + "\n")
}

func main() {
	e := echo.New()

	s := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: e,
	}
	// Set routes
	e.GET("/", handleInput)
	e.GET("/part1", handlePart1)
	e.GET("/part2", handlePart2)

	logger.Printf("Listening on port: 8080")

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		logger.Printf("Server closed")
	} else if err != nil {
		logger.Fatalf("Error while starting the server: %v", err)
	}
}
