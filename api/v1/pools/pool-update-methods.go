package pools

import (
	"fmt"

	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
)

func (pool *Pool) UpdatePoolData() error {

	poolProvider, err := pool.NewProviderInterface()
	if err != nil {
		return fmt.Errorf("could not create pool interface: %w", err)
	}

	scrapedTotals, err := poolProvider.ScrapeTotals(pool.ObserverURL)
	if err != nil {
		return fmt.Errorf("could not scrape n store totals: %w", err)
	}

	totalsChanged := checkIfTotalsChange(pool, &scrapedTotals)
	if !totalsChanged {
		return nil
	}

	scrapedRewards, err := poolProvider.ScrapeDailyRewards(pool.ObserverURL)
	if err != nil {
		return fmt.Errorf("could not scrape daily rewards: %w", err)
	}

	hasNewData := CheckForNewData(pool, scrapedTotals, scrapedRewards)
	if !hasNewData {
		return nil
	}

	err = pool.SaveToDB()
	if err != nil {
		return fmt.Errorf("could not save pool: %w", err)
	}

	return nil // no errors
}

func checkIfTotalsChange(pool *Pool, scrapedTotals *poolproviders.NeatblockTotals) bool {

	if pool.BTCRevenue == scrapedTotals.TotalBtcProfit {
		return false
	}

	return true
}

func CheckForNewData(pool *Pool, sTotals poolproviders.NeatblockTotals, sRewards []poolproviders.NeatblockReward) bool {

	return false
}
