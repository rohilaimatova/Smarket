package repository

import (
	"Smarket/internal/db"
	"Smarket/models"
	"Smarket/pkg/logger"
)

func GetAllProducts() ([]models.Product, error) {
	logger.Info.Println("repository.GetAllProducts: fetching all products")

	var products []models.Product
	query := `
		SELECT id, name, price, quantity, added_by, category_id, created_at, updated_at
	FROM products
	WHERE deleted_at IS NULL`

	err := db.GetDBConn().Select(&products, query)
	if err != nil {
		logger.Error.Printf("repository.GetAllProducts: failed to fetch products: %v", err)
		return nil, translateError(err)
	}

	logger.Info.Printf("repository.GetAllProducts: fetched %d products", len(products))
	return products, nil
}

func GetProductByID(id int) (models.Product, error) {
	logger.Info.Printf("repository.GetProductByID: fetching product with ID %d", id)

	var product models.Product
	query := `
		SELECT id, name, price, quantity, added_by, category_id, created_at, updated_at
	FROM products
	WHERE id = $1 AND deleted_at IS NULL`

	err := db.GetDBConn().Get(&product, query, id)
	if err != nil {
		logger.Error.Printf("repository.GetProductByID: failed to fetch product ID %d: %v", id, err)
		return models.Product{}, translateError(err)
	}

	logger.Info.Printf("repository.GetProductByID: found product ID %d", id)
	return product, nil
}

func CreateProduct(product models.Product) error {
	logger.Info.Printf("repository.CreateProduct: creating product: %+v", product)

	query := `
		INSERT INTO products (name, price, quantity, added_by, category_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	var id int
	err := db.GetDBConn().Get(&id, query,
		product.Name,
		product.Price,
		product.Quantity,
		product.AddedBy,
		product.CategoryID,
	)
	if err != nil {
		logger.Error.Printf("repository.CreateProduct: failed to create product: %v", err)
		return translateError(err)
	}

	logger.Info.Printf("repository.CreateProduct: product created with ID %d", id)
	return nil
}

func UpdateProduct(id int, product models.Product) (models.Product, error) {
	logger.Info.Printf("repository.UpdateProduct: updating product ID %d with data: %+v", id, product)

	query := `
		UPDATE products
	SET name = $1,
		price = $2,
		quantity = $3,
		added_by = $4,
		category_id = $5,
		updated_at = NOW()
	WHERE id = $6
	RETURNING id, name, price, quantity, added_by, category_id, created_at, updated_at, deleted_at`

	var updated models.Product
	err := db.GetDBConn().QueryRow(query,
		product.Name,
		product.Price,
		product.Quantity,
		product.AddedBy,
		product.CategoryID,
		id,
	).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Price,
		&updated.Quantity,
		&updated.AddedBy,
		&updated.CategoryID,
		&updated.CreatedAt,
		&updated.UpdateAt,
		&updated.DeletedAt,
	)
	if err != nil {
		logger.Error.Printf("repository.UpdateProduct: failed to update product ID %d: %v", id, err)
		return models.Product{}, translateError(err)
	}

	logger.Info.Printf("repository.UpdateProduct: product ID %d updated successfully", id)
	return updated, nil
}

func DeleteProducts(id int) error {
	logger.Info.Printf("repository.DeleteProducts: deleting product ID %d", id)

	query := `DELETE FROM products WHERE id = $1`
	_, err := db.GetDBConn().Exec(query, id)
	if err != nil {
		logger.Error.Printf("repository.DeleteProducts: failed to delete product ID %d: %v", id, err)
	}

	return translateError(err)
}
