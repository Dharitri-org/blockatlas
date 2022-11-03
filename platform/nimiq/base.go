package nimiq

import (
	"github.com/Dharitri-org/tw-go-libs/client"
	"github.com/Dharitri-org/tw-go-libs/coin"
	"github.com/Dharitri-org/tw-go-libs/network/middleware"
)

type Platform struct {
	client Client
}

func Init(api string) *Platform {
	return &Platform{
		client: Client{client.InitJSONClient(api, middleware.SentryErrorHandler)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Nimiq()
}
