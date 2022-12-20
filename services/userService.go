package services

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"

	"github.com/smg/easy-bazaar/models"
	"github.com/smg/easy-bazaar/repo"
)

type UserService struct {
	Client *redis.Client
}

func (i *UserService) getAll() []models.User {
	byteResult := repo.ReadUsersData()
	var users []models.User
	err := json.Unmarshal(byteResult, &users)
	if err != nil {
		log.Fatal(err)
	}

	return users
}

func (i *UserService) GetUser(id int) models.User {
	users := i.getAll()

	for _, it := range users {
		if it.ID == id {
			return it
		}
	}

	return models.User{}
}
