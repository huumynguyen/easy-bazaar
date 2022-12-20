package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"github.com/smg/easy-bazaar/services"
)

func main() {
	var redisHost string
	runningMode := os.Getenv("RUNNING_MODE")
	fmt.Printf("RUNNING_MODE %s\n", runningMode)
	if runningMode == "production" {
		redisHost = "redis:6379"
	} else {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})
	fmt.Printf("ping redis: %s\n", client.Ping(context.Background()).Val())

	h := handlers{
		itemService: services.ItemService{
			Client: client,
		},
		userService: services.UserService{
			Client: client,
		},
	}
	r := gin.New()
	r.GET("/items", h.getItems)
	r.GET("/item", h.getItem) // item?id=123
	r.POST("/postRequest", h.postRequest)

	r.GET("/listRequests", h.getRequests) // ?userId=123

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func HeathCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	c.JSON(http.StatusOK, res)
}
