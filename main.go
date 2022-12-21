package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/smg/easy-bazaar/docs"
	"github.com/smg/easy-bazaar/services"
)

// @title           Swagger API easy-bazaar
// @version         1.0
// @description     Swagger API easy-bazaar.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
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
		bazaarService: services.BazaarService{
			Client: client,
		},
	}
	r := gin.New()
	r.GET("/items", h.getItems)
	r.GET("/item", h.getItem) // item?id=123
	r.POST("/postRequest", h.postRequest)

	r.GET("/listRequests", h.getRequests) // ?userId=123
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/user", h.getUser) // ?userId=123
	r.GET("/itembyname", h.GetItemsByName)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
