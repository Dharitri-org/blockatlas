package tezos

import (
	"github.com/Dharitri-org/tw-go-libs/client"
	"github.com/Dharitri-org/tw-go-libs/coin"
	"github.com/Dharitri-org/tw-go-libs/network/middleware"
)

type Platform struct {
	client      Client
	rpcClient   RpcClient
	bakerClient BakerClient
}

func Init(api, rpc, baker string) *Platform {
	p := &Platform{
		client:      Client{client.InitClient(api, middleware.SentryErrorHandler)},
		rpcClient:   RpcClient{client.InitClient(rpc, middleware.SentryErrorHandler)},
		bakerClient: BakerClient{client.InitClient(baker, middleware.SentryErrorHandler)},
	}
	p.client.SetTimeout(35)
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Tezos()
}
