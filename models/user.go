package models

import "time"

type User struct {
	ID           int       `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"-" db:"password_hash" ` // НЕ возвращаем в JSON
	Name         string    `json:"name" db:"name"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at" db:"deleted_at"`
}

type UserSignUp struct {
	Username string `json:"username"`
	Password string `json:"password"` // Пользователь вводит обычный пароль
	Name     string `json:"name"`
}

type UserSignIn struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
}
