package models

import "time"

type Product struct {
	ID         int        `json:"id" db:"id"`
	Name       string     `json:"name	"`
	Price      float64    `json:"price"`
	Quantity   int        `json:"quantity"`
	AddedBy    int        `json:"added_by" db:"added_by"`
	CategoryID int        `json:"category_id" db:"category_id"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdateAt   *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
}
