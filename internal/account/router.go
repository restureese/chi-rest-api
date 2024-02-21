package account

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"main/pkg/utils"
	"net/http"
)

func Router() *chi.Mux {
	r := chi.NewMux()

	//r.Get("/", listItemsHandler)
	r.Post("/", createItemHandler)

	return r
}

func createItemHandler(w http.ResponseWriter, req *http.Request) {
	var account AccountItem
	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ctx := req.Context()

	item, err := createItem(ctx, account.Username, account.Password)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteResponse(item, w, "Success")
}
