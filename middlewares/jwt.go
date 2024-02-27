package middlewares

import (
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"main/utils"
	"net/http"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrExpired      = errors.New("token is expired")
	ErrNBFInvalid   = errors.New("token nbf validation failed")
	ErrIATInvalid   = errors.New("token iat validation failed")
	ErrNoTokenFound = errors.New("no token found")
	ErrAlgoInvalid  = errors.New("algorithm mismatch")
)

func Authenticator(ja *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			token, _, err := jwtauth.FromContext(r.Context())

			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, ErrUnauthorized)
				return
			}

			if token == nil {
				utils.WriteError(w, http.StatusUnauthorized, ErrUnauthorized)
				return
			}

			// Token is authenticated, pass it through
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}
