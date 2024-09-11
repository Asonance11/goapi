package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"samuel": {
		AuthToken: "123ABC",
		Username:  "samuel",
	},
	"esther": {
		AuthToken: "456DEF",
		Username:  "esther",
	},
	"betty": {
		AuthToken: "789GHI",
		Username:  "betty",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"samuel": {
		Coins:    500,
		Username: "samuel",
	},
	"esther": {
		Coins:    200,
		Username: "esther",
	},
	"betty": {
		Coins:    400,
		Username: "betty",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
