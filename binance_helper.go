package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const BinanceUrl = "https://api.binance.com/api/v3/ticker/price?symbol=BTCUAH"

type BinanceResponse struct {
	Symbol string
	Price  string
}

func getBtcUahRate() BinanceResponse {
	response, err := http.Get(BinanceUrl)
	check(err)

	responseData, err := ioutil.ReadAll(response.Body)
	check(err)

	var binanceResponse BinanceResponse
	json.Unmarshal([]byte(responseData), &binanceResponse)

	return binanceResponse
}
