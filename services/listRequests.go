package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/smg/easy-bazaar/models"
)

func (b *BazaarService) GetRequests(ctx context.Context, userId int) []models.UserItemResponse {
	var results = make([]models.UserItemResponse, 0)

	requests, err := b.Client.Keys(ctx, fmt.Sprintf(models.USER_TRACKING_ID2, userId, "*")).Result()
	if err != nil {
		fmt.Printf(`get requests  error %v`, err)
		return results
	}

	for _, key := range requests {
		var uq models.UserItem
		err := json.Unmarshal([]byte(b.Client.Get(ctx, key).Val()), &uq)
		if err != nil {
			fmt.Printf(`unmarshal get requests  error %v`, err)
			return results
		}

		userName := b.GetUser(uq.UserId).Name
		itemName := b.GetItem(uq.ItemId).Item

		uqr := models.UserItemResponse{
			UserName: userName,
			ItemName: itemName,
			UserItem: uq,
		}

		results = append(results, uqr)
	}

	return results
}

func (b *BazaarService) GetAllRequests(ctx context.Context) []models.UserItemResponse {
	var results = make([]models.UserItemResponse, 0)

	iter := b.Client.Scan(ctx, 0, "prefix:*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		var uq models.UserItem
		err := json.Unmarshal([]byte(b.Client.Get(ctx, key).Val()), &uq)
		if err != nil {
			fmt.Printf(`unmarshal get requests  error %v`, err)
			return results
		}

		userName := b.GetUser(uq.UserId).Name
		itemName := b.GetItem(uq.ItemId).Item

		uqr := models.UserItemResponse{
			UserName: userName,
			ItemName: itemName,
			UserItem: uq,
		}

		results = append(results, uqr)
	}
	return results
}
