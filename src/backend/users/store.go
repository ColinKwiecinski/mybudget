package users

import (
	"errors"
	"mybudget/src/backend/auth"

)

var ErrUserNotFound = errors.New("User Not Found")

type Store interface {
	GetByID(id int64) (*auth.User, error)

	GetByEmail(email string) (*auth.User, error)

	GetByContactNum(contactNum string) (*auth.User, error)

	Insert(user *auth.User) (*auth.User, error)

	Update(id int64, updates *auth.Updates) (*auth.User, error)

	Delete(id int64) error

	InsertTransaction(t *auth.Transaction) error

	DeleteTransaction(id int64) error

	GetTransactions(selector string, value string) (*[]auth.Transaction, error)
}