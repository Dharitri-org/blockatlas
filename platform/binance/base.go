package binance

import (
	"github.com/Dharitri-org/blockatlas/platform/binance/staking"
	"github.com/Dharitri-org/tw-go-libs/coin"
)

type Platform struct {
	client        Client
	stakingClient staking.Client
}

func Init(api, apiKey, stakingApi string) *Platform {
	p := Platform{
		client:        InitClient(api, apiKey),
		stakingClient: staking.InitClient(stakingApi),
	}
	return &p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Binance()
}
