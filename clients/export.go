package clients

type CoinClientService interface {
	FetchCurrentPrice() (CoinResponse, error)
}
