package api_handlers

import (
	"Smarket/internal/service"
	"Smarket/pkg/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetReceipt godoc
// @Summary Получить чек по ID продажи
// @Security BearerAuth
// @Description Возвращает информацию о чеке продажи по идентификатору
// @Tags receipts
// @Produce json
// @Param id path int true "ID продажи"
// @Success 200 {object} models.Receipt
// @Failure 400 {object} models.ErrorResponse "Неверный ID продажи"
// @Failure 404 {object} models.ErrorResponse "Чек не найден"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера при получении чека"
// @Router /api/sales/{id}/receipt [get]
func GetReceipt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid sale ID", err)
		return
	}

	receipt, err := service.GetSaleReceipt(id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			respondWithError(c, http.StatusNotFound, "Receipt not found", err)
			return
		}

		respondWithError(c, http.StatusInternalServerError, "Could not fetch receipt", err)
		return
	}

	c.JSON(http.StatusOK, receipt)
}
