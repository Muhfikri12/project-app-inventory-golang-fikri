package model

type Inventory struct {
	ProductId int `json:"product_id"`
    Row       int `json:"row"`
    Part      int `json:"part"`
    ID        int `json:"id,omitempty"`
}