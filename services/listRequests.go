package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/smg/easy-bazaar/models"
)

func (b *BazaarService) GetRequests(ctx context.Context, userId int) []models.UserItem {
	var results = make([]models.UserItem, 0)

	requests, err := b.Client.Keys(ctx, fmt.Sprintf(models.USER_TRACKING_ID, userId, "*")).Result()
	if err != nil {
		fmt.Printf(`get requests  error %v`, err)
		return results
	}

	for _, key := range requests {
		var uq models.UserItem
		err := json.Unmarshal([]byte(b.Client.Get(ctx, key).Val()), &uq)
		if err != nil {
			fmt.Printf(`get requests  error %v`, err)
			return results
		}

		results = append(results, uq)
	}

	return results
}
