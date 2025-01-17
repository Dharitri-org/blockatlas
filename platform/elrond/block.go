package elrond

import "github.com/Dharitri-org/tw-go-libs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	return p.client.GetBlockByNumber(num)
}
