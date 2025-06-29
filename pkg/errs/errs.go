package errs

import (
	"errors"
)

var (
	ErrUnauthorized                = errors.New("unauthorized")
	ErrProductNotFound             = errors.New("no product found")
	ErrInternal                    = errors.New("internal error")
	ErrNoCategoriesFound           = errors.New("no categories found")
	ErrUserIDNotFound              = errors.New("user ID Not Found")
	ErrUserAlreadyExists           = errors.New(`user already exists`)
	ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password")
	ErrAccountNotFound             = errors.New("account not found")
	ErrUserNotFound                = errors.New("user not found")
	ErrInvalidOperationType        = errors.New("invalid operation type")
	ErrNotFound                    = errors.New("not found")
	ErrValidationFailed            = errors.New("validation failed")
	ErrSomethingWentWrong          = errors.New("something went wrong")
	ErrTaskNotFound                = errors.New("task not found")
	ErrInvalidValue                = errors.New("invalid value")
	ErrCategoryNotFound            = errors.New("category not found")
	ErrCategoryExists              = errors.New("category already exists")
)
