package repository

import (
	"Smarket/internal/db"
	"Smarket/models"
	"Smarket/pkg/logger"
)

func GetAllSales() ([]models.Sale, error) {
	logger.Info.Println("repository.GetAllSales: fetching all sales")

	var sales []models.Sale
	query := `
		SELECT id, user_id, total_sum, created_at, updated_at
	    FROM sales`

	err := db.GetDBConn().Select(&sales, query)
	if err != nil {
		logger.Error.Printf("repository.GetAllSales: failed to fetch sales: %v", err)
		return nil, translateError(err)
	}

	logger.Info.Printf("repository.GetAllSales: fetched %d sales", len(sales))
	return sales, nil
}

func GetSaleByID(id int) (models.Sale, error) {
	logger.Info.Printf("repository.GetSaleByID: fetching sale with ID %d", id)

	var sale models.Sale
	query := `
		SELECT id, user_id, total_sum, created_at, updated_at
	    FROM sales
	    WHERE id = $1`

	err := db.GetDBConn().Get(&sale, query, id)
	if err != nil {
		logger.Error.Printf("repository.GetSaleByID: failed to fetch sale ID %d: %v", id, err)
		return models.Sale{}, translateError(err)
	}

	logger.Info.Printf("repository.GetSaleByID: found sale ID %d", id)
	return sale, nil
}

func CreateSale(sale models.Sale) (int, error) {
	logger.Info.Printf("repository.CreateSale: creating sale: %+v", sale)

	query := `
		INSERT INTO sales(user_id, total_sum)
	    VALUES ($1, $2)
	RETURNING id`

	var id int
	err := db.GetDBConn().Get(&id, query, sale.UserID, sale.TotalSum)
	if err != nil {
		logger.Error.Printf("repository.CreateSale: failed to create sale: %v", err)
		return id, translateError(err)
	}

	logger.Info.Printf("repository.CreateSale: sale created with ID %d", id)
	return id, nil
}

func UpdateSale(id int, sale models.Sale) (models.Sale, error) {
	logger.Info.Printf("repository.UpdateSale: updating sale ID %d with data: %+v", id, sale)

	query := `
		UPDATE sales
	SET user_id = $1,
		total_sum = $2,
		updated_at = NOW()
	WHERE id = $3
	RETURNING id, user_id, total_sum, created_at, updated_at`

	var updated models.Sale
	err := db.GetDBConn().Get(&updated, query, sale.UserID, sale.TotalSum, id)
	if err != nil {
		logger.Error.Printf("repository.UpdateSale: failed to update sale ID %d: %v", id, err)
		return models.Sale{}, translateError(err)
	}

	logger.Info.Printf("repository.UpdateSale: sale ID %d updated successfully", id)
	return updated, nil
}

func DeleteSale(id int) error {
	logger.Info.Printf("repository.DeleteSale: deleting sale ID %d", id)

	query := `DELETE FROM sales 
              WHERE id = $1`
	_, err := db.GetDBConn().Exec(query, id)
	if err != nil {
		logger.Error.Printf("repository.DeleteSale: failed to delete sale ID %d: %v", id, err)
	}
	return translateError(err)
}

func GetSaleReceipt(saleID int) (models.Receipt, error) {
	logger.Info.Printf("repository.GetSaleReceipt: generating receipt for sale ID %d", saleID)

	var receipt models.Receipt
	receipt.SaleID = saleID

	dateQuery := `
		SELECT s.created_at AS date, u.name AS cashier
	FROM sales s
	JOIN users u ON s.user_id = u.id
	WHERE s.id = $1`

	err := db.GetDBConn().Get(&receipt, dateQuery, saleID)
	if err != nil {
		logger.Error.Printf("repository.GetSaleReceipt: failed to get base receipt info: %v", err)
		return receipt, translateError(err)
	}

	itemsQuery := `
		SELECT
	p.name AS product_name,
		si.quantity,
		p.price AS unit_price,
		si.price AS total_price
	FROM sale_items si
	JOIN products p ON si.product_id = p.id
	WHERE si.sale_id = $1`

	var items []models.ReceiptItem
	err = db.GetDBConn().Select(&items, itemsQuery, saleID)
	if err != nil {
		logger.Error.Printf("repository.GetSaleReceipt: failed to get receipt items: %v", err)
		return receipt, translateError(err)
	}

	receipt.Items = items
	for _, item := range items {
		receipt.TotalSum += item.TotalPrice
	}

	logger.Info.Printf("repository.GetSaleReceipt: receipt for sale ID %d generated successfully", saleID)
	return receipt, nil
}
