package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"

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

func (i *BazaarService) SaveBorrowedItem(userId, itemId, from, to int, requestStatus string) error {
	trackingId := fmt.Sprintf(models.USER_TRACKING_ID, userId, itemId)
	var sttt = models.PendingStatus
	if requestStatus != `` {
		sttt = requestStatus
	}
	data := models.UserItem{
		UserId:   userId,
		ItemId:   itemId,
		FromDate: from,
		ToDate:   to,
		Status:   sttt,
		Created:  time.Now(),
	}

	ctx := context.Background()

	resultBytes, _ := json.Marshal(data)
	status := i.Client.Set(ctx, trackingId, string(resultBytes), 0)
	fmt.Println(status.Val())
	if status.Val() != "OK" {
		return fmt.Errorf("error while set redis.")
	}

	// send to firebase for PN
	fbApp := initializeServiceAccountID(ctx)
	user := i.GetUser(userId)
	var token string
	var msg map[string]string

	item := i.GetItem(itemId)
	if sttt == models.PendingStatus {
		// get token of admin
		token = i.GetUser(0).Token
		msg = map[string]string{
			"title":   "New request",
			"message": fmt.Sprintf(`%v requests to borrow %v`, user.Name, item.Item),
			"type":    "request", // TODO: this is using for loading pages
		}
	} else {
		// get token of user
		token = user.Token
		msg = map[string]string{
			"title":   fmt.Sprintf(`%v request`, sttt),
			"message": fmt.Sprintf(`The %v you borrowed was %v`, item.Item, sttt),
			"type":    "request", // TODO: this is using for loading pages
		}
	}

	sendToToken(ctx, fbApp, token, msg)

	return nil
}

func initializeServiceAccountID(ctx context.Context) *firebase.App {
	// [START initialize_sdk_with_service_account_id]
	opt := option.WithCredentialsFile("repo/serviceAccountKey.json")
	conf := &firebase.Config{
		ServiceAccountID: "firebase-adminsdk-67qi2@easy-bazaar-nvg.iam.gserviceaccount.com",
		ProjectID:        "easy-bazaar-nvg",
	}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// [END initialize_sdk_with_service_account_id]
	return app
}

func sendToToken(ctx context.Context, app *firebase.App, token string, msg map[string]string) {
	// [START send_to_token_golang]
	// Obtain a messaging.Client from the App.
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := token

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data:  msg,
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	// [END send_to_token_golang]
}
