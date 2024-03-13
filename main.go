package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"main/cmd"
	"main/configs"
	_ "main/docs"
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
	r := cmd.NewApplication(pool)

	http.ListenAndServe("0.0.0.0:8000", r)
}
