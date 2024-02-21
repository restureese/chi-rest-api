//go:build fake

package account

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func saveItem(ctx context.Context, tx pgx.Tx, item AccountItem) error {

	log.Debug().Msg("Fake save item")

	var found bool

	for i, v := range fake_items {
		if item.Id == v.Id {
			fake_items[i] = item
			return nil
		}
	}

	if !found {
		fake_items = append(fake_items, item)
	}
	return nil

}
