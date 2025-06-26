package api_handlers

import (
	"Smarket/internal/service"
	"Smarket/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GetAllCategories godoc
// @Summary Получить все категории
// @Description Возвращает список всех категорий
// @Tags categories
// @Produce json
// @Success 200 {array} models.Category
// @Failure 500 {object} models.ErrorResponse
// @Router /categories [get]
func GetAllCategories(c *gin.Context) {
	categories, err := service.GetAllCategories()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to fetch categories", err)
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID godoc
// @Summary Получить категорию по ID
// @Description Возвращает категорию по ID
// @Tags categories
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} models.Category
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /categories/{id} [get]
func GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid Category ID", err)
		return
	}

	category, err := service.GetCategoryByID(id)
	if err != nil {
		respondWithError(c, http.StatusNotFound, "Category not found", err)
		return
	}

	c.JSON(http.StatusOK, category)
}

// CreateCategory godoc
// @Summary Создать новую категорию
// @Description Принимает JSON и создаёт категорию
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Категория"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid JSON body", err)
		return
	}

	if err := service.CreateCategory(input); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to create category", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})
}

// UpdateCategory godoc
// @Summary Обновить категорию
// @Description Обновляет категорию по ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "ID категории"
// @Param category body models.Category true "Обновлённая категория"
// @Success 200 {object} models.Category
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid category ID", err)
		return
	}

	var updateCategory models.Category
	if err := c.ShouldBindJSON(&updateCategory); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid JSON body", err)
		return
	}

	result, err := service.UpdateCategory(id, updateCategory)
	if err != nil {
		respondWithError(c, http.StatusNotFound, "Failed to update category", err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteCategory godoc
// @Summary Удалить категорию
// @Description Удаляет категорию по ID. Если используется — выдаёт ошибку.
// @Tags categories
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 409 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid category ID", err)
		return
	}

	err = service.DeleteCategory(id)
	if err != nil {
		if strings.Contains(err.Error(), "used by one or more products") {
			respondWithError(c, http.StatusConflict, "Cannot delete category: it is used by products", err) // 409
			return
		}

		if strings.Contains(err.Error(), "not found") {
			respondWithError(c, http.StatusNotFound, "Category not found", err)
			return
		}

		respondWithError(c, http.StatusInternalServerError, "Failed to delete category", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
