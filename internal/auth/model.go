package auth

import (
	"github.com/oklog/ulid/v2"
	"gopkg.in/guregu/null.v4"
	"time"
)

type LoginItem struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AccountItem struct {
	Id        ulid.ULID
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt null.Time
	DeletedAt null.Time
}
