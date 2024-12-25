package pools

import (
	"fmt"

	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
)

func (pool *Pool) ScrapeMiningData() error {

	poolProvider, err := pool.NewProviderInterface()
	if err != nil {
		return fmt.Errorf("could not create pool interface: %w", err)
	}

	scrapedTotals, err := poolProvider.ScrapeTotals(pool.ObserverURL)
	if err != nil {
		return fmt.Errorf("could not scrape totals: %w", err)
	}

	totalsChanged := checkIfTotalsChange(pool, &scrapedTotals)
	if !totalsChanged {
		return nil
	}

	scrapedRewards, err := poolProvider.ScrapeDailyRewards(pool.ObserverURL, pool.ID)
	if err != nil {
		return fmt.Errorf("could not scrape daily rewards: %w", err)
	}

	storedPoolRewards, err := pool.GetAllRewards()
	if err != nil {
		return err
	}

	newRewards, hasNewData := CheckForNewData(storedPoolRewards, scrapedRewards)
	if !hasNewData {
		return nil
	}

	err = pool.StoreRewards(newRewards)
	if err != nil {
		return err
	}

	err = pool.StorePoolStructState()
	if err != nil {
		return fmt.Errorf("could not save pool: %w", err)
	}

	return nil // no errors
}

func checkIfTotalsChange(pool *Pool, scrapedTotals *poolproviders.MiningTotals) bool {

	if pool.TotalBtcMined == scrapedTotals.TotalBtcMined {
		return false
	}

	return true
}

func CheckForNewData(storedRewards []poolproviders.MiningReward, scrapedRewards []poolproviders.MiningReward) (newRewards []poolproviders.MiningReward, hasNewRewards bool) {

	hash := map[string]any{}
	hasNewRewards = false

	for _, reward := range storedRewards {
		hash[reward.Date] = 1
	}

	for _, reward := range scrapedRewards {

		_, exists := hash[reward.Date]
		if !exists {
			if !hasNewRewards {
				hasNewRewards = true
			}
			newRewards = append(newRewards, reward)
		}

	}

	if !hasNewRewards {
		return newRewards, false
	}

	return newRewards, true
}
