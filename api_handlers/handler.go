package api_handlers

import (
	"Smarket/pkg/errs"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping godoc
// @Summary Проверка доступности сервера
// @Description Проверяет, работает ли сервер, и возвращает сообщение
// @Tags system
// @Produce json
// @Success 200 {object} map[string]string "Сервер работает"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is running",
	})

}

func respondWithError(c *gin.Context, status int, message string, err error) {
	c.JSON(status, gin.H{
		"error":   message,
		"details": err.Error(),
	})
}

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	} else if errors.Is(err, errs.ErrValidationFailed) ||
		errors.Is(err, errs.ErrInvalidOperationType) ||
		errors.Is(err, errs.ErrUserAlreadyExists) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errs.ErrAccountNotFound) ||
		errors.Is(err, errs.ErrUserNotFound) ||
		errors.Is(err, errs.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else if errors.Is(err, errs.ErrIncorrectUsernameOrPassword) ||
		errors.Is(err, errs.ErrUserIDNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("something went wrong: %s", err.Error()),
		})
	}
}
