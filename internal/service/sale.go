package service

import (
	"Smarket/internal/repository"
	"Smarket/models"
	"Smarket/pkg/errs"
	"Smarket/pkg/logger"
	"database/sql"
	"errors"
)

func GetAllSales() ([]models.Sale, error) {
	sales, err := repository.GetAllSales()
	if err != nil {
		logger.Error.Printf("GetAllSales: failed to fetch sales: %v", err)
		return nil, errors.Join(errs.ErrInternal, err)
	}

	if len(sales) == 0 {
		logger.Warn.Println("GetAllSales: no sales found")
		return nil, errs.ErrNotFound
	}

	logger.Info.Printf("GetAllSales: fetched %d sales", len(sales))
	return sales, nil
}

func GetSaleByID(id int) (models.Sale, error) {
	sale, err := repository.GetSaleByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			logger.Warn.Printf("GetSaleByID: sale ID %d not found", id)
			return models.Sale{}, errs.ErrNotFound
		}
		logger.Error.Printf("GetSaleByID: failed to fetch sale ID %d: %v", id, err)
		return models.Sale{}, errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("GetSaleByID: fetched sale ID %d", id)
	return sale, nil
}

func CreateSale(sale models.Sale) error {
	if sale.UserID == 0 || sale.TotalSum <= 0 {
		logger.Warn.Printf("CreateSale: invalid input: %+v", sale)
		return errs.ErrInvalidValue
	}

	if err := repository.CreateSale(sale); err != nil {
		logger.Error.Printf("CreateSale: failed to create sale: %+v, error: %v", sale, err)
		return errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("CreateSale: sale created successfully: %+v", sale)
	return nil
}

func UpdateSale(id int, sale models.Sale) (models.Sale, error) {
	updated, err := repository.UpdateSale(id, sale)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			logger.Warn.Printf("UpdateSale: sale ID %d not found", id)
			return models.Sale{}, errs.ErrNotFound
		}

		logger.Error.Printf("UpdateSale: failed to update sale ID %d: %v", id, err)
		return models.Sale{}, errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("UpdateSale: sale ID %d updated successfully", id)
	return updated, nil
}

func DeleteSale(id int) error {
	if err := repository.DeleteSale(id); err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			logger.Warn.Printf("DeleteSale: sale ID %d not found", id)
			return errs.ErrNotFound
		}

		logger.Error.Printf("DeleteSale: failed to delete sale ID %d: %v", id, err)
		return errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("DeleteSale: sale ID %d deleted successfully", id)
	return nil
}

func GetSaleReceipt(saleID int) (models.Receipt, error) {
	receipt, err := repository.GetSaleReceipt(saleID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) || errors.Is(err, sql.ErrNoRows) {
			logger.Warn.Printf("GetSaleReceipt: receipt for sale ID %d not found", saleID)
			return models.Receipt{}, errs.ErrNotFound
		}

		logger.Error.Printf("GetSaleReceipt: error fetching receipt for sale ID %d: %v", saleID, err)
		return models.Receipt{}, errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("GetSaleReceipt: fetched receipt for sale ID %d", saleID)
	return receipt, nil
}
