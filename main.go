package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"main/configs"
	"main/internal/account"
	"main/utils"
	"net/http"
	"net/http/pprof"
)

func main() {
	configs.LoadConfig(".")
	ctx := context.Background()
	data := map[string]interface{}{"identity": "restu"}
	token := utils.GenerateJWT(data)
	fmt.Println(token)

	pool, err := pgxpool.NewWithConfig(ctx, configs.PgConfig())

	if err != nil {
		log.Error().Err(err).Msg("unable to connect to database")
	}

	account.SetPool(pool)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.HandleFunc("/pprof/profile", pprof.Profile)
	r.Mount("/account", account.Router())
	http.ListenAndServe("0.0.0.0:8000", r)
}
