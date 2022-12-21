package services

import (
	"encoding/json"
	"log"

	"github.com/smg/easy-bazaar/models"
	"github.com/smg/easy-bazaar/repo"
)

func (i *BazaarService) GetItems(pi, ps int) []models.Item {
	items := i.getAllItems()
	fromIndex := pi * ps
	toIndex := fromIndex + ps

	return items[fromIndex:toIndex]
}

func (i *BazaarService) getAllItems() []models.Item {
	byteResult := repo.ReadItemsData()
	var items []models.Item
	err := json.Unmarshal(byteResult, &items)
	if err != nil {
		log.Fatal(err)
	}

	return items
}

func (i *BazaarService) GetItem(id int) models.Item {
	items := i.getAllItems()

	for _, it := range items {
		if it.ID == id {
			return it
		}
	}

	return models.Item{}
}
