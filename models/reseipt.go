package models

import "time"

type Receipt struct {
	SaleID   int           `json:"id" db:"id"`
	Cashier  string        `json:"cashier" db:"cashier"`
	Date     time.Time     `json:"date" db:"date"`
	Items    []ReceiptItem `json:"items" db:"items"`
	TotalSum float64       `json:"total_sum" db:"total_sum"`
}

type ReceiptItem struct {
	ProductName string  `json:"product_name" db:"product_name"`
	Quantity    int     `json:"quantity" db:"quantity"`
	UnitPrice   float64 `json:"unit_price" db:"unit_price"`
	TotalPrice  float64 `json:"total_price" db:"total_price"`
}
