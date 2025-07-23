package api_handlers

import (
	"Smarket/internal/service"
	"Smarket/models"
	"Smarket/pkg/smRedis"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetAllCategories godoc
// @Summary Получить все категории
// @Security BearerAuth
// @Description Возвращает список всех категорий
// @Tags categories
// @Produce json
// @Success 200 {array} models.Category
// @Failure 500 {object} models.ErrorResponse
// @Router /api/categories [get]
func GetAllCategories(c *gin.Context) {
	cacheKey := "categories:all"

	// Пробуем взять из Redis
	cached, err := smRedis.Rdb.Get(smRedis.Ctx, cacheKey).Result()
	if err == nil {
		var categories []models.Category
		if err := json.Unmarshal([]byte(cached), &categories); err == nil {
			c.JSON(http.StatusOK, categories)
			return
		}
	}

	// Если в Redis нет — берём из базы
	categories, err := service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch categories",
		})
		return
	}

	// Кладём результат в Redis
	data, _ := json.Marshal(categories)
	smRedis.Rdb.Set(smRedis.Ctx, cacheKey, data, 5*time.Minute)

	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID godoc
// @Summary Получить категорию по ID
// @Security BearerAuth
// @Description Возвращает категорию по ID
// @Tags categories
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} models.Category
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/categories/{id} [get]
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
// @Security BearerAuth
// @Description Принимает JSON и создаёт категорию
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.CreateCategoryRequest true "Категория"
// @Success 201 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/categories [post]
func CreateCategory(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		respondWithError(c, http.StatusBadRequest, "User ID not found in context", nil)
		return
	}
	userID, ok := userIDInterface.(int)
	if !ok {
		respondWithError(c, http.StatusBadRequest, "invalid user ID type", nil)
		return
	}
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid JSON body", err)
		return
	}
	input.AddedBy = userID

	if err := service.CreateCategory(input); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to create category", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

// UpdateCategory godoc
// @Summary Обновить категорию
// @Security BearerAuth
// @Description Обновляет категорию по ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "ID категории"
// @Param category body models.UpdateCategoryRequest true "Обновлённая категория"
// @Success 200 {object} models.Category
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		respondWithError(c, http.StatusUnauthorized, "User ID not found", nil)
		return
	}
	userID, ok := userIDInterface.(int)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Invalid user ID type", nil)
		return
	}

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
	updateCategory.AddedBy = userID

	result, err := service.UpdateCategory(id, updateCategory)
	if err != nil {
		respondWithError(c, http.StatusNotFound, "Failed to update category", err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteCategory godoc
// @Summary Удалить категорию
// @Security BearerAuth
// @Description Удаляет категорию по ID. Если используется — выдаёт ошибку.
// @Tags categories
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 409 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/categories/{id} [delete]
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
