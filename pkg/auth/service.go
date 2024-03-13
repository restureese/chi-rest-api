package auth

import (
	"context"
)

func findItem(ctx context.Context, username string) (item AccountItem, err error) {
	tx, err := pool.Begin(ctx)

	if err != nil {
		return
	}

	item, err = findItemByUsername(ctx, tx, username)

	if err != nil {
		return AccountItem{}, err
	}

	err = tx.Commit(ctx)
	return
}
