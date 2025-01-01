package api

//Coins Balance Params
type CoinsBalanceParams struct {
	Username string
}

//Coins Balance Result
type CoinsBalanceResult struct {
	Code int

	Balance int64
}

//Error Result
type ErrorResult struct {
	Code int

	Message string
}
