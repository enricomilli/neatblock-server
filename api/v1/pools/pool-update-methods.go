package pools

import "fmt"

func (pool *Pool) UpdatePoolData() error {

	// create interface based off provider
	poolProvider, err := pool.NewProviderInterface()
	if err != nil {
		return fmt.Errorf("could not create pool interface: %w", err)
	}

	newTotalsData, err := poolProvider.ScrapeTotals(pool.ObserverURL)
	if err != nil {
		return fmt.Errorf("could not scrape totals: %w", err)
	}

	newRewardsData, err := poolProvider.ScrapeDailyRewards(pool.ObserverURL)
	if err != nil {
		return fmt.Errorf("could not scrape daily rewards: %w", err)
	}

	fmt.Println("New Totals:", newTotalsData)
	fmt.Println("New Rewards:", newRewardsData)

	// TODO:
	// check if theres new data
	// i've written this function for bps
	hasNewData := true

	if !hasNewData {
		return nil
	}

	err = pool.SaveToDB()
	// store the new data
	if err != nil {
		return fmt.Errorf("could not save pool: %w", err)
	}

	return nil
}
