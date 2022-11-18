package models

type Order struct {
	Products Products
	UserId   int64 `json:"user_id"`
}

type Products []int64
