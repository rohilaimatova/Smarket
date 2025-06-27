package service

import (
	"Smarket/internal/repository"
	"Smarket/models"
	"Smarket/pkg/errs"
	"Smarket/pkg/logger"
	"errors"
)

func GetAllSaleItems() ([]models.SaleItem, error) {
	saleItems, err := repository.GetAllSaleItems()
	if err != nil {
		logger.Error.Printf("GetAllSaleItems: failed to fetch sale items: %v", err)
		return nil, errors.Join(errs.ErrInternal, err)
	}

	if len(saleItems) == 0 {
		logger.Warn.Println("GetAllSaleItems: no sale items found")
		return nil, errs.ErrNotFound
	}

	logger.Info.Printf("GetAllSaleItems: fetched %d sale items", len(saleItems))
	return saleItems, nil
}

func GetSaleItemByID(id int) (models.SaleItem, error) {
	item, err := repository.GetSaleItemByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			logger.Warn.Printf("GetSaleItemByID: item with ID %d not found", id)
			return models.SaleItem{}, errs.ErrNotFound
		}

		logger.Error.Printf("GetSaleItemByID: failed to fetch item ID %d: %v", id, err)
		return models.SaleItem{}, errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("GetSaleItemByID: fetched item ID %d", id)
	return item, nil
}

//func CreateSaleItem(item models.SaleItem) error {
//	if item.ProductID == 0 || item.SaleID == 0 || item.Quantity <= 0 {
//		logger.Warn.Printf("CreateSaleItem: invalid input: %+v", item)
//		return errs.ErrInvalidValue
//	}
//
//	product, err := repository.GetProductByID(item.ProductID)
//	if err != nil {
//		logger.Error.Printf("CreateSaleItem: product not found (product_id=%d): %v", item.ProductID, err)
//		return errors.Join(errs.ErrInvalidValue, err)
//	}
//
//	item.Price = product.Price
//
//	err = repository.CreateSaleItem(item)
//	if err != nil {
//		logger.Error.Printf("CreateSaleItem: failed to create sale item %+v: %v", item, err)
//		return errors.Join(errs.ErrInternal, err)
//	}
//
//	logger.Info.Printf("CreateSaleItem: sale item created successfully: %+v", item)
//	return nil
//}

func UpdateSaleItem(id int, item models.SaleItem) (models.SaleItem, error) {
	updatedItem, err := repository.UpdateSaleItem(id, item)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			logger.Warn.Printf("UpdateSaleItem: item ID %d not found", id)
			return models.SaleItem{}, errs.ErrNotFound
		}

		logger.Error.Printf("UpdateSaleItem: failed to update item ID %d: %v", id, err)
		return models.SaleItem{}, errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("UpdateSaleItem: item ID %d updated successfully", id)
	return updatedItem, nil
}

func DeleteSaleItem(id int) error {
	err := repository.DeleteSaleItem(id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			logger.Warn.Printf("DeleteSaleItem: item ID %d not found", id)
			return errs.ErrNotFound
		}

		logger.Error.Printf("DeleteSaleItem: failed to delete item ID %d: %v", id, err)
		return errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("DeleteSaleItem: item ID %d deleted successfully", id)
	return nil
}
