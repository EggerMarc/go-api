package tools

import (
    "time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails {
    "marc": {
        AuthToken: "001aaa",
        Username: "marc",
    },
    "fayum": {
        AuthToken: "4091ll",
        Username: "fayum",
    },
    "logan": {
        AuthToken: "411aaa",
        Username: "logan",
    },
    "patricia": {
        AuthToken: "2m1ss1",
        Username: "patricia",
    }
}

var mockBalanceDetails = map[string]BalanceDetails {
    "marc": {
        Balance: 1,
        Username: "marc",
    },
    "fayum": {
        Balance: 4100,
        Username: "fayum",
    },

    "logan": {
        Balance: 2999,
        Username: "logan",
    },
    "patricia": {
        Balance: 1021,
        Username: "patricia",

    }
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
    time.Sleep(time.Milliseconds * 20)
    var clientData = LoginDetails{}

    clientData, ok := mockLoginDetails[username]
    if !ok {
        return nil
    }
    return &clientData
}

func (d *mockDB) GetUserBalanceDetails(username string) *BalanceDetails {
    time.sleep(time.Milliseconds * 20)
    var clientData = BalanceDetails{}

    clientData, ok := mockBalanceDetails[username]

    if !ok {
        return nil
    }
    return &clientData
}


func (d *mockDB) SetupDatabase() error {
    return nil
}




