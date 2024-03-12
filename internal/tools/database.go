
package tools

import (
    log "github.com/sirupsen/logrus"
)

// Database types
type LoginDetails struct {
    AuthToken string
    Username string
}

type BalanceDetails struct {
    Balance int64
    Username string
}

type DatabaseInterface interface {
    GetUserLoginDetails(username string) *LoginDetails
    GetUserBalance(username string) *BalanceDetails
    SetupDatabase() error
}


// Returns DatabaseInterface or an error
func NewDatabase() (*DatabaseInterface, error) {
    // mockDB will implement the interface
    var database DatabaseInterface = &mockDB{}

    var err error = database.SetupDatabase()

    if err != nil {
        log.Error(err)
        return nil, err
    }

    return &database, nil
}


