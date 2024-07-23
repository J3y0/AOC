package cmd

import (
	"github.com/spf13/cobra"
	"main/days"
	"strconv"
)

type NoDayArgumentError struct{}

func (e *NoDayArgumentError) Error() string {
	return "Invalid argument: a day should be provided.\n"
}

var (
	partNb int
	dayCmd = &cobra.Command{
		Use:   "day",
		Short: "day's number solution to run",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return &NoDayArgumentError{}
			}

			dayNb, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			if err := days.RunSelectedSolution(dayNb, partNb); err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(dayCmd)
	dayCmd.PersistentFlags().IntVarP(&partNb, "part", "p", 0, "Define which part to run (1 or 2). By default it is equal to 0, meaning both are run.")
}
