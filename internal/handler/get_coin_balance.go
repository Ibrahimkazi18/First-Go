package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Ibrahimkazi18/First-Go/api"
	"github.com/Ibrahimkazi18/First-Go/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinsBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

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

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinsBalanceResult{
		Balance: (*tokenDetails).Coin,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
