package models

type ProductReport struct {
	ProductName     string  `json:"product_name" db:"product_name"`
	ProductQuantity int     `json:"total_quantity" db:"quantity"`
	ProductPrice    float64 `json:"total_price" db:"price"`
}

type SalesReport struct {
	ID          int             `json:"id" db:"id"`
	Cashier     string          `json:"cashier" db:"cashier"`
	TotalAmount float64         `json:"total_amount" db:"total_sum"`
	Products    []ProductReport `json:"products"`
}
type Report struct {
	Sales             []SalesReport `json:"sales_report"`
	TotalSalesAmount  float64       `json:"total_sales_amount" db:"total_amount"`
	SalesCount        int           `json:"sales_count" db:"sales_count"`
	TotalProductCount int           `json:"total_product_count" db:"total_items"`
}
