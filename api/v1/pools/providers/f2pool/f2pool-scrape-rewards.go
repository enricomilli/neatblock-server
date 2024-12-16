package f2pool

import poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"

// TODO:
// There are two requests you have to make for F2Pool
// One of the hashrate and one for the profit share
// Need to give the option for users to put in their take home revenue share
func (provider *F2Pool) ScrapeDailyRewards(observerURL string, poolID string) ([]poolproviders.MiningReward, error) {

	return []poolproviders.MiningReward{}, nil
}
