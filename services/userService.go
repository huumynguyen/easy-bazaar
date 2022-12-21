package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/smg/easy-bazaar/models"
	"github.com/smg/easy-bazaar/repo"
)

func (i *BazaarService) getAllUsers() []models.User {
	byteResult := repo.ReadUsersData()
	var users []models.User
	err := json.Unmarshal(byteResult, &users)
	if err != nil {
		log.Fatal(err)
	}

	return users
}

func (i *BazaarService) GetUser(id int) models.User {
	users := i.getAllUsers()

	for _, it := range users {
		if it.ID == id {
			return it
		}
	}

	return models.User{}
}

func (i *BazaarService) SaveBorrowedItem(userId, itemId, from, to int) error {
	trackingId := fmt.Sprintf(models.USER_TRACKING_ID, userId, itemId)
	data := models.UserItem{
		UserId:   userId,
		ItemId:   itemId,
		FromDate: from,
		ToDate:   to,
	}

	resultBytes, _ := json.Marshal(data)
	status := i.Client.Set(context.Background(), trackingId, string(resultBytes), 0)
	fmt.Println(status.Val())
	if status.Val() != "OK" {
		return fmt.Errorf("error while set redis.")
	}

	return nil
}
