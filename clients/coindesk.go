package clients

import (
	"crypto_price_tracker/config"
	"crypto_price_tracker/plerrors"
	"crypto_price_tracker/plog"
	"encoding/json"
	"net/http"
)

type CoinDesk struct {
	URL string
}

func Configure() *CoinDesk {
	return &CoinDesk{
		URL: config.GetCryptoURL(),
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
	err = json.NewDecoder(resp.Body).Decode(&responseJSON)
	if err != nil {
		return CoinResponse{}, plerrors.NewAppError("FetchCurrentPrice", "",
			"Response Decoding Error", plerrors.InternalServiceError,
			err.Error(), plog.Params{})
	}

	return responseJSON, nil
}
