package api_handlers

import (
	"Smarket/internal/service"
	"Smarket/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllProducts godoc
// @Summary Получить все продукты
// @Description Возвращает список всех продуктов
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} models.ErrorResponse
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
	products, err := service.GetAllProducts()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not fetch product", err)
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductByID godoc
// @Summary Получить продукт по ID
// @Description Возвращает продукт по его идентификатору
// @Tags products
// @Produce json
// @Param id path int true "ID продукта"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid product ID", err)
		return
	}
	category, err := service.GetProductByID(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Product not found", err)
		return
	}
	c.JSON(http.StatusOK, category)
}

// CreateProduct godoc
// @Summary Создать продукт
// @Description Принимает JSON и создаёт новый продукт
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Новый продукт"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid JSON payload", err)
		return
	}
	err := service.CreateProduct(newProduct)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not create product", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
	})
}

// UpdateProduct godoc
// @Summary Обновить продукт
// @Description Обновляет данные продукта по ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Param product body models.Product true "Обновлённые данные продукта"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid product ID", err)
		return
	}
	var updateProduct models.Product
	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid JSON payload", err)
		return
	}
	result, err := service.UpdateProduct(id, updateProduct)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Product update faild", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteProduct godoc
// @Summary Удалить продукт
// @Description Удаляет продукт по ID
// @Tags products
// @Produce json
// @Param id path int true "ID продукта"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid product ID", err)
		return
	}
	if err := service.DeleteProduct(id); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Product delete faild", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})

}
