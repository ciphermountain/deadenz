package main

import (
	"fmt"
	"os"

	"github.com/ciphermountain/deadenz/pkg/actions"
	"github.com/ciphermountain/deadenz/pkg/characters"
	"github.com/ciphermountain/deadenz/pkg/events"
	"github.com/ciphermountain/deadenz/pkg/items"
)

func loadItems(data *actions.WithData, config Config) {
	dat, err := config.ItemsSource.Data()
	if err != nil {
		os.Exit(1)
	}

	it, err := items.LoadItems(dat)
	if err != nil {
		os.Exit(1)
	}

	data.Items = it
}

func loadCharacters(data *actions.WithData, config Config) {
	dat, err := config.CharactersSource.Data()
	if err != nil {
		os.Exit(1)
	}

	it, err := characters.Load(dat)
	if err != nil {
		os.Exit(1)
	}

	data.Characters = it
}

func loadItemDecisionEvents(data *actions.WithData, config Config) {
	dat, err := config.ItemDecisionEventsSource.Data()
	if err != nil {
		os.Exit(1)
	}

	it, err := events.LoadItemDecisions(dat)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data.ItemDecisions = it
}

func loadActionEvents(data *actions.WithData, config Config) {
	dat, err := config.ActionEventsSource.Data()
	if err != nil {
		os.Exit(1)
	}

	it, err := events.LoadActionEvents(dat)
	if err != nil {
		os.Exit(1)
	}

	data.Actions = it
}

func loadMutationEvents(data *actions.WithData, config Config) {
	dat, err := config.MutationEventsSource.Data()
	if err != nil {
		os.Exit(1)
	}

	live, die, err := events.LoadMutations(dat)
	if err != nil {
		os.Exit(1)
	}

	data.LiveMutations = live
	data.DieMutations = die
}

func loadEncounterEvents(data *actions.WithData, config Config) {
	dat, err := config.EncounterEventsSource.Data()
	if err != nil {
		os.Exit(1)
	}

	it, err := events.LoadEncounterEvents(dat)
	if err != nil {
		os.Exit(1)
	}

	data.EncounterEvents = it
}
