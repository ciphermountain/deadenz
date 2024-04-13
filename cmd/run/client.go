package run

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ciphermountain/deadenz/internal/listeners"
	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/service/core"
)

var (
	runClient = &cobra.Command{
		Use:   "client",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			addr := fmt.Sprintf("%s:%d", host, port)

			// create grpc client
			client, err := core.NewClient(addr)
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "no client connection: %s", err.Error())
				os.Exit(1)
			}

			// start command loop
			commands := listeners.NewCommandEvent(deadenz.SpawninCommandType, listeners.EnglishCommands)

			// TODO: start a process to listen for events from the multiverse

			// TODO: get profile from another source
			profile := components.Profile{
				UUID:          "1",
				XP:            0,
				Currency:      0,
				BackpackLimit: 10,
				Backpack:      []components.ItemType{},
				Stats: components.Stats{
					Wit:   1,
					Skill: 1,
					Humor: 1,
				},
			}
			defaultCmd := deadenz.WalkCommandType

			if profile.Active == nil {
				defaultCmd = deadenz.SpawninCommandType
			}

			commands.SetDefaultCommand(defaultCmd)

			for {
				input := <-commands.Next()

				switch input {
				case deadenz.SpawninCommandType, deadenz.WalkCommandType:
					// action commands get routed to the game service
					next := runActionCommand(cmd, client, input, profile)
					if next != nil {
						commands.SetDefaultCommand(*next)
					}
				case deadenz.BackpackCommandType, deadenz.CurrencyCommandType, deadenz.XPCommandType:
					// data read commands can be run directly on the client
					runDataReadCommand(cmd, client, input, profile)
				case deadenz.ExitCommandType:
					if err := client.Close(); err != nil {
						fmt.Fprintf(cmd.ErrOrStderr(), "client exited with error: %s", err.Error())
						os.Exit(1)
					}

					fmt.Fprintln(cmd.OutOrStdout(), "client exited successfully")
					os.Exit(0)
				default:
					fmt.Fprintln(cmd.ErrOrStderr(), "unrecognized command")
				}
			}
		},
	}
)

func runActionCommand(
	cmd *cobra.Command,
	client *core.Client,
	input deadenz.CommandType,
	profile components.Profile,
) *deadenz.CommandType {
	var (
		next   deadenz.CommandType
		events []string
		err    error
	)

	switch input {
	case deadenz.SpawninCommandType:
		events, profile, err = client.Spawnin(context.Background(), profile)
		next = deadenz.WalkCommandType
	case deadenz.WalkCommandType:
		events, profile, err = client.Walk(context.Background(), profile)
		next = deadenz.WalkCommandType

		if profile.Active == nil {
			next = deadenz.SpawninCommandType
		}
	default:
		return nil
	}

	if err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "%s", err.Error())

		return nil
	}

	for _, event := range events {
		fmt.Fprintln(cmd.OutOrStdout(), event)
	}

	return &next
}

func runDataReadCommand(
	cmd *cobra.Command,
	client *core.Client,
	input deadenz.CommandType,
	profile components.Profile,
) {
	switch input {
	case deadenz.BackpackCommandType:
		if len(profile.Backpack) == 0 {
			fmt.Println("you have no items in your backpack")
		} else {
			items, err := client.Items(context.Background())
			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err.Error())

				return
			}

			fmt.Println("your backpack includes:")

			for _, itemType := range profile.Backpack {
				for _, item := range items {
					if item.Type == itemType {
						fmt.Fprintln(cmd.OutOrStdout(), item.Name)

						continue
					}
				}
			}
		}
	case deadenz.XPCommandType:
		fmt.Fprintf(cmd.OutOrStdout(), "you have %d xp\n", profile.XP)
	case deadenz.CurrencyCommandType:
		fmt.Fprintf(cmd.OutOrStdout(), "you have %d currency\n", profile.Currency)
	default:
		return
	}
}
