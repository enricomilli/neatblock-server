package f2pool

import (
	"fmt"
	"testing"

	"github.com/enricomilli/neat-server/api/v1/pools"
	"github.com/enricomilli/neat-server/api/v1/pools/providers/f2pool"
	"github.com/stretchr/testify/assert"
)

// TODO: find more pool links and test them
func TestF2Pool_RewardsEndpointBuilder(t *testing.T) {
	provider := &f2pool.F2Pool{}
	testURL := "https://www.f2pool.com/mining-user/6f2e4214d79688ab7697aca78f243d50?user_name=adamantiumnewvz"

	rewardsEndpoint := provider.GetRewardsEndpoint(testURL)
	assert.NotNil(t, rewardsEndpoint)

	fmt.Println("Test pool's rewards endpoint:", rewardsEndpoint)
}

func TestF2Pool_PayoutsEndpointBuilder(t *testing.T) {
	provider := &f2pool.F2Pool{}
	testURL := "https://www.f2pool.com/mining-user/6f2e4214d79688ab7697aca78f243d50?user_name=adamantiumnewvz"

	payoutsEndpoint := provider.GetPayoutsEndpoint(testURL)
	assert.NotNil(t, payoutsEndpoint)

	fmt.Println("Test pool's payouts endpoint:", payoutsEndpoint)
}

func TestF2Pool_ScrapeRewards(t *testing.T) {
	provider := &f2pool.F2Pool{}
	testPool := pools.Pool{
		ID:          "this-test-id",
		ObserverURL: "https://www.f2pool.com/mining-user/6f2e4214d79688ab7697aca78f243d50?user_name=adamantiumnewvz",
	}

	rewards, err := provider.ScrapeDailyRewards(testPool.ObserverURL, testPool.ID)
	assert.NoError(t, err)
	assert.NotNil(t, rewards)

	fmt.Printf("f2pool rewards: %+v\n", rewards)
}
