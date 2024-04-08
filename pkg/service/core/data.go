package core

import (
	"context"
	"os"

	"github.com/ciphermountain/deadenz/pkg/components"
)

type DataLoader interface {
	Data() ([]byte, error)
}

type Loaders struct {
	Items              DataLoader
	Characters         DataLoader
	ItemDecisionEvents DataLoader
	ActionEvents       DataLoader
	MutationEvents     DataLoader
	EncounterEvents    DataLoader
}

type Data struct{}

func (d *Data) Items() []components.Item {
	return nil
}

type FileLoader struct {
	Path string `json:"path"`
}

func NewFileLoader(path string) *FileLoader {
	return &FileLoader{Path: path}
}

func (l *FileLoader) Data(_ context.Context) ([]byte, error) {
	dat, err := os.ReadFile(l.Path)
	if err != nil {
		return nil, err
	}

	return dat, nil
}
