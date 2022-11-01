package main

import (
	"crypto_price_tracker/clients"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func GetBitCoinPrice() (interface{}, error) {
	var coinAPIClient clients.CoinClientService
	coinAPIClient = clients.CoinDesk{}

	resp, err := coinAPIClient.FetchCurrentPrice()

	if err != nil {
		return clients.CoinResponse{}, err
	}

	formattedData := map[string]map[string]string{
		"bitcoin": {
			"EUR": resp.Bpi["EUR"].Rate,
			"USD": resp.Bpi["USD"].Rate,
		},
	}

	SetPriceCache("bitcoin_price_USD", formattedData["bitcoin"]["USD"])
	SetPriceCache("bitcoin_price_EUR", formattedData["bitcoin"]["EUR"])
	return formattedData, nil
}

func GetBitCoinPriceCache() map[string]map[string]string {
	rc := clients.RedisClient{}
	valUSD := rc.GetValue("bitcoin_price_USD")
	valEUR := rc.GetValue("bitcoin_price_EUR")

	if valUSD != "" {
		log.Default().Println("Fetching Redis....")
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
	expirySeconds, _ := strconv.Atoi(os.Getenv("EXPIRY_SECONDS"))
	fmt.Println(expirySeconds)
	rc.SetValue(key, val, time.Duration(expirySeconds)*time.Second)
}
