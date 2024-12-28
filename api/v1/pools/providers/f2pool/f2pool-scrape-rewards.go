package f2pool

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
)

// TODO:
// There are two requests you have to make for F2Pool
// One of the hashrate and one for the profit share
// Need to give the option for users to put in their take home revenue share
func (provider *F2Pool) ScrapeDailyRewards(observerURL string, poolID string) ([]poolproviders.MiningReward, error) {

	rewards := []poolproviders.MiningReward{}

	// rewards contains the hashrate, total mined
	rewardsEndpoint := provider.GetRewardsEndpoint(observerURL)

	rewardsReq, err := http.NewRequest("GET", rewardsEndpoint, nil)
	if err != nil {
		return rewards, err
	}
	addRewardsHeaders(rewardsReq, observerURL)

	response, err := http.DefaultClient.Do(rewardsReq)
	if err != nil {
		return rewards, fmt.Errorf("error completing request: %v", err)
	}
	defer response.Body.Close()

	f2RewardsResponse := &F2PoolRewardsResponse{}
	err = json.NewDecoder(response.Body).Decode(f2RewardsResponse)
	if err != nil {
		return rewards, err
	}

	// payouts contains the payout and the distribution
	payoutsEndpoint := provider.GetPayoutsEndpoint(observerURL)

	payoutsReq, err := http.NewRequest("GET", payoutsEndpoint, nil)
	if err != nil {
		return rewards, err
	}
	addRewardsHeaders(payoutsReq, observerURL)

	response, err = http.DefaultClient.Do(payoutsReq)
	if err != nil {
		return rewards, fmt.Errorf("error completing request: %v", err)
	}
	defer response.Body.Close()

	f2PayoutResponse := &F2PoolPayoutResponse{}
	err = json.NewDecoder(response.Body).Decode(f2PayoutResponse)
	if err != nil {
		return rewards, err
	}

	hasRevShare := checkForRevShare(f2PayoutResponse.Data[0])
	if hasRevShare {
		f2PayoutResponse.Data = removePayoutsToProvider(f2PayoutResponse.Data)
	}

	return reformatF2Rewards(f2PayoutResponse.Data, f2RewardsResponse.Data.IncomeData, poolID)
}

func removePayoutsToProvider(payouts []F2PoolPayoutData) []F2PoolPayoutData {

	response := []F2PoolPayoutData{}
	distributionMap := map[float64]string{}

	for _, payout := range payouts {
		percentStr := strings.Split(payout.Comment, "%")[0]
		percentFloat, _ := strconv.ParseFloat(percentStr, 64)
		distributionMap[percentFloat] = percentStr
	}

	highestDist := getHighestRevShare(distributionMap)

	// assuming the highest revenue share is going to the owner
	for _, payout := range payouts {
		if strings.HasPrefix(payout.Comment, highestDist) {
			response = append(response, payout)
		}
	}

	return response
}

// return the string of the higest float
func getHighestRevShare(distMap map[float64]string) string {
	highest := 0.0
	for key := range distMap {
		if key > highest {
			highest = key
		}
	}
	return distMap[highest]
}

func checkForRevShare(payout F2PoolPayoutData) bool {
	if payout.Type == "revenue_distribution" {
		return true
	}
	return false
}
