package ethereum

import (
	"github.com/Dharitri-org/blockatlas/platform/bitcoin/blockbook"
	"github.com/Dharitri-org/blockatlas/platform/ethereum/bounce"
	"github.com/Dharitri-org/blockatlas/platform/ethereum/opensea"
	"github.com/Dharitri-org/tw-go-libs/client"
	"github.com/Dharitri-org/tw-go-libs/coin"
	"github.com/Dharitri-org/tw-go-libs/network/middleware"
)

type Platform struct {
	CoinIndex   uint
	client      EthereumClient
	collectible CollectibleClient
}

func InitWithBlockbook(coinType uint, blockbookApi string) *Platform {
	return &Platform{
		CoinIndex: coinType,
		client:    &blockbook.Client{Request: client.InitClient(blockbookApi, middleware.SentryErrorHandler)},
	}
}

func InitWithOpenSea(coinType uint, blockbookApi, collectionApi, collectionKey string) *Platform {
	platform := InitWithBlockbook(coinType, blockbookApi)
	platform.collectible = opensea.InitClient(collectionApi, collectionKey)
	return platform
}

func InitWithBounce(coinType uint, blockbookApi, collectionApi string) *Platform {
	platform := InitWithBlockbook(coinType, blockbookApi)
	platform.collectible = bounce.InitClient(collectionApi)
	return platform
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
