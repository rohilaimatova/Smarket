package api_handlers

import (
	"Smarket/internal/service"
	"Smarket/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllSaleItems godoc
// @Summary Получить все позиции продажи
// @Security BearerAuth
// @Description Возвращает список всех позиций продажи (SaleItems)
// @Tags sale-items
// @Produce json
// @Success 200 {array} models.SaleItem
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера при получении позиций продажи"
// @Router /api/sale-items [get]
func GetAllSaleItems(c *gin.Context) {
	sale, err := service.GetAllSaleItems()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not fetch sale items", err)
	}
	c.JSON(http.StatusOK, sale)
}

// GetSaleItemByID godoc
// @Summary Получить позицию продажи по ID
// @Security BearerAuth
// @Description Возвращает позицию продажи по идентификатору
// @Tags sale-items
// @Produce json
// @Param id path int true "ID позиции продажи"
// @Success 200 {object} models.SaleItem
// @Failure 400 {object} models.ErrorResponse "Неверный ID позиции продажи"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера при получении позиции продажи"
// @Router /api/sale-items/{id} [get]
func GetSaleItemByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid Sale item  ID", err)
		return
	}
	sale, err := service.GetSaleItemByID(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not fetch ale item", err)
		return
	}
	c.JSON(http.StatusOK, sale)
}

// UpdateSaleItem godoc
// @Summary Обновить позицию продажи
// @Security BearerAuth
// @Description Обновляет данные позиции продажи по ID
// @Tags sale-items
// @Accept json
// @Produce json
// @Param id path int true "ID позиции продажи"
// @Param saleItem body models.UpdateSaleItemRequest true "Обновленные данные позиции продажи"
// @Success 201 {object} models.SaleItem
// @Failure 400 {object} models.ErrorResponse "Неверный запрос или ID"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера при обновлении позиции продажи"
// @Router /api/sale-items/{id} [put]
func UpdateSaleItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid Sale item ID", err)
		return
	}
	var updateItem models.SaleItem
	if err := c.ShouldBindJSON(&updateItem); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}
	result, err := service.UpdateSaleItem(id, updateItem)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not update sale item", err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// DeleteSaleItem godoc
// @Summary Удалить позицию продажи
// @Security BearerAuth
// @Description Удаляет позицию продажи по ID
// @Tags sale-items
// @Produce json
// @Param id path int true "ID позиции продажи"
// @Success 200 {object} map[string]string "Позиция продажи успешно удалена"
// @Failure 400 {object} models.ErrorResponse "Неверный ID позиции продажи"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера при удалении позиции продажи"
// @Router /api/sale-items/{id} [delete]
func DeleteSaleItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid Sale item ID", err)
		return
	}
	if err := service.DeleteSaleItem(id); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not delete sale item", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "SaleItem deleted",
	})

}
