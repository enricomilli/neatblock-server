package f2pool

import (
	"fmt"

	"github.com/alpacahq/alpacadecimal"
	apiutil "github.com/enricomilli/neat-server/api/api-utils"
	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
)

func reformatF2Rewards(payouts []F2PoolPayoutData, rewards []F2PoolIncomeData, poolID string) ([]poolproviders.MiningReward, error) {

	response := []poolproviders.MiningReward{}

	pLen := len(payouts)
	rLen := len(rewards)

	if pLen != rLen {
		return nil, fmt.Errorf("payouts and rewards are not of equal lengths")
	}
	if pLen < 1 || rLen < 1 {
		return nil, fmt.Errorf("more than 1 payout and reward required")
	}

	for i := 0; i < pLen; i++ {

		date := int64(payouts[i].CreatedAt)

		rewardBTC := alpacadecimal.NewFromFloat(rewards[i].Amount)
		rewardTxFee := alpacadecimal.NewFromFloat(rewards[i].Txfee)
		rewardTotal := rewardBTC.Add(rewardTxFee)

		payoutAmount, err := alpacadecimal.NewFromString(payouts[i].Amount)
		if err != nil {
			return nil, err
		}

		day := poolproviders.MiningReward{
			PoolID:    poolID,
			Date:      apiutil.UnixToDate(date),
			Hashrate:  float64(rewards[i].HashRate),
			BtcReward: rewards[i].Amount,
			BtcTxFee:  rewards[i].Txfee,
			Total:     rewardTotal.InexactFloat64(),
			Payout:    payoutAmount.InexactFloat64(),
		}

		response = append(response, day)
	}

	return response, nil
}
