package auth

import (
	"mybudget.com/src/backend/sessions"
)

type HandlerContext struct {
	SigningKey string
	Sessions   sessions.Store
	Users      Store
	DB         MysqlStore
}
