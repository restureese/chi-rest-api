package account

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/oklog/ulid/v2"
	"main/utils"
	"net/http"
)

func Router() *chi.Mux {
	r := chi.NewMux()

	r.Get("/", listItemsHandler)
	r.Post("/", createItemHandler)
	r.Put("/{id}", updateItemHandler)
	r.Delete("/{id}", deleteItemHandler)

	return r
}

// listItemsHandler - Returns all the available APIs
// @Summary Get List Account
// @Description get list account.
// @Tags Account
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Response
// @Router /accounts [get]
// @Security BearerAuth
func listItemsHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	items, err := listItems(ctx)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	var resp struct {
		data interface{}
	}
	resp.data = items

	utils.WriteResponse(items, w, "Success")
}

// createItemHandler - Returns all the available APIs
// @Summary Create Account
// @Description create account.
// @Tags Account
// @Accept  json
// @Produce  json
//
//	@Param account	body CreateAccountItem true "account"
//
// @Success 200 {object} utils.Response
// @Router /accounts [post]
// @Security BearerAuth
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

// updateItemHandler - Returns all the available APIs
// @Summary Update Account
// @Description update account.
// @Tags Account
// @Accept  json
// @Produce  json
// @Param	id	path	string  true  "Account ID"
//
//	@Param account	body CreateAccountItem true "account"
//
// @Success 200 {object} utils.Response
// @Router /accounts/{id} [put]
// @Security BearerAuth
func updateItemHandler(w http.ResponseWriter, req *http.Request) {
	var account AccountItem
	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ctx := req.Context()

	itemId := chi.URLParam(req, "id")

	id, err := ulid.Parse(itemId)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var resp AccountItem

	item, err := makeItemUpdate(ctx, id, account.Username, account.Password)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp = item
	utils.WriteResponse(resp, w, "Success")
}

// deleteItemHandler - Returns all the available APIs
// @Summary Delete Account
// @Description delete account.
// @Tags Account
// @Accept  json
// @Produce  json
// @Param	id	path	string  true  "Account ID"
// @Success 200 {object} utils.Response
// @Router /accounts/{id} [delete]
// @Security BearerAuth
func deleteItemHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	itemId := chi.URLParam(req, "id")

	id, err := ulid.Parse(itemId)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = makeItemDelete(ctx, id)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var resp struct {
		data interface{}
	}

	resp.data = nil

	utils.WriteResponse(resp, w, "Success")
}
