package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/actions"
	"github.com/ciphermountain/deadenz/pkg/events"
)

func main() {
	// load default characters from assets

	fmt.Println("This is the console version of the game and will take input from stdin!")

	reader := bufio.NewReader(os.Stdin)

	profile := deadenz.Profile{
		UUID:          "1",
		XP:            0,
		Currency:      0,
		BackpackLimit: 10,
		Backpack:      []deadenz.Item{},
		Stats: deadenz.Stats{
			Wit:   1,
			Skill: 1,
			Humor: 1,
		},
	}

	action := &actions.WithData{}

	loadItems(action, "./assets/default_items.json")
	loadCharacters(action, "./assets/default_characters.json")
	loadItemDecisionEvents(action, "./assets/default_item_decision_events.json")

	defaultAction := "spawnin"

CommandLoop:
	for {
		fmt.Printf("Enter command (%s): ", defaultAction)

		// ReadString will block until the delimiter is entered
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}

		// remove the delimeter from the string
		input = strings.TrimSuffix(input, "\n")

		// set default input
		if len(input) == 0 {
			input = defaultAction
		}

		var evts []events.Event

		switch input {
		case "spawnin":
			var err error

			profile, evts, err = action.Spawn(profile)
			if err != nil {
				fmt.Println(err.Error())
			}

			defaultAction = "walk"
		case "walk":
			var err error

			profile, evts, err = action.Walk(profile)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "backpack":
			if profile.Active == nil {
				fmt.Println(actions.ErrNotSpawnedIn)

				continue CommandLoop
			}

			if len(profile.Backpack) == 0 {
				fmt.Println("you have no items in your backpack")
			} else {
				fmt.Println("your backpack includes:")

				for _, item := range profile.Backpack {
					fmt.Println(item.Name)
				}
			}
		case "xp":
			fmt.Println(profile.XP)
		case "currency":
			fmt.Println(profile.Currency)
		case "exit", "quit":
			break CommandLoop
		default:
			fmt.Println("unrecognized command")
		}

		for _, event := range evts {
			fmt.Println(event)
		}

		if profile.Active == nil {
			defaultAction = "spawnin"
		}

		fmt.Println("")
	}
}
