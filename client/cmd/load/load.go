package load

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	LoadCmd = &cobra.Command{
		Use:   "load",
		Short: "Load resources for the deadenz game.",
		Long:  ``,
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}

			if _, err := typeFromString(args[0]); err != nil {
				return fmt.Errorf("invalid value type: %s", args[0])
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = typeFromString(args[0])
		},
	}
)

type LoadValue int

const (
	LoadValueItems LoadValue = iota
)

func typeFromString(value string) (LoadValue, error) {
	switch value {
	case "items":
		return LoadValueItems, nil
	}

	return LoadValue(-1), fmt.Errorf("invalid type")
}
