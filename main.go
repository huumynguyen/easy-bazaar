package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"github.com/smg/easy-bazaar/services"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	h := handlers{
		itemService: services.ItemService{
			Client: client,
		},
	}
	r := gin.New()
	r.GET("/items", h.getItems)
	r.GET("/item", h.getItem) // item?id=123

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
