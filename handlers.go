package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/smg/easy-bazaar/services"
)

type handlers struct {
	itemService services.ItemService
	userService services.UserService
}

func (h *handlers) getItems(c *gin.Context) {
	pageIndexParam := c.DefaultQuery("pi", "0")
	pageSizeParam := c.DefaultQuery("ps", "10")

	pi, _ := strconv.Atoi(pageIndexParam)
	ps, _ := strconv.Atoi(pageSizeParam)
	items := h.itemService.GetItems(pi, ps)

	c.JSON(200, items)
}

func (h *handlers) getItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	item := h.itemService.GetItem(id)

	c.JSON(200, item)
}

func (h *handlers) postRequest(c *gin.Context) {
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	itemId, _ := strconv.Atoi(c.PostForm("itemId"))
	df, _ := strconv.Atoi(c.PostForm("df"))
	dt, _ := strconv.Atoi(c.PostForm("dt"))

	err := h.userService.SaveBorrowedItem(userId, itemId, df, dt)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, nil)
		return
	}

	c.JSON(200, nil)
}
