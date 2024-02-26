package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"main/configs"
	"main/internal/account"
	"main/internal/auth"
	"main/middlewares"
	"main/utils"
	"net/http"
	"net/http/pprof"
)

func main() {
	configs.LoadConfig(".")
	ctx := context.Background()

	pool, err := pgxpool.NewWithConfig(ctx, configs.PgConfig())

	if err != nil {
		log.Error().Err(err).Msg("unable to connect to database")
	}

	account.SetPool(pool)
	auth.SetPool(pool)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.Group(func(r chi.Router) {
		r.Mount("/auth", auth.Router())
	})

	r.Group(func(r chi.Router) {
		//r.Mount("/", utils.RegisterAuth())
		r.Use(jwtauth.Verifier(utils.TokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(middlewares.Authenticator(utils.TokenAuth))
		r.HandleFunc("/pprof/profile", pprof.Profile)
		r.Mount("/account", account.Router())
	})

	http.ListenAndServe("0.0.0.0:8000", r)
}
