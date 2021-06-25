package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fsa-authorities",
	Short: "A CLI tool to compare the food hygiene rating distribution local authorities",
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(compareCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
