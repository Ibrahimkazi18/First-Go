package api

import (
	"encoding/json"
	"net/http"
)

// Coins Balance Params
type CoinsBalanceParams struct {
	Username string
}

// Coins Balance Result
type CoinsBalanceResult struct {
	Code int

	Balance int64
}

// Error Result
type Error struct {
	Code int

	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured", http.StatusInternalServerError)
	}
)
