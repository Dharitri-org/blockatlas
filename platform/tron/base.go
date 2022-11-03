package tron

import (
	"github.com/Dharitri-org/tw-go-libs/client"
	"github.com/Dharitri-org/tw-go-libs/coin"
	"github.com/Dharitri-org/tw-go-libs/network/middleware"
)

type Platform struct {
	client Client
}

func Init(api, apiKey string) *Platform {
	request := client.InitClient(api, middleware.SentryErrorHandler)
	//TODO: Add when ready
	//request.Headers = map[string]string{"TRON-PRO-API-KEY": apiKey}
	return &Platform{
		client: Client{request},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Tron()
}
