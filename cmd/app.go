package cmd

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"main/middlewares"
	"main/pkg/account"
	"main/pkg/auth"
	"main/utils"
)

func NewApplication(pool *pgxpool.Pool) chi.Router {
	account.SetPool(pool)
	auth.SetPool(pool)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Mount("/debug", middleware.Profiler())
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
	return r
}
