package utils

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog/log"
	"main/configs"
	"time"
)

var TokenAuth = jwtauth.New("HS256", []byte(configs.Config.JwtSecret), nil)

func GenerateJWT(data map[string]interface{}) string {
	jwtauth.SetExpiryIn(data, time.Duration(configs.Config.JwtDuration)*time.Hour)
	_, token, err := TokenAuth.Encode(data)
	if err != nil {
		log.Error().Err(err).Msg("failed encode jwt")
	}
	return token
}

func RegisterAuth() *chi.Mux {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(TokenAuth))
	r.Use(jwtauth.Authenticator(TokenAuth))
	return r
}
