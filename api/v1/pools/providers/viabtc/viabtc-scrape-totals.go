package viabtc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
)

func (provider *ViaBTC) ScrapeTotals(observerURL string) (poolproviders.MiningTotals, error) {

	url, err := url.Parse(observerURL)
	if err != nil {
		return poolproviders.MiningTotals{}, fmt.Errorf("could not parse url: %v", err)
	}
	queries := url.Query()

	accessKey := queries.Get("access_key")
	if accessKey == "" {
		return poolproviders.MiningTotals{}, fmt.Errorf("no access key found in url: %s", observerURL)
	}

	coin := queries.Get("coin")
	if coin == "" {
		return poolproviders.MiningTotals{}, fmt.Errorf("no coin parameter in url: %s", observerURL)
	}

	// user_id does not need an error check because some links don't need it
	webEndpoint := provider.GetTotalsEndpoint(queries.Get("user_id"), accessKey, coin)

	req, err := http.NewRequest("GET", webEndpoint, nil)
	if err != nil {
		return poolproviders.MiningTotals{}, fmt.Errorf("could not init request: %v", err)
	}
	addViaBTCHeaders(req)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return poolproviders.MiningTotals{}, fmt.Errorf("error completing request: %v", err)
	}
	defer response.Body.Close()

	viaBtcResponse := ViaBTCTotalsResponse{}
	err = json.NewDecoder(response.Body).Decode(&viaBtcResponse)
	if err != nil {
		return poolproviders.MiningTotals{}, fmt.Errorf("could not parse json body: %v", err)
	}

	totalProfit, err := strconv.ParseFloat(viaBtcResponse.Data.ProfitTotal, 64)
	if err != nil {
		return poolproviders.MiningTotals{}, fmt.Errorf("could not parse float from string: %s", viaBtcResponse.Data.ProfitTotal)
	}

	fmt.Println(totalProfit)

	return poolproviders.MiningTotals{TotalBtcProfit: totalProfit}, nil
}
