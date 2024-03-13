package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"main/configs"
	_ "main/docs"
	"main/internal/account"
	"main/internal/auth"
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
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Mount("/debug", middleware.Profiler())
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://0.0.0.0:8000/docs/doc.json"),
	))

	//r

	r.Group(func(r chi.Router) {
		r.Mount("/auth", auth.Router())
		r.Mount("/accounts", account.Router())
	})

	http.ListenAndServe("0.0.0.0:8000", r)
}
