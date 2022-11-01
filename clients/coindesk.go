package clients

import (
	"encoding/json"
	"net/http"
)

type CoinDesk struct {
	URL string
}

func Configure() *CoinDesk {
	return &CoinDesk{
		URL: "https://api.coindesk.com/v1/bpi/currentprice.json",
	}
}

func (c CoinDesk) FetchCurrentPrice() (CoinResponse, error) {
	client := Configure()
	resp, err := http.Get(client.URL)
	if err != nil {
		return CoinResponse{}, err
	}

	defer resp.Body.Close()

	var responseJSON CoinResponse
	json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}
