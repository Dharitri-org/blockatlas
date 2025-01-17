package ripple

import "github.com/Dharitri-org/tw-go-libs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	if srcBlock, err := p.client.GetBlockByNumber(num); err == nil {
		txs := NormalizeTxs(srcBlock)
		return &types.Block{
			Number: num,
			Txs:    txs,
		}, nil
	} else {
		return nil, err
	}
}
