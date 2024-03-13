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

type CreateAccountItem struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountList struct {
	Items []AccountItem `json:"items"`
	Count int           `json:"count"`
}

func (t AccountItem) IsDeleted() bool {
	return t.DeletedAt.IsZero() == false
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

func (t *AccountItem) UpdateItem(username string, password string) error {
	if err := validateTitle(username); err != nil {
		return err
	}
	t.Username = username
	t.Password = password
	t.UpdatedAt = null.TimeFrom(time.Now())
	return nil
}

func (t *AccountItem) DeleteItem() error {
	if t.IsDeleted() {
		return ErrIsDeleted
	}
	t.DeletedAt = null.TimeFrom(time.Now())
	return nil
}
