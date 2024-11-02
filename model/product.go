package model

import "time"

type Products struct {
	ID int
	Name string
	Code string
	Stocks int
	CategoryID int
	Updated_at time.Time
}