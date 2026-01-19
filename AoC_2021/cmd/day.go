package cmd

import (
	"fmt"
	"main/days"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	partNb int
	dayCmd = &cobra.Command{
		Use:          "day <1-25>",
		Short:        "day's number solution to run",
		Example:      "aoc day 12",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			dayNb, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			if dayNb < 1 || dayNb > 25 {
				return fmt.Errorf("invalid day %d (should be between 1 and 25)", dayNb)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			dayNb, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return days.RunSolution(dayNb)
		},
	}
)

func init() {
	rootCmd.AddCommand(dayCmd)
}
