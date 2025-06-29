package models

import "time"

type Sale struct {
	ID        int        `json:"id" db:"id"`
	UserID    int        `json:"user_id" db:"user_id"`
	TotalSum  float64    `json:"total_sum" db:"total_sum"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type CreateSaleRequest struct {
	UserId   int            `json:"user_id"`
	Products []ProductItems `json:"products"`
}

type ProductItems struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}

type SaleRequest struct {
	Products []ProductItems `json:"products"`
}
