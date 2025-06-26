package api_handlers

import (
	"Smarket/internal/service"
	"Smarket/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllSales godoc
// @Summary Получить все продажи
// @Description Возвращает список всех продаж
// @Tags sales
// @Produce json
// @Success 200 {array} models.Sale
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера при получении продаж"
// @Router /sales [get]
func GetAllSales(c *gin.Context) {
	sale, err := service.GetAllSales()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not fetch receipt", err)
	}
	c.JSON(http.StatusOK, sale)
}

// GetSaleByID godoc
// @Summary Получить продажу по ID
// @Description Возвращает продажу по идентификатору
// @Tags sales
// @Produce json
// @Param id path int true "ID продажи"
// @Success 200 {object} models.Sale
// @Failure 400 {object} models.ErrorResponse "Неверный ID продажи"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера при получении продажи"
// @Router /sales/{id} [get]
func GetSaleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid Sale ID", err)
		return
	}
	sale, err := service.GetSaleByID(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Can not get sale", err)
		return
	}
	c.JSON(http.StatusOK, sale)
}

// CreateSale godoc
// @Summary Создать продажу
// @Description Создает новую продажу
// @Tags sales
// @Accept json
// @Produce json
// @Param sale body models.Sale true "Данные продажи"
// @Success 200 {object} map[string]string "Успешное создание"
// @Failure 400 {object} models.ErrorResponse "Неверный запрос"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера при создании продажи"
// @Router /sales [post]
func CreateSale(c *gin.Context) {
	var newSale models.Sale
	if err := c.ShouldBindJSON(&newSale); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request payload", err)

		return
	}
	err := service.CreateSale(newSale)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not create sale", err)

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Sale created successfully",
	})
}

// UpdateSale godoc
// @Summary Обновить продажу
// @Description Обновляет данные продажи по ID
// @Tags sales
// @Accept json
// @Produce json
// @Param id path int true "ID продажи"
// @Param sale body models.Sale true "Обновленные данные продажи"
// @Success 200 {object} models.Sale
// @Failure 400 {object} models.ErrorResponse "Неверный запрос или ID"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера при обновлении"
// @Router /sales/{id} [put]
func UpdateSale(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid Category ID", err)
		return
	}
	var updateSale models.Sale
	if err := c.ShouldBindJSON(&updateSale); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}
	result, err := service.UpdateSale(id, updateSale)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not update sale", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteSale godoc
// @Summary Удалить продажу
// @Description Удаляет продажу по ID
// @Tags sales
// @Produce json
// @Param id path int true "ID продажи"
// @Success 200 {object} map[string]string "Удаление успешно"
// @Failure 400 {object} models.ErrorResponse "Неверный ID"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера при удалении"
// @Router /sales/{id} [delete]
func DeleteSale(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid Category ID", err)
		return
	}
	if err := service.DeleteSale(id); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not delete sale", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Sale deleted",
	})

}
