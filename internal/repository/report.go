package repository

import (
	"Smarket/internal/db"
	"Smarket/models"
	"Smarket/pkg/logger"
)

func GetSalesReport(fromDate, toDate string) ([]models.CashierSalesReport, error) {
	db := db.GetDBConn()

	cashierQuery := `
		SELECT DISTINCT u.name AS cashier
	FROM sales s
	JOIN users u ON s.user_id = u.id
	WHERE s.created_at BETWEEN $1 AND $2`

	var cashiers []string
	err := db.Select(&cashiers, cashierQuery, fromDate, toDate)
	if err != nil {
		logger.Error.Printf("repository.GetSalesReport(): failed to fetch cashiers: %v\n", err)
	}

	var reports []models.CashierSalesReport

	for _, cashier := range cashiers {
		var report models.CashierSalesReport
		report.Cashier = cashier

		summaryQuery := `
			SELECT
		COALESCE(SUM(si.quantity), 0) AS total_items,
			COALESCE(SUM(si.price), 0) AS total_amount,
			COUNT(DISTINCT s.id) AS sales_count
		FROM sale_items si
		JOIN sales s ON s.id = si.sale_id
		JOIN users u ON s.user_id = u.id
		WHERE s.created_at BETWEEN $1 AND $2 AND u.name = $3`

		err = db.Get(&report, summaryQuery, fromDate, toDate, cashier)
		if err != nil {
			logger.Error.Printf("repository.GetSalesReport(): failed to fetch cashiers: %v\n", err)
		}

		productsQuery := `
			SELECT
		p.name AS product_name,
			SUM(si.quantity) AS total_quantity,
			SUM(si.price) AS total_amount
		FROM sale_items si
		JOIN products p ON si.product_id = p.id
		JOIN sales s ON s.id = si.sale_id
		JOIN users u ON s.user_id = u.id
		WHERE s.created_at BETWEEN $1 AND $2 AND u.name = $3
		GROUP BY p.name`

		err = db.Select(&report.Products, productsQuery, fromDate, toDate, cashier)
		if err != nil {
			logger.Error.Printf("repository.GetSalesReport(): failed to fetch products: %v\n", err)
		}

		reports = append(reports, report)
	}
	return reports, nil
}
