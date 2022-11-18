package main

import (
	"crypto_price_tracker/clients"
	"crypto_price_tracker/config"
	"crypto_price_tracker/plerrors"
	"crypto_price_tracker/plog"
	"encoding/json"
	"time"
)

func GetBitCoinPrice() (CoinAPIResponse, error) {
	var coinAPIClient clients.CoinClientService
	coinAPIClient = clients.CoinDesk{}

	resp, err := coinAPIClient.FetchCurrentPrice()

	if err != nil {
		return CoinAPIResponse{}, err
	}

	formattedData := map[string]map[string]string{
		"bitcoin": {
			"EUR": resp.Bpi["EUR"].Rate,
			"USD": resp.Bpi["USD"].Rate,
		},
	}

	formattedDataJSON, err := json.Marshal(formattedData)
	var response CoinAPIResponse
	err = json.Unmarshal(formattedDataJSON, &response)
	if err != nil {
		return CoinAPIResponse{}, plerrors.NewAppError("GetBitCoinPrice", "", "",
			plerrors.InternalServiceError, err.Error(), plog.Params{})
	}

	SetPriceCache("bitcoin_price_USD", formattedData["bitcoin"]["USD"])
	SetPriceCache("bitcoin_price_EUR", formattedData["bitcoin"]["EUR"])
	return response, nil
}

func GetBitCoinPriceCache() map[string]map[string]string {
	rc := clients.RedisClient{}
	valUSD := rc.GetValue("bitcoin_price_USD")
	valEUR := rc.GetValue("bitcoin_price_EUR")

	if valUSD != "" {
		plog.Info("Fetching Redis....", nil)
		formattedData := map[string]map[string]string{
			"bitcoin": {
				"EUR": valEUR,
				"USD": valUSD,
			},
		}
		return formattedData
	}
	return nil
}

func SetPriceCache(key string, val string) {
	rc := clients.RedisClient{}
	expirySeconds, _ := config.GetRedisExpiry()
	rc.SetValue(key, val, time.Duration(expirySeconds)*time.Second)
}
