package f2pool

import poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"

// TODO:
// Find the endpoints for the totals
// Create the types for the totals response
func (provider *F2Pool) ScrapeTotals(observerURL string) (poolproviders.MiningTotals, error) {

	return poolproviders.MiningTotals{}, nil
}
