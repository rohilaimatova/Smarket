package models

import "time"

type SaleItem struct {
	ID        int       `json:"id" db:"id"`
	SaleID    int       `json:"sale_id" db:"sale_id"`
	ProductID int       `json:"product_id" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UpdateSaleItemRequest struct {
	SaleID    int     `json:"sale_id" db:"sale_id"`
	ProductID int     `json:"product_id" db:"product_id"`
	Quantity  int     `json:"quantity" db:"quantity"`
	Price     float64 `json:"price" db:"price"`
}
