package api_handlers

import (
	"Smarket/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetSalesReport godoc
// @Summary Получить отчёт по продажам за период
// @Description Возвращает отчёт по продажам между датами from и to (в формате YYYY-MM-DD)
// @Tags reports
// @Produce json
// @Param from query string true "Дата начала периода (YYYY-MM-DD)"
// @Param to query string true "Дата конца периода (YYYY-MM-DD)"
// @Success 200 {object} models.CashierSalesReport
// @Failure 400 {object} models.ErrorResponse "Параметры запроса from и to обязательны"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера при формировании отчёта"
// @Router /reports/sales [get]
func GetSalesReport(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")

	if from == "" || to == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "from and to query parameters are required"})
		return
	}

	report, err := service.GetSalesReport(from, to)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not generate report", err)
		return
	}

	c.JSON(http.StatusOK, report)
}
