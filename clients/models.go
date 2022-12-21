package clients

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
