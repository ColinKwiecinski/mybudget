package auth

import (
	"errors"
)

var ErrUserNotFound = errors.New("User Not Found")

type Store interface {
	GetByID(id int64) (*User, error)

	GetByEmail(email string) (*User, error)

	GetByContactNum(contactNum string) (*User, error)

	Insert(user *User) (*User, error)

	Update(id int64, updates *Updates) (*User, error)

	Delete(id int64) error

	InsertTransaction(t *Transaction) error

	DeleteTransaction(id int64) error

	GetTransactions(selector string, value string) (*[]Transaction, error)
}
