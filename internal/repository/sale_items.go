package repository

import (
	"Smarket/internal/db"
	"Smarket/models"
	"Smarket/pkg/logger"
)

func GetAllSaleItems() ([]models.SaleItem, error) {
	logger.Info.Println("repository.GetAllSaleItems: fetching all sale items")

	var items []models.SaleItem
	err := db.GetDBConn().Select(&items,
		`SELECT id, sale_id, product_id, quantity, price, created_at, updated_at 
               FROM sale_items`)
	if err != nil {
		logger.Error.Printf("repository.GetAllSaleItems: failed to fetch sale items: %v", err)
		return nil, translateError(err)
	}

	logger.Info.Printf("repository.GetAllSaleItems: fetched %d sale items", len(items))
	return items, nil
}

func GetSaleItemByID(id int) (models.SaleItem, error) {
	logger.Info.Printf("repository.GetSaleItemByID: fetching sale item with ID %d", id)

	var item models.SaleItem
	err := db.GetDBConn().Get(&item,
		`SELECT id, sale_id, product_id, quantity, price, created_at, updated_at
               FROM sale_items 
               WHERE id = $1, id`)
	if err != nil {
		logger.Error.Printf("repository.GetSaleItemByID: failed to fetch item ID %d: %v", id, err)
		return models.SaleItem{}, translateError(err)
	}

	logger.Info.Printf("repository.GetSaleItemByID: found item ID %d", id)
	return item, nil
}

func CreateSaleItem(item models.SaleItem) error {
	logger.Info.Printf("repository.CreateSaleItem: creating sale item: %+v", item)

	var price float64
	err := db.GetDBConn().Get(&price,
		`SELECT price
               FROM products
               WHERE id = $1, item.ProductID`)
	if err != nil {
		logger.Error.Printf("repository.CreateSaleItem: failed to fetch product price: %v", err)
		return translateError(err)
	}

	total := price * float64(item.Quantity)

	query := `
		INSERT INTO sale_items (sale_id, product_id, quantity, price, created_at, updated_at)
	    VALUES ($1, $2, $3, $4, NOW(), NOW())`
	_, err = db.GetDBConn().Exec(query, item.SaleID, item.ProductID, item.Quantity, total)
	if err != nil {
		logger.Error.Printf("repository.CreateSaleItem: failed to insert sale item: %v", err)
		return translateError(err)
	}

	updateQuery := `UPDATE sales SET total_sum = total_sum + $1 
                    WHERE id = $2`
	_, err = db.GetDBConn().Exec(updateQuery, total, item.SaleID)
	if err != nil {
		logger.Error.Printf("repository.CreateSaleItem: failed to update sales total_sum: %v", err)
		return translateError(err)
	}

	logger.Info.Println("repository.CreateSaleItem: sale item created and total updated successfully")
	return nil
}

func UpdateSaleItem(id int, updated models.SaleItem) (models.SaleItem, error) {
	logger.Info.Printf("repository.UpdateSaleItem: updating sale item ID %d with data: %+v", id, updated)

	query := `
		UPDATE sale_items
	SET sale_id = $1,
		product_id = $2,
		quantity = $3,
		price = $4,
		updated_at = NOW()
	WHERE id = $5
	RETURNING id, sale_id, product_id, quantity, price, updated_at, created_at`

	var item models.SaleItem
	err := db.GetDBConn().Get(&item, query, updated.SaleID, updated.ProductID, updated.Quantity, updated.Price, id)
	if err != nil {
		logger.Error.Printf("repository.UpdateSaleItem: failed to update item ID %d: %v", id, err)
		return models.SaleItem{}, translateError(err)
	}

	logger.Info.Printf("repository.UpdateSaleItem: item ID %d updated successfully", id)
	return item, nil
}

func DeleteSaleItem(id int) error {
	logger.Info.Printf("repository.DeleteSaleItem: deleting sale item ID %d", id)

	_, err := db.GetDBConn().Exec(
		`DELETE FROM sale_items 
               WHERE id = $1, id`)
	if err != nil {
		logger.Error.Printf("repository.DeleteSaleItem: failed to delete item ID %d: %v", id, err)
	}
	return translateError(err)
}
