package api

import (
    "encoding/json"
    "net/http"
)

// Account Balance Parameters
type AccountBalanceParams struct{
    Username string
}

type AccountBalanceResponse struct {
    Code int // 200 for success, 500 for internal server errors, 400 for input errors
    Balance int64 // Balance in the account
}

type Error struct {
    Code int // 400 - 500
    Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
    resp := Error{
        Code: code,
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
        writeError(w, "An unexpected error ocurred", http.StatusInternalServerError)
    }
)
