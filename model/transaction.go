package model

type Transaction struct {
	ID int
	ProductId int `json:"product_id"`
	Qty int	`json:"qty"`
	IsOut bool `json:"is_out"`
}