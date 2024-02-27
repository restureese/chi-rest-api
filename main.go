package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"main/configs"
	_ "main/docs"
	"main/internal/account"
	"main/internal/auth"
	"main/middlewares"
	"main/utils"
	"net/http"
)

// @title Swagger Example API
// @version 2.0
// @description This is a sample rest api server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 0.0.0.0:8000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

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
	r.Mount("/debug", middleware.Profiler())
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://0.0.0.0:8000/docs/doc.json"),
	))

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(middlewares.Authenticator(utils.TokenAuth))
		r.Mount("/accounts", account.Router())
	})

	r.Group(func(r chi.Router) {
		r.Mount("/auth", auth.Router())
	})

	http.ListenAndServe("0.0.0.0:8000", r)
}
