package model

type Response struct {
	StatusCode int
	Message string
	Page int
	Limit int
	TotalItems int
	TotalPages int
	Data interface{}
}

type ResponseCreate struct {
	StatusCode int
	Message string
	Data interface{}
}