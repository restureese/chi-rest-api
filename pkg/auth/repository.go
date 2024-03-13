package auth

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func saveItem(ctx context.Context, tx pgx.Tx, item AccountItem) error {
	q := `INSERT INTO accounts(id, username, password, created_at) VALUES ( $1, $2, $3, $4) 
			ON CONFLICT(id) DO UPDATE SET username=$2, password=$3, created_at=$4`

	_, err := tx.Exec(ctx, q, item.Id, item.Username, item.Password, item.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func findItemByUsername(ctx context.Context, tx pgx.Tx, username string) (AccountItem, error) {
	q := `SELECT id, username, password, created_at FROM accounts WHERE username = $1`

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
