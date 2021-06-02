package auth

import (
	"mybudget/src/backend/sessions"
)

type HandlerContext struct {
	SigningKey string
	Sessions   sessions.Store
	Users      Store
	DB         MysqlStore
}
