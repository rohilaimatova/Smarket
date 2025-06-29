package models

import "time"

type Category struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	AddedBy   int        `json:"added_by" db:"added_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name"`
}
