package auth

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"main/utils"
	"net/http"
)

func Router() *chi.Mux {
	r := chi.NewMux()

	r.Post("/login", loginHandler)

	return r
}

// loginHandler - Returns all the available APIs
// @Summary Login
// @Description login account.
// @Tags Auth
// @Accept  json
// @Produce  json
//
//	@Param account	body LoginItem true "Login"
//
// @Success 200 {object} utils.Response
// @Router /auth/login [post]
func loginHandler(w http.ResponseWriter, req *http.Request) {
	var account LoginItem
	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ctx := req.Context()

	item, err := findItem(ctx, account.Username)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	err = utils.VerifyPassword(item.Password, account.Password)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, ErrANotAuthenticated)
		return
	}

	data := map[string]interface{}{"identity": item.Username}
	token := utils.GenerateJWT(data)

	var resp struct {
		AccessToken string `json:"access_token"`
	}
	resp.AccessToken = token

	utils.WriteResponse(resp, w, "Success")
}
