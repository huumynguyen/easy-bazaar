package models

import (
	"time"
)

type Item struct {
	ID                 int
	Item               string
	Picture            string
	AvailableItem      int
	AvailableQuantity  string
	AvailableStatus    interface{}
	Description        string
	Category           string
	Note               string
	Contact            string
	Created            time.Time
	Duration           string // days
	Keeper             string
	NoteWhoIsBorrowing string
	Popularity         int
	PriceDay           interface{} // for borrow
	ProductLink        string
	Rating             string
	Value              string
}

type User struct {
	ID   int
	Name string
	Role string
}

type UserItem struct {
	UserId   int
	ItemId   int
	FromDate int
	ToDate   int
	Status   string // Approved, Pending, Rejected
	Created  time.Time
}

type UserItemResponse struct {
	UserName string
	ItemName string
	Picture  string
	Created  string
	UserItem UserItem
}
