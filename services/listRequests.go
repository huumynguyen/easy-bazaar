package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
		uqr := b.buildUserItem(ctx, key)

		results = append(results, uqr)
	}

	return results
}

func (b *BazaarService) GetAllRequests(ctx context.Context) []models.UserItemResponse {
	var results = make([]models.UserItemResponse, 0)

	requests, err := b.Client.Keys(ctx, fmt.Sprintf(models.USER_TRACKING_ID2, "*", "*")).Result()
	if err != nil {
		fmt.Printf(`get requests error %v`, err)
		return results
	}

	for _, key := range requests {
		uqr := b.buildUserItem(ctx, key)

		results = append(results, uqr)
	}
	return results
}

func (b *BazaarService) buildUserItem(ctx context.Context, key string) models.UserItemResponse {
	var uq models.UserItem
	err := json.Unmarshal([]byte(b.Client.Get(ctx, key).Val()), &uq)
	if err != nil {
		fmt.Printf(`unmarshal get requests  error %v`, err)
		return models.UserItemResponse{}
	}

	userName := b.GetUser(uq.UserId).Name
	itemObj := b.GetItem(uq.ItemId)
	itemName := itemObj.Item

	pics := itemObj.Picture
	var pic string
	if pics != `` {
		arr := strings.Split(pics, ",")
		if len(arr) > 0 {
			pic = arr[0]
		}
	}

	uqr := models.UserItemResponse{
		UserName: userName,
		ItemName: itemName,
		Picture:  pic,
		Created:  uq.Created.Format("01/02/2006"),
		UserItem: uq,
	}

	return uqr
}
