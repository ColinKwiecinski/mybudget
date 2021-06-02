package auth

import (
	"mybudget/src/backend/sessions"
	"mybudget/src/backend/users"
)

type HandlerContext struct {
	SigningKey string
	Sessions   sessions.Store
	Users      users.Store
	DB 				 users.MysqlStore
}
