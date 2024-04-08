package cmd

import (
	"fmt"
	"os"

	"github.com/ciphermountain/deadenz/cmd/run"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(run.RootCmd)
}

var (
	rootCmd = &cobra.Command{
		Short: "commands",
		Long:  "commands",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
