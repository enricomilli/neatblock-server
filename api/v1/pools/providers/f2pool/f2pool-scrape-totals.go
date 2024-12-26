package f2pool

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
)

// TODO:
// Find the endpoints for the totals
// Create the types for the totals response
func (provider *F2Pool) ScrapeTotals(observerURL string) (poolproviders.MiningTotals, error) {

	totals := poolproviders.MiningTotals{}

	totalsEndpoint, err := provider.GetTotalsEndpoint(observerURL)
	if err != nil {
		return totals, fmt.Errorf("error generating endpoint: %v\n", err)
	}

	request, err := http.NewRequest("GET", totalsEndpoint, nil)
	if err != nil {
		return totals, fmt.Errorf("could not create request: %v\n", err)
	}
	addTotalsHeaders(request)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return totals, fmt.Errorf("error completing request: %v\n", err)
	}

	htmlDoc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return totals, fmt.Errorf("error creating html doc: %v\n", err)
	}

	queryRes := htmlDoc.Find("div.num")

	for i, s := range queryRes.EachIter() {
		if i < 2 {
			btc, err := strconv.ParseFloat(s.Text(), 64)
			if err != nil {
				return totals, fmt.Errorf("could not parse f2pool html float: %v\n", err)
			}

			switch i {
			case 0:
				totals.TotalBtcMined = btc
			case 1:
				totals.TotalBtcPayout = btc
			}
		}
	}

	return totals, nil
}
