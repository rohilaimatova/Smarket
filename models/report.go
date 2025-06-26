package models

type ProductReport struct {
	ProductName   string  `json:"product_name" db:"product_name"`
	TotalQuantity int     `json:"total_quantity" db:"total_quantity"`
	TotalAmount   float64 `json:"total_amount" db:"total_amount"`
}

type CashierSalesReport struct {
	Cashier     string          `json:"cashier" db:"cashier"`
	Products    []ProductReport `json:"products" db:"products"`
	TotalItems  int             `json:"total_items" db:"total_items"`
	TotalAmount float64         `json:"total_amount" db:"total_amount"`
	SalesCount  int             `json:"sales_count" db:"sales_count"`
}
