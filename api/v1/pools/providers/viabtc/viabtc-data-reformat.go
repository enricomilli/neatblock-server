package viabtc

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/alpacahq/alpacadecimal"
	apiutil "github.com/enricomilli/neat-server/api/api-utils"
	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
)

// Reformats rewards response from viabtc to a neatblock reward
func reformatData(data ViaBTCRewardsResponse, poolID string) ([]poolproviders.MiningReward, error) {

	formated := []poolproviders.MiningReward{}

	for _, day := range data.Data.Data {

		parsedHashrate, err := strconv.ParseFloat(day.Hashrate, 64)
		if err != nil {
			return nil, errors.New("could not cover string to interger")
		}

		// viabtc hashrate returns in a weird format
		newHashrate := apiutil.RoundFloat(parsedHashrate/1000000000000, 2)

		rewardBTC, err := alpacadecimal.NewFromString(day.RewardBTC)
		if err != nil {
			return nil, errors.New("could not cover string to decimal")
		}

		rewardTxFee, err := alpacadecimal.NewFromString(day.RewardTxFee)
		if err != nil {
			return nil, errors.New("could not cover string to decimal")
		}

		rewardTotal := rewardBTC.Add(rewardTxFee)

		payout, err := alpacadecimal.NewFromString(day.PayoutBTC)
		if err != nil {
			return nil, fmt.Errorf("could not cover payout: %v", err)
		}

		formatedDay := poolproviders.MiningReward{
			Date:      day.Date,
			PoolID:    poolID,
			Hashrate:  newHashrate,
			BtcReward: rewardBTC.InexactFloat64(),
			BtcTxFee:  rewardTxFee.InexactFloat64(),
			Total:     rewardTotal.InexactFloat64(),
			Payout:    payout.InexactFloat64(),
		}

		formated = append(formated, formatedDay)
	}

	if len(formated) < 1 {
		return formated, fmt.Errorf("did not retrieve any data")
	}

	return formated, nil
}
