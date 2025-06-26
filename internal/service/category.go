package service

import (
	"Smarket/internal/repository"
	"Smarket/models"
	"Smarket/pkg/errs"
	"Smarket/pkg/logger"
	"errors"
	"strings"
)

func GetAllCategories() ([]models.Category, error) {
	categories, err := repository.GetAll()
	if err != nil {
		logger.Error.Printf("[service] GetAllCategories(): failed to fetch categories: %v\n", err)
		return nil, errors.Join(errs.ErrInternal, err)
	}

	if len(categories) == 0 {
		logger.Warn.Println("[service] GetAllCategories(): no categories found")
		return nil, errs.ErrNoCategoriesFound
	}

	logger.Info.Printf("[service] GetAllCategories(): %d categories fetched\n", len(categories))
	return categories, nil
}

func GetCategoryByID(id int) (models.Category, error) {
	category, err := repository.GetByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrCategoryNotFound) {
			logger.Warn.Printf("[service] GetCategoryByID(): category with ID %d not found\n", id)
			return models.Category{}, errs.ErrCategoryNotFound
		}
		logger.Error.Printf("[service] GetCategoryByID(): error fetching category ID %d: %v\n", id, err)
		return models.Category{}, errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("[service] GetCategoryByID(): category ID %d fetched successfully\n", id)
	return category, nil
}

func CreateCategory(category models.Category) error {
	if category.Name == "" {
		logger.Warn.Println("[service] CreateCategory(): empty category name provided")
		return errs.ErrInvalidValue
	}

	if err := repository.Create(category); err != nil {
		logger.Error.Printf("[service] CreateCategory(): failed to create category %+v: %v\n", category, err)
		return errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("[service] CreateCategory(): category created successfully: %+v\n", category)
	return nil
}

func UpdateCategory(id int, category models.Category) (models.Category, error) {
	updated, err := repository.Update(id, category)
	if err != nil {
		if errors.Is(err, errs.ErrCategoryNotFound) {
			logger.Warn.Printf("[service] UpdateCategory(): category ID %d not found\n", id)
			return models.Category{}, errs.ErrCategoryNotFound
		}
		logger.Error.Printf("[service] UpdateCategory(): failed to update category ID %d: %v\n", id, err)
		return models.Category{}, errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("[service] UpdateCategory(): category ID %d updated successfully\n", id)
	return updated, nil
}

func DeleteCategory(id int) error {
	rowsAffected, err := repository.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			logger.Warn.Printf("[service] DeleteCategory(): category ID %d is used by products\n", id)
			return errors.New("cannot delete category: it is used by one or more products")
		}
		logger.Error.Printf("[service] DeleteCategory(): failed to delete category ID %d: %v\n", id, err)
		return errors.New("internal error")
	}

	if rowsAffected == 0 {
		logger.Warn.Printf("[service] DeleteCategory(): category ID %d not found\n", id)
		return errors.New("category not found")
	}

	logger.Info.Printf("[service] DeleteCategory(): category ID %d deleted\n", id)
	return nil
}
