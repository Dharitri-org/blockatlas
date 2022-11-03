package bitcoin

import (
	"github.com/Dharitri-org/blockatlas/platform/bitcoin/blockbook"
	"github.com/Dharitri-org/tw-go-libs/client"
	"github.com/Dharitri-org/tw-go-libs/coin"
	"github.com/Dharitri-org/tw-go-libs/network/middleware"
)

type Platform struct {
	client    blockbook.Client
	CoinIndex uint
}

func Init(coin uint, api string) *Platform {
	return &Platform{
		CoinIndex: coin,
		client:    blockbook.Client{Request: client.InitClient(api, middleware.SentryErrorHandler)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) GetAddressesFromXpub(xpub string) ([]string, error) {
	tokens, err := p.client.GetAddressesFromXpub(xpub)
	addresses := make([]string, 0)
	for _, token := range tokens {
		addresses = append(addresses, token.Name)
	}
	return addresses, err
}
