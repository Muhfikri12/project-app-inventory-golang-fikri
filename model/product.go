package model

import "time"

type Products struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Stocks int	`json:"stocks"`
	CategoryID int `json:"category_id"`
	Updated_at time.Time
}
type ProductsIs struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Stocks int	`json:"stocks"`
	Category string `json:"category"`
	Updated_at time.Time
}