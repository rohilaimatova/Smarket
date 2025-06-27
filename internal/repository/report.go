package repository

import (
	"Smarket/internal/db"
	"Smarket/models"
	"Smarket/pkg/logger"
)

func GetSalesReport(fromDate, toDate string) (models.Report, error) {
	var (
		report         models.Report
		salesReport    []models.SalesReport
		productsReport []models.ProductReport
		totalItems     int
		salesCount     int
		totalAmount    float64
	)

	db := db.GetDBConn()

	salesQuery := ` select s.id, u.name AS cashier, s.total_sum
       from sales s
       join users u on u.id = s.user_id
       where s.created_at BETWEEN $1 AND $2::pg_catalog.timestamp + interval '1 DAY'`

	err := db.Select(&salesReport, salesQuery, fromDate, toDate)
	if err != nil {
		logger.Error.Printf("repository.GetSalesReport(): failed to fetch sales: %v\n", err)
		return report, err
	}

	saleItemsQuery := `select p.name AS product_name, p.price,  si.quantity
      FROM sale_items si
      JOIN products p ON si.product_id = p.id
      WHERE si.sale_id = $1`

	for _, sale := range salesReport {
		err = db.Select(&productsReport, saleItemsQuery, sale.ID)
		if err != nil {
			logger.Error.Printf("repository.GetSalesReport(): failed to fetch saleItems: %v\n", err)
		}

		sale.Products = productsReport
	}

	summaryQuery := `
			SELECT
		    COALESCE(SUM(si.quantity), 0) AS total_items,
			COALESCE(SUM(si.price), 0) AS total_amount,
			COUNT(DISTINCT s.id) AS sales_count
		FROM sale_items si
		JOIN sales s ON s.id = si.sale_id
		JOIN users u ON s.user_id = u.id
		WHERE s.created_at BETWEEN $1 AND $2::pg_catalog.timestamp + interval '1 DAY'`

	err = db.QueryRow(summaryQuery, fromDate, toDate).Scan(&totalItems, &totalAmount, &salesCount)
	if err != nil {
		logger.Error.Printf("repository.GetSalesReport(): failed to fetch summery counts: %v\n", err)
		return report, err
	}

	report.Sales = salesReport
	report.TotalSalesAmount = totalAmount
	report.SalesCount = salesCount
	report.TotalProductCount = totalItems

	return report, nil

}
