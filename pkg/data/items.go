package data

import (
	"errors"

	"github.com/ciphermountain/deadenz/pkg/components"
)

type ItemProvider struct {
	loader *DataLoader
}

func NewItemProviderFromLoader(loader *DataLoader) *ItemProvider {
	return &ItemProvider{
		loader: loader,
	}
}

func (p *ItemProvider) Item(iType components.ItemType) (*components.Item, error) {
	var items []components.Item
	if err := p.loader.Load(&items); err != nil {
		return nil, err
	}

	for idx, item := range items {
		if item.Type == iType {
			return &items[idx], nil
		}
	}

	return nil, errors.New("item not found")
}

func (p *ItemProvider) Items() ([]components.Item, error) {
	var items []components.Item
	if err := p.loader.Load(&items); err != nil {
		return nil, err
	}

	return items, nil
}
