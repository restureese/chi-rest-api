package account

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/guregu/null.v4"
	"time"
)

var emptyList AccountList

func findAllItems(ctx context.Context, tx pgx.Tx) (AccountList, error) {
	var itemCount int

	row := tx.QueryRow(ctx, "SELECT COUNT(id) as cnt FROM accounts;")
	err := row.Scan(&itemCount)

	if err != nil {
		log.Warn().Err(err).Msg("cannot find a count in account list")
		return emptyList, err
	}

	if itemCount == 0 {
		return emptyList, nil
	}

	log.Debug().Int("count", itemCount).Msg("found account items")

	items := make([]AccountItem, itemCount)

	rows, err := tx.Query(ctx, "SELECT id, username, created_at, updated_at, deleted_at FROM accounts")

	if err != nil {
		return emptyList, err
	}

	defer rows.Close()

	var i int

	for i = range items {
		var id ulid.ULID
		var username string
		var createdAt time.Time
		var updatedAt null.Time
		var deletedAt null.Time

		if !rows.Next() {
			break
		}

		if err := rows.Scan(&id, &username, &createdAt, &updatedAt, &deletedAt); err != nil {
			log.Warn().Err(err).Msg("cannot scan an item")
			return emptyList, err
		}
		items[i] = AccountItem{
			Id: id, Username: username, CreatedAt: createdAt, UpdatedAt: updatedAt, DeletedAt: deletedAt,
		}
	}

	list := AccountList{
		Items: items,
		Count: itemCount,
	}

	return list, nil
}

func findItemById(ctx context.Context, tx pgx.Tx, id ulid.ULID) (AccountItem, error) {
	q := `SELECT id, username, password, created_at, updated_at, deleted_at FROM accounts WHERE id = $1`

	row := tx.QueryRow(ctx, q, id)

	var item AccountItem
	if err := row.Scan(&item.Id, &item.Username, &item.Password, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); err != nil {
		if err == pgx.ErrNoRows {
			log.Debug().Err(err).Msg("can't find any item")
			return AccountItem{}, ErrAccountNotFound
		}
		return AccountItem{}, err
	}

	return item, nil
}

func saveItem(ctx context.Context, tx pgx.Tx, item AccountItem) error {
	q := `INSERT INTO accounts(id, username, password, created_at) VALUES ( $1, $2, $3, $4) 
			ON CONFLICT(id) DO UPDATE SET username=$2, password=$3, created_at=$4, updated_at=$5, deleted_at=$6`

	_, err := tx.Exec(ctx, q, item.Id, item.Username, item.Password, item.CreatedAt, item.UpdatedAt, item.DeletedAt)

	if err != nil {
		return err
	}

	return nil
}
