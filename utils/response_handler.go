package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

func writeMessage(data interface{}, w http.ResponseWriter, status int, msg string) {

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(&Response{
		Status:  status,
		Data:    data,
		Message: msg,
	})
}

func WriteError(w http.ResponseWriter, status int, err error) {
	writeMessage(nil, w, status, err.Error())
}

func WriteResponse(data interface{}, w http.ResponseWriter, msg string) {
	writeMessage(data, w, http.StatusOK, msg)
}
