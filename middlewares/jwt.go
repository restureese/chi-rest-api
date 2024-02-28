package middlewares

import (
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"main/utils"
	"net/http"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
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
			if token == nil || jwt.Validate(token) != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			// Token is authenticated, pass it through
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}
