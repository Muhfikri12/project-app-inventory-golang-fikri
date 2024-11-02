package model

type Transaction struct {
	ID int
	ProductId int
	Qty int
	IsOut bool
}