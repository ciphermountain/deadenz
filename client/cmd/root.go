package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ciphermountain/deadenz/client/cmd/load"
)

func init() {
	rootCmd.AddCommand(load.LoadCmd)
}

var (
	rootCmd = &cobra.Command{
		Use: "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.ErrOrStderr(), "invalid command")
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
