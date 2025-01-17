package algorand

import (
	"github.com/Dharitri-org/tw-go-libs/coin"
)

type Platform struct {
	client Client
}

func Init(api, apiKey string) *Platform {
	return &Platform{
		client: InitClient(api, apiKey),
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Algorand()
}
