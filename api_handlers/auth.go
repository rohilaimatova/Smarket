package api_handlers

import (
	"Smarket/internal/service"
	"Smarket/models"
	"Smarket/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth SignUp godoc
// @Summary Регистрация нового пользователя
// @Description Создаёт нового пользователя на основе переданных данных (username, password и т.п.)
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.UserSignUp true "Данные нового пользователя"
// @Success 200 {object} map[string]string "Успешная регистрация"
// @Failure 400 {object} models.ErrorResponse "Неверный JSON"
// @Failure 500 {object} models.ErrorResponse "Ошибка при создании пользователя"
// @Router /auth/sign-up [post]
func SignUp(c *gin.Context) {
	var user models.UserSignUp

	if err := c.ShouldBindJSON(&user); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid JSON payload", err)
		return
	}

	if err := service.CreateUser(user); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

// Auth SignIn godoc
// @Summary Авторизация пользователя
// @Description Проверяет имя пользователя и пароль, возвращает JWT токен
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.UserSignIn true "Данные для входа"
// @Success 200 {object} map[string]string "JWT access token"
// @Failure 400 {object} models.ErrorResponse "Неверный JSON"
// @Failure 401 {object} models.ErrorResponse "Неверное имя пользователя или пароль"
// @Failure 500 {object} models.ErrorResponse "Ошибка генерации токена"
// @Router /auth/sign-in [post]
func SignIn(c *gin.Context) {
	var requestUser models.UserSignIn

	if err := c.ShouldBindJSON(&requestUser); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid credentials payload", err)
		return
	}

	user, err := service.GetUserByUsernameAndPassword(requestUser.Username, requestUser.PasswordHash)
	if err != nil {
		respondWithError(c, http.StatusUnauthorized, "Invalid username or password", err)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to generate access token", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
