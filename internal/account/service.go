package account

import (
	"context"
	"fmt"
	"github.com/oklog/ulid/v2"
	"main/utils"
)

func listItems(ctx context.Context) (AccountList, error) {
	tx, err := pool.Begin(ctx)

	if err != nil {
		return AccountList{}, err
	}

	list, err := findAllItems(ctx, tx)

	if err != nil {
		return AccountList{}, err
	}

	tx.Commit(ctx)

	return list, nil
}

func findItem(ctx context.Context, id ulid.ULID) (item AccountItem, err error) {
	tx, err := pool.Begin(ctx)

	if err != nil {
		return
	}

	item, err = findItemById(ctx, tx, id)

	if err != nil {
		return AccountItem{}, err
	}

	err = tx.Commit(ctx)
	return
}

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

func makeItemUpdate(ctx context.Context, id ulid.ULID, username string, password string) (account AccountItem, err error) {
	hashPassword, err := utils.GeneratePassword(password)
	if err != nil {
		return
	}

	tx, err := pool.Begin(ctx)

	if err != nil {
		return
	}

	item, err := findItemById(ctx, tx, id)
	if err != nil {
		return item, err
	}

	err = item.UpdateItem(username, hashPassword)
	if err != nil {
		tx.Rollback(ctx)
		return
	}

	if err = saveItem(ctx, tx, item); err != nil {
		tx.Rollback(ctx)
		return item, err
	}

	err = tx.Commit(ctx)

	if err != nil {
		return
	}

	return item, nil
}

func makeItemDelete(ctx context.Context, id ulid.ULID) error {
	tx, err := pool.Begin(ctx)

	if err != nil {
		return err
	}

	item, err := findItemById(ctx, tx, id)
	if err != nil {
		return err
	}

	err = item.DeleteItem()
	fmt.Println(err)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err = saveItem(ctx, tx, item); err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}
