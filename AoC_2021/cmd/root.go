package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	dayNb   int
	rootCmd = &cobra.Command{
		Use:               "aoc",
		Short:             "aoc - run a selected solution from CLI",
		Long:              "aoc can be used to run a selected day's solution for the year 2021 of Advent of Code",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
