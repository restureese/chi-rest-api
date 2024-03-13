package tests

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"main/cmd"
	"main/configs"
	"net/http/httptest"
	"testing"
)

func TestApplication(t *testing.T) {
	configs.LoadConfig(".")
	ctx := context.Background()

	pool, err := pgxpool.NewWithConfig(ctx, configs.PgConfig())
	if err != nil {
		log.Error().Err(err).Msg("unable to connect to database")
	}
	r := cmd.NewApplication(pool)
	ts := httptest.NewServer(r)
	defer ts.Close()

	if _, body := testRequest(t, ts, "GET", "/", nil); body != "404 page not found\n" {
		t.Fatalf(body)
	}
}
