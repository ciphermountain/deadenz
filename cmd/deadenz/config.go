package main

import (
	"os"
)

type DataLoader interface {
	Data() ([]byte, error)
}

type Config struct {
	ItemsSource              DataLoader
	CharactersSource         DataLoader
	ItemDecisionEventsSource DataLoader
	ActionEventsSource       DataLoader
	MutationEventsSource     DataLoader
	EncounterEventsSource    DataLoader
}

type FileLoader struct {
	Path string `json:"path"`
}

func NewFileLoader(path string) *FileLoader {
	return &FileLoader{Path: path}
}

func (l *FileLoader) Data() ([]byte, error) {
	dat, err := os.ReadFile(l.Path)
	if err != nil {
		return nil, err
	}

	return dat, nil
}
