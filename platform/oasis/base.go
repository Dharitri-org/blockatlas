package oasis

import (
	"github.com/Dharitri-org/tw-go-libs/client"
	"github.com/Dharitri-org/tw-go-libs/coin"
	"github.com/Dharitri-org/tw-go-libs/network/middleware"
)

type Platform struct {
	client Client
}

func Init(api string) *Platform {
	p := &Platform{
		client: Client{client.InitClient(api, middleware.SentryErrorHandler)},
	}
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Oasis()
}
