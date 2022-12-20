package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smg/easy-bazaar/services"
)

type handlers struct {
}

func (h *handlers) getItems(c *gin.Context) {
	pageIndexParam := c.DefaultQuery("pi", "0")
	pageSizeParam := c.DefaultQuery("ps", "10")

	pi, _ := strconv.Atoi(pageIndexParam)
	ps, _ := strconv.Atoi(pageSizeParam)
	items := services.GetItems(pi, ps)

	c.JSON(200, items)
}
