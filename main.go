package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	h := handlers{}
	r := gin.New()
	r.GET("/items", h.getItems)

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
