package service

import (
	"Smarket/internal/repository"
	"Smarket/models"
	"Smarket/pkg/errs"
	"Smarket/pkg/logger"
	"Smarket/pkg/utils"
	"errors"
)

func CreateUser(u models.UserSignUp) error {
	if u.Username == "" || u.Password == "" {
		logger.Warn.Printf("CreateUser: invalid input: %+v", u)
		return errs.ErrInvalidValue
	}

	_, err := repository.GetUserByUsername(u.Username)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		logger.Error.Printf("CreateUser: failed to check username existence: %v", err)
		return errors.Join(errs.ErrInternal, err)
	} else if err == nil {
		logger.Warn.Printf("CreateUser: user %s already exists", u.Username)
		return errs.ErrUserAlreadyExists
	}

	u.Password = utils.GenerateHash(u.Password)

	if err := repository.CreateUser(u); err != nil {
		logger.Error.Printf("CreateUser: failed to create user %+v: %v", u, err)
		return errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("CreateUser: user %s created successfully", u.Username)
	return nil
}

func GetUserByUsernameAndPassword(username string, password string) (models.User, error) {
	if username == "" || password == "" {
		logger.Warn.Println("GetUserByUsernameAndPassword: empty credentials")
		return models.User{}, errs.ErrInvalidValue
	}

	hashedPassword := utils.GenerateHash(password)
	user, err := repository.GetUserByUsernameAndPassword(username, hashedPassword)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			logger.Warn.Printf("GetUserByUsernameAndPassword: invalid credentials for user %s", username)
			return models.User{}, errs.ErrIncorrectUsernameOrPassword
		}

		logger.Error.Printf("GetUserByUsernameAndPassword: internal error for user %s: %v", username, err)
		return models.User{}, errors.Join(errs.ErrInternal, err)
	}

	logger.Info.Printf("GetUserByUsernameAndPassword: user %s authenticated successfully", username)
	return user, nil
}
