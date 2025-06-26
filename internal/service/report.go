package service

import (
	"Smarket/internal/repository"
	"Smarket/models"
	"Smarket/pkg/errs"
	"Smarket/pkg/logger"
	"errors"
)

func GetSalesReport(fromDate, toDate string) ([]models.CashierSalesReport, error) {
	logger.Info.Printf("GetSalesReport: fetching report from %s to %s", fromDate, toDate)

	report, err := repository.GetSalesReport(fromDate, toDate)
	if err != nil {
		logger.Error.Printf("GetSalesReport: failed to fetch report from %s to %s: %v", fromDate, toDate, err)
		return nil, errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("GetSalesReport: successfully fetched report from %s to %s", fromDate, toDate)
	return report, nil
}
