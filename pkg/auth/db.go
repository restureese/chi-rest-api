package auth

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool                 *pgxpool.Pool
	ErrANotAuthenticated = errors.New("username or password is wrong")
)

func SetPool(newPool *pgxpool.Pool) error {

	if newPool == nil {
		return errors.New("cannot assign nil pool")
	}

	pool = newPool

	return nil
}
