package main

type CoinAPIResponse struct {
	Bitcoin struct {
		EUR string `json:"EUR"`
		USD string `json:"USD"`
	} `json:"bitcoin"`
}
