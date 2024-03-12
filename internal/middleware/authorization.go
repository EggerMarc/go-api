package middleware

import (
    "errors"
    "net/http"

    "github.com/eggermarc/go-api/api"
    "github.com/eggermarc/go-api/internal/tools"
    log "github.com/sirupsen/logrus"
)

// Create an Unauthorized error
var UnAuthorizedError = errors.New("Invalid username or token")

// Takes in an http handler, and returns the same
func Authorization(next http.Handler) http.Handler {
    // http is handled, by the function we define below
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Define logic for authorization
        // get the username from the http
        var username string = r.URL.Query().Get("username")
        // get the token from http
        var token = r.Header.Get("Authorization")
        // set an error variable
        var error error

        // if no username or token passed, throw error
        if username == "" || token == "" {
            log.Error(UnAuthorizedError)
            api.RequestErrorHandler(w, UnAuthorizedError)
            return
        }

        // set a pointer to the database
        var database *tools.DatabaseInterface

        // unpack
        database, err = tools.NewDatabase()
        if err != nil {
            api.InternalErrorHandler(w)
            return
        }

        // set a pointer to the loging
        var loginDetails *tools.loginDetails
        // from the database, get the logins
        loginDetails = (*database).GetUserLoginDetails(username)

        if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
            log.Error(UnAuthorizedError)
            api.RequestErrorHandler(w, UnAuthorizedError)
            return
        }

        // calls in the next middleware or endpoint
        next.Serve(w, r)
    })
}

