package services

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"

	"github.com/smg/easy-bazaar/models"
	"github.com/smg/easy-bazaar/repo"
)

type ItemService struct {
	Client *redis.Client
}

func (i *ItemService) GetItems(pi, ps int) []models.Item {
	items := i.getAll()
	fromIndex := pi * ps
	toIndex := fromIndex + ps

	return items[fromIndex:toIndex]
}

func (i *ItemService) getAll() []models.Item {
	byteResult := repo.ReadDataFile()
	var items []models.Item
	err := json.Unmarshal(byteResult, &items)
	if err != nil {
		log.Fatal(err)
	}

	return items
}

func (i *ItemService) GetItem(id int) models.Item {
	items := i.getAll()

	for _, it := range items {
		if it.ID == id {
			return it
		}
	}

	return models.Item{}
}
