package theta

import (
	"github.com/Dharitri-org/tw-go-libs/client"
	"github.com/Dharitri-org/tw-go-libs/coin"
	"github.com/Dharitri-org/tw-go-libs/network/middleware"
)

type Platform struct {
	client Client
}

func Init(api, key string) *Platform {
	request := client.InitClient(api, middleware.SentryErrorHandler)
	request.Headers = map[string]string{"x-api-token": key}
	return &Platform{
		client: Client{request},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.THETA]
}
