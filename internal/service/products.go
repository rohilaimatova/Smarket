package service

import (
	"Smarket/internal/repository"
	"Smarket/models"
	"Smarket/pkg/errs"
	"Smarket/pkg/logger"
	"errors"
)

func GetAllProducts() ([]models.Product, error) {
	products, err := repository.GetAllProducts()
	if err != nil {
		logger.Error.Printf("[service] GetAllProducts(): failed to fetch all products: %v\n", err)
		return nil, errs.ErrProductNotFound
	}

	logger.Info.Printf("[service] GetAllProducts(): %d products fetched\n", len(products))
	return products, nil
}

func GetProductByID(id int) (models.Product, error) {
	product, err := repository.GetProductByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrProductNotFound) {
			logger.Warn.Printf("[service] GetProductByID(): product with ID %d not found\n", id)
			return models.Product{}, errs.ErrProductNotFound
		}
		logger.Error.Printf("[service] GetProductByID(): error fetching product ID %d: %v\n", id, err)
		return models.Product{}, errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("[service] GetProductByID(): product ID %d fetched\n", id)
	return product, nil
}

func CreateProduct(product models.Product) error {
	if product.Name == "" {
		logger.Warn.Println("[service] CreateProduct(): empty product name")
		return errs.ErrInvalidValue
	}
	if product.AddedBy == 0 {
		logger.Warn.Println("[service] CreateCategory(): user ID is missing")
		return errs.ErrUnauthorized
	}

	if err := repository.CreateProduct(product); err != nil {
		logger.Error.Printf("[service] CreateProduct(): failed to create product %+v: %v\n", product, err)
		return errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("[service] CreateProduct(): product created successfully: %+v\n", product)
	return nil
}

func UpdateProduct(id int, product models.Product) (models.Product, error) {
	updatedProduct, err := repository.UpdateProduct(id, product)
	if err != nil {
		if errors.Is(err, errs.ErrProductNotFound) {
			logger.Warn.Printf("[service] UpdateProduct(): product ID %d not found\n", id)
			return models.Product{}, errs.ErrProductNotFound
		}
		logger.Error.Printf("[service] UpdateProduct(): failed to update product ID %d: %v\n", id, err)
		return models.Product{}, errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("[service] UpdateProduct(): product ID %d updated successfully\n", id)
	return updatedProduct, nil
}

func DeleteProduct(id int) error {
	if err := repository.DeleteProducts(id); err != nil {
		if errors.Is(err, errs.ErrProductNotFound) {
			logger.Warn.Printf("[service] DeleteProduct(): product ID %d not found\n", id)
			return errs.ErrProductNotFound
		}
		logger.Error.Printf("[service] DeleteProduct(): failed to delete product ID %d: %v\n", id, err)
		return errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("[service] DeleteProduct(): product ID %d deleted successfully\n", id)
	return nil
}
