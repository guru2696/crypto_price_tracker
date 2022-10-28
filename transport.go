package main

import (
	"encoding/json"
	"net/http"
)

type CoinResponse struct {
	Time       map[string]string               `json:"time"`
	Disclaimer string                          `json:"disclaimer"`
	ChartName  string                          `json:"chartName"`
	Bpi        map[string]CoinResponseCurrency `json:"bpi"`
}

type CoinResponseCurrency struct {
	Code        string  `json:"code"`
	Symbol      string  `json:"symbol"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}

func GetBitCoinPrice() (CoinResponse, error) {
	resp, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		return CoinResponse{}, err
	}

	defer resp.Body.Close()

	var responseJSON CoinResponse
	json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}
