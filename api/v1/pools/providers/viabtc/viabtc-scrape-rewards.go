package viabtc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/alpacahq/alpacadecimal"
	apiutil "github.com/enricomilli/neat-server/api/api-utils"
	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
)

func (provider *ViaBTC) ScrapeDailyRewards(observerURL string, poolID string) ([]poolproviders.MiningReward, error) {

	url, err := url.Parse(observerURL)
	if err != nil {
		return []poolproviders.MiningReward{}, fmt.Errorf("could not parse url: %v", err)
	}
	queries := url.Query()

	accessKey := queries.Get("access_key")
	if accessKey == "" {
		return []poolproviders.MiningReward{}, fmt.Errorf("no access key found in url: %s", observerURL)
	}

	coin := queries.Get("coin")
	if coin == "" {
		return []poolproviders.MiningReward{}, fmt.Errorf("no coin parameter in url: %s", observerURL)
	}

	// user_id does not need an error check because some links don't need it
	webEndpoint := provider.GetRewardsEndpoint(queries.Get("user_id"), accessKey, coin)

	req, err := http.NewRequest("GET", webEndpoint, nil)
	if err != nil {
		return []poolproviders.MiningReward{}, fmt.Errorf("could not init request: %v", err)
	}
	addViaBTCHeaders(req)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return []poolproviders.MiningReward{}, fmt.Errorf("error completing request: %v", err)
	}

	defer response.Body.Close()
	rewardsRes := ViaBTCRewardsResponse{}
	err = json.NewDecoder(response.Body).Decode(&rewardsRes)
	if err != nil {
		return []poolproviders.MiningReward{}, fmt.Errorf("error parsing json body: %v", err)
	}

	// format from viabtc rewards to neatblock mining rewards

	return reformatData(rewardsRes, poolID)
}

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
			Date:          day.Date,
			PoolReference: poolID,
			Hashrate:      newHashrate,
			Reward:        rewardBTC.InexactFloat64(),
			TxFee:         rewardTxFee.InexactFloat64(),
			RewardPlusTx:  rewardTotal.InexactFloat64(),
			Payout:        payout.InexactFloat64(),
		}

		formated = append(formated, formatedDay)
	}

	if len(formated) < 1 {
		return formated, fmt.Errorf("did not retrieve any data")
	}

	return formated, nil
}
