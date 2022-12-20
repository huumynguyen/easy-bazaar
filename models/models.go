package models

import (
	"time"
)

type Item struct {
	ItemName          string
	Picture           string
	AvailableItem     int
	AvailableQuantity int
	AvailableStatus   string
	Description       string
	Category          string
	Note              string
	Contact           string
	Created           time.Time
	Duration          string // days
	Keeper            string
	WhoIsBorrowing    string
	Popularity        int
	PriceDay          string // for borrow
	ProductLink       string
	Rating            string
	Value             string
}
