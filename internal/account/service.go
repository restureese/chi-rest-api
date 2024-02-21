package account

import (
	"context"
	"main/pkg/utils"
)

func createItem(ctx context.Context, username string, password string) (account AccountItem, err error) {
	hashPassword, err := utils.GeneratePassword(password)
	if err != nil {
		return
	}
	accountItem, err := NewAccountItem(username, hashPassword)

	if err != nil {
		return
	}

	tx, err := pool.Begin(ctx)

	if err != nil {
		return
	}

	err = saveItem(ctx, tx, accountItem)

	if err != nil {
		tx.Rollback(ctx)
		return
	}

	err = tx.Commit(ctx)

	if err != nil {
		return
	}

	return accountItem, nil
}
