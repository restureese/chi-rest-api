package account

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool               *pgxpool.Pool
	ErrAccountNotFound = errors.New("account not found")
	ErrIsDeleted       = errors.New("account not found")
)

func SetPool(newPool *pgxpool.Pool) error {

	if newPool == nil {
		return errors.New("cannot assign nil pool")
	}

	pool = newPool

	return nil
}
