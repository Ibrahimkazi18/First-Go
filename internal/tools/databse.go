package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database Collection
type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coin     int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetUpDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}
	var err error = database.SetUpDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
