package account

import (
	"github.com/oklog/ulid/v2"
	"gopkg.in/guregu/null.v4"
	"time"
)

type AccountItem struct {
	Id        ulid.ULID
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt null.Time
	DeletedAt null.Time
}

func NewAccountItem(username string, password string) (AccountItem, error) {
	if err := validateTitle(username); err != nil {
		return AccountItem{}, err
	}

	item := AccountItem{
		Id:        ulid.Make(),
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}

	return item, nil
}
