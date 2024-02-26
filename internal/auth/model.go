package auth

import (
	"github.com/oklog/ulid/v2"
	"gopkg.in/guregu/null.v4"
	"time"
)

type LoginItem struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountItem struct {
	Id        ulid.ULID
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt null.Time
	DeletedAt null.Time
}

func NewLoginItem(username string, password string) (LoginItem, error) {
	if err := validateTitle(username); err != nil {
		return LoginItem{}, err
	}

	item := LoginItem{
		Username: username,
		Password: password,
	}

	return item, nil
}
