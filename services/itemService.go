package services

import (
	"encoding/json"
	"log"

	"github.com/smg/easy-bazaar/models"
	"github.com/smg/easy-bazaar/repo"
)

func GetItems(pi, ps int) []models.Item {
	items := GetAll()
	fromIndex := pi * ps
	toIndex := fromIndex + ps

	return items[fromIndex:toIndex]
}

func GetAll() []models.Item {
	byteResult := repo.ReadDataFile()
	var items []models.Item
	err := json.Unmarshal(byteResult, &items)
	if err != nil {
		log.Fatal(err)
	}

	return items
}
