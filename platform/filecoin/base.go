package filecoin

import (
	"github.com/Dharitri-org/blockatlas/platform/filecoin/explorer"
	"github.com/Dharitri-org/blockatlas/platform/filecoin/rpc"
	"github.com/Dharitri-org/tw-go-libs/client"
	"github.com/Dharitri-org/tw-go-libs/coin"
	"github.com/Dharitri-org/tw-go-libs/network/middleware"
)

type Platform struct {
	client   rpc.Client
	explorer explorer.Client
}

func Init(api, explorerApi string) *Platform {
	p := &Platform{
		client:   rpc.Client{Request: client.InitClient(api, middleware.SentryErrorHandler)},
		explorer: explorer.Client{Request: client.InitClient(explorerApi, middleware.SentryErrorHandler)},
	}
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Filecoin()
}
