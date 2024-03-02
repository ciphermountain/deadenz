package main

import (
	"os"

	"github.com/ciphermountain/deadenz/pkg/actions"
	"github.com/ciphermountain/deadenz/pkg/characters"
	"github.com/ciphermountain/deadenz/pkg/events"
	"github.com/ciphermountain/deadenz/pkg/items"
)

func loadItems(data *actions.WithData, path string) {
	// load data from json file
	dat, err := os.ReadFile(path)
	if err != nil {
		os.Exit(1)
	}

	it, err := items.LoadItems(dat)
	if err != nil {
		os.Exit(1)
	}

	data.Items = it
}

func loadCharacters(data *actions.WithData, path string) {
	// load data from json file
	dat, err := os.ReadFile(path)
	if err != nil {
		os.Exit(1)
	}

	it, err := characters.Load(dat)
	if err != nil {
		os.Exit(1)
	}

	data.Characters = it
}

func loadItemDecisionEvents(data *actions.WithData, path string) {
	// load data from json file
	dat, err := os.ReadFile(path)
	if err != nil {
		os.Exit(1)
	}

	it, err := events.LoadItemDecisions(dat)
	if err != nil {
		os.Exit(1)
	}

	data.ItemDecisions = it
}
