package viabtc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
)

// TODO: support for pagination
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
