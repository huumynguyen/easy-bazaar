package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/smg/easy-bazaar/services"
)

type handlers struct {
	bazaarService services.BazaarService
}

// getItems godoc
// @Summary      Get items
// @Description  get items
// @Tags         Root
// @Accept       json
// @Produce      json
// @Param pi query int false "page index" default(0)
// @Param ps query int false "page size" default(20)
// @Success      200 {object} []models.Item
// @Router       /getItems [get]
func (h *handlers) getItems(c *gin.Context) {
	pageIndexParam := c.DefaultQuery("pi", "0")
	pageSizeParam := c.DefaultQuery("ps", "10")

	pi, _ := strconv.Atoi(pageIndexParam)
	ps, _ := strconv.Atoi(pageSizeParam)
	items := h.bazaarService.GetItems(pi, ps)

	c.JSON(200, items)
}

// getItem godoc
// @Summary      Get item
// @Description  get item
// @Tags         Root
// @Accept       json
// @Produce      json
// @Param id query int true "item id"
// @Success      200 {object} models.Item
// @Router       /getItem [get]
func (h *handlers) getItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	item := h.bazaarService.GetItem(id)

	c.JSON(200, item)
}

// postRequest godoc
// @Summary      Tracking borrowed items
// @Description  Tracking borrowed items
// @Tags         Root
// @Accept       json
// @Produce      json
// @Param userId query int true "user id"
// @Param itemId query int true "item id"
// @Param df query int true "date to"
// @Param dt query int true "date from"
// @Success      200
// @Router       /postRequest [POST]
func (h *handlers) postRequest(c *gin.Context) {
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	itemId, _ := strconv.Atoi(c.PostForm("itemId"))
	df, _ := strconv.Atoi(c.PostForm("df"))
	dt, _ := strconv.Atoi(c.PostForm("dt"))
	st := c.PostForm("status")

	err := h.bazaarService.SaveBorrowedItem(userId, itemId, df, dt, st)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, nil)
		return
	}

	c.JSON(200, nil)
}

// listRequests godoc
// @Summary      Get list requests
// @Description   Get list requests
// @Tags         Root
// @Accept       json
// @Produce      json
// @Param userId query int true "user id"
// @Success      200 {object} []models.UserItemResponse
// @Router       /listRequests [get]
func (h *handlers) getRequests(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("userId"))

	if userId > 0 {
		res := h.bazaarService.GetRequests(context.Background(), userId)
		c.JSON(200, res)
	} else {
		res := h.bazaarService.GetAllRequests(context.Background())
		c.JSON(200, res)
	}
}

// getUser godoc
// @Summary      Get user
// @Description   Get user
// @Tags         Root
// @Accept       json
// @Produce      json
// @Param userId query int true "user id"
// @Success      200 {object} []models.User
// @Router       /getUser [get]
func (h *handlers) getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("userId"))
	item := h.bazaarService.GetUser(id)

	c.JSON(200, item)
}
