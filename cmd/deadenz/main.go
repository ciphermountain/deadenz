package main

import (
	"context"
	"encoding/json"
	"fmt"

	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/actions"
	"github.com/ciphermountain/deadenz/pkg/events"
	"github.com/ciphermountain/deadenz/pkg/multiverse"
	"github.com/ciphermountain/deadenz/pkg/multiverse/service"
)

func main() {
	fmt.Println("This is the console version of the game and will take input from stdin!")

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
	config := Config{
		ItemsSource:              NewFileLoader("./assets/default_items.json"),
		CharactersSource:         NewFileLoader("./assets/default_characters.json"),
		ItemDecisionEventsSource: NewFileLoader("./assets/default_item_decision_events.json"),
		ActionEventsSource:       NewFileLoader("./assets/default_action_events.json"),
		MutationEventsSource:     NewFileLoader("./assets/default_mutation_events.json"),
		EncounterEventsSource:    NewFileLoader("./assets/default_encounter_events.json"),
	}

	loadItems(action, config)
	loadCharacters(action, config)
	loadItemDecisionEvents(action, config)
	loadActionEvents(action, config)
	loadMutationEvents(action, config)
	loadEncounterEvents(action, config)

	client, err := multiverse.NewMultiverseClient(":9090")
	if err != nil {
		panic(err)
	}

	commands := NewCommandEventListener("spawnin")

	multi, err := NewMultiverseMessageListener(client)
	if err != nil {
		panic(err)
	}

CommandLoop:
	for {
		select {
		case evt := <-multi.Next():
			fmt.Printf("event: %+v\n", evt)
			continue
		case input := <-commands.Next():
			var evts []events.Event

			switch input {
			case "spawnin":
				var err error

				profile, evts, err = action.Spawn(profile)
				if err != nil {
					fmt.Println(err.Error())
				}

				commands.SetDefaultCommand("walk")
			case "walk":
				var err error

				profile, evts, err = action.Walk(profile)
				if err != nil {
					fmt.Println(err.Error())
				}

				for _, evt := range evts {
					switch evt.(type) {
					case events.DieMutationEvent:
						bts, err := json.Marshal(evt)
						if err != nil {
							continue
						}

						client.PublishEvent(context.Background(), &service.Event{
							Data: bts,
						})
					}
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
				commands.SetDefaultCommand("spawnin")
			}

			fmt.Println("")
		}
	}
}
