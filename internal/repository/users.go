package repository

import (
	"Smarket/internal/db"
	"Smarket/models"
	"Smarket/pkg/errs"
	"Smarket/pkg/logger"
	"database/sql"
	"errors"
)

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Get(&user, `SELECT id, 
					   name, 
					   username, 
					   created_at
				FROM users 
				WHERE deleted_at IS NULL 
				  AND username = $1
				  AND password_hash = $2`, username, password)
	if err != nil {
		logger.Error.
			Printf("[repository] GetUserByUsernameAndPassword(): error duriing getting from database: %s\n", err.Error())
		return models.User{}, translateError(err)
	}

	return user, nil
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	query := `SELECT id, username, password_hash FROM users WHERE username = $1`
	err := db.GetDBConn().Get(&user, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error.Printf("repository.GetUserByUsername(): failed to fetch from database: %s\n", err.Error())
			return user, errs.ErrNotFound
		}

		logger.Error.Printf("repository.GetUserByUsername(): failed to fetch from database: %s\n", err.Error())
		return user, err
	}
	return user, nil
}

func CreateUser(user models.UserSignUp) error {
	_, err := db.GetDBConn().Exec(`
			INSERT INTO users (name, username, password_hash)
			VALUES ($1, $2, $3)`, user.Name, user.Username, user.Password)
	if err != nil {
		logger.Error.Printf("repository.CreateUser(): failed to insert into users: %s\n", err.Error())

		return translateError(err)
	}

	return nil
}
