//go:build fake

package auth

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func findItemByUsername(ctx context.Context, tx pgx.Tx, username string) (AccountItem, error) {
	q := `SELECT id, title, created_at FROM accounts WHERE username = $1`

	row := tx.QueryRow(ctx, q, username)

	var item AccountItem
	if err := row.Scan(&item.Id, &item.Username, &item.Password, &item.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			log.Debug().Err(err).Msg("can't find any item")
			return AccountItem{}, ErrANotAuthenticated
		}
		return AccountItem{}, err
	}

	return item, nil
}
