package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	dayNb   int
	rootCmd = &cobra.Command{
		Use:   "aoc",
		Short: "aoc - run a selected solution from CLI",
		Long:  "aoc can be used to run a selected day's solution for the year 2021 of Advent of Code",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Day selected", dayNb)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error while executing CLI: %v", err)
		os.Exit(1)
	}
}
