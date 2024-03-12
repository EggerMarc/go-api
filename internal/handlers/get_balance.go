package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/eggermarc/go-api/api"
    "github.com/eggermarc/go-api/internal/tools"
    log "github.com/sirupsen/logrus"
    "github.com/gorilla/schema"
)

func GetUserBalance(w http.ResponseWriter, r *http.Request) {
    var params = api.UserBalanceParams{}
    var decoder *schema.Decoder = schema.NewDecoder()
    var err error

    // this will take in the parameters and add the http info to the error
    err = decoder.Decode(&params, r.URL.Query())

    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }
    var database *tools.DatabaseInterface
    database, err = tools.NewDatabase()

    if err != nil {
        api.InternalErrorHandler(w)
        return
    }

    var tokenDetails *tools.BalanceDetails
    tokenDetails = (*database).GetUserBalance(params.Username)
    if tokenDetails == nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    var response = api.AccountBalanceResponse{
        Balance: (*tokenDetails).Balance,
        Code: http.StatusOK,
    }

    w.Header().Set("Content-Type", "application/json")

    err = json.NewEncoder(w).Encode(response)

    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }
}

