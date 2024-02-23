package utils

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog/log"
	"main/configs"
	"time"
)

var tokenAuth = jwtauth.New("HS256", []byte(configs.Config.JwtSecret), nil)

func GenerateJWT(data map[string]interface{}) string {
	jwtauth.SetExpiryIn(data, time.Duration(configs.Config.JwtDuration)*time.Hour)
	_, token, err := tokenAuth.Encode(data)
	if err != nil {
		log.Error().Err(err).Msg("failed encode jwt")
	}
	return token
}
