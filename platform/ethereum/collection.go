package ethereum

import (
	"github.com/Dharitri-org/tw-go-libs/types"
)

func (p *Platform) GetCollections(owner string) (types.CollectionPage, error) {
	return p.collectible.GetCollections(owner, p.CoinIndex)
}

func (p *Platform) GetCollectibles(owner, collectionID string) (types.CollectiblePage, error) {
	return p.collectible.GetCollectibles(owner, collectionID, p.CoinIndex)
}
