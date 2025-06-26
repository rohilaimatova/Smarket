package repository

import (
	"Smarket/internal/db"
	"Smarket/models"
	"Smarket/pkg/logger"
)

func GetAll() ([]models.Category, error) {
	logger.Info.Println("repository.GetAll: fetching all categories")

	var categories []models.Category
	query := `
		SELECT id, name, added_by, created_at, updated_at
	FROM category_products`

	err := db.GetDBConn().Select(&categories, query)
	if err != nil {
		logger.Error.Printf("repository.GetAll: failed to fetch categories: %v", err)
		return nil, translateError(err)
	}

	logger.Info.Printf("repository.GetAll: fetched %d categories", len(categories))
	return categories, nil
}

func GetByID(id int) (models.Category, error) {
	logger.Info.Printf("repository.GetByID: fetching category with ID %d", id)

	var category models.Category
	query := `
		SELECT id, name, added_by, created_at, updated_at
	FROM category_products
	WHERE id = $1`

	err := db.GetDBConn().Get(&category, query, id)
	if err != nil {
		logger.Error.Printf("repository.GetByID: failed to fetch category ID %d: %v", id, err)
		return models.Category{}, translateError(err)
	}

	logger.Info.Printf("repository.GetByID: found category ID %d", id)
	return category, nil
}

func Create(cat models.Category) error {
	logger.Info.Printf("repository.Create: creating category: %+v", cat)

	query := `
		INSERT INTO category_products(name, added_by)
	VALUES ($1, $2)
	RETURNING id`

	var id int
	err := db.GetDBConn().Get(&id, query, cat.Name, cat.AddedBy)
	if err != nil {
		logger.Error.Printf("repository.Create: failed to create category: %v", err)
		return translateError(err)
	}

	logger.Info.Printf("repository.Create: category created with ID %d", id)
	return nil
}

func Update(id int, cat models.Category) (models.Category, error) {
	logger.Info.Printf("repository.Update: updating category ID %d with data: %+v", id, cat)

	query := `
		UPDATE category_products
	SET name = $1,
		added_by = $2,
		updated_at = NOW()
	WHERE id = $3
	RETURNING id, name, added_by, created_at, updated_at, deleted_at`

	var updated models.Category
	err := db.GetDBConn().QueryRow(query, cat.Name, cat.AddedBy, id).Scan(
		&updated.ID,
		&updated.Name,
		&updated.AddedBy,
		&updated.CreatedAt,
		&updated.UpdatedAt,
		&updated.DeletedAt,
	)
	if err != nil {
		logger.Error.Printf("repository.Update: failed to update category ID %d: %v", id, err)
		return models.Category{}, translateError(err)
	}

	logger.Info.Printf("repository.Update: category ID %d updated successfully", id)
	return updated, nil
}

func Delete(id int) (int64, error) {
	logger.Info.Printf("repository.Delete: deleting category ID %d", id)

	query := `DELETE FROM category_products WHERE id = $1`
	res, err := db.GetDBConn().Exec(query, id)
	if err != nil {
		logger.Error.Printf("repository.Delete: failed to delete category ID %d: %v", id, err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Error.Printf("repository.Delete: failed to get affected rows for ID %d: %v", id, err)
		return 0, err
	}

	logger.Info.Printf("repository.Delete: category ID %d deleted, rows affected: %d", id, rowsAffected)
	return rowsAffected, nil
}
