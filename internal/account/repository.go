//go:build !fake

package account

import (
	"context"
	"github.com/jackc/pgx/v5"
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
