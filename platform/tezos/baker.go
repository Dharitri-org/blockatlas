package tezos

import (
	"math"
	"strconv"
	"time"

	"github.com/Dharitri-org/blockatlas/pkg/blockatlas"
	"github.com/Dharitri-org/blockatlas/services/assets"
	"github.com/Dharitri-org/tw-go-libs/client"
	"github.com/Dharitri-org/tw-go-libs/coin"
	"github.com/Dharitri-org/tw-go-libs/types"
)

const (
	cacheTime = 1 * time.Hour
)

type BakerClient struct {
	client.Request
}

func (c *BakerClient) GetBakers() (validators blockatlas.StakeValidators, err error) {
	var bakers []Baker
	err = c.GetWithCache(&bakers, "v2/bakers", nil, cacheTime)
	if err != nil {
		return
	}
	assetsValidators, err := assets.GetValidatorsInfo(coin.Tezos())
	if err != nil {
		return
	}
	validatorMap := assetsValidators.ToMap()
	for _, baker := range bakers {
		if av, ok := validatorMap[baker.Address]; ok {
			validators = append(validators, NormalizeStakeValidator(baker, av))
		}
	}
	return
}

func NormalizeStakeValidator(baker Baker, assetValidator assets.AssetValidator) blockatlas.StakeValidator {
	status := true
	if baker.FreeSpace < 0 || baker.ServiceHealth != "active" || !baker.OpenForDelegation {
		status = false
	}

	amount := uint64(math.Ceil(baker.MinDelegation))

	return blockatlas.StakeValidator{
		ID:     baker.Address,
		Status: status,
		Info: blockatlas.StakeValidatorInfo{
			Name:        assetValidator.Name,
			Description: assetValidator.Description,
			Website:     assetValidator.Website,
			Image:       assets.GetImageURL(coin.Tezos(), baker.Address),
		},
		Details: blockatlas.StakingDetails{
			Reward: blockatlas.StakingReward{
				Annual: math.Round(baker.EstimatedRoi*10000) / 100,
			},
			LockTime:      LockTime,
			MinimumAmount: types.Amount(strconv.FormatUint(amount, 10)),
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}
