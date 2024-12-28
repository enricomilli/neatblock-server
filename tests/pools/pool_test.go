package tests

import (
	"fmt"
	"testing"

	"github.com/enricomilli/neat-server/api/v1/pools"
	"github.com/stretchr/testify/assert"
)

func TestPool_NewProviderInterface(t *testing.T) {
	testPool := pools.Pool{
		ObserverURL: "https://www.viabtc.com/observer/dashboard?access_key=c0a80baded29bb989c06458a8d410a13&coin=BTC",
		Name:        "Test pool",
		UserID:      "d0d4ad7e-6c89-4a44-8e58-2b0d01170c2f",
	}

	provider, err := testPool.NewProviderInterface()
	assert.NoError(t, err)
	assert.NotNil(t, provider, "Provider should not be nil")
}

func TestPool_ValidateURL(t *testing.T) {
	testCases := []struct {
		name        string
		observerURL string
		expectError bool
	}{
		{
			name:        "Valid URL",
			observerURL: "https://www.viabtc.com/observer/dashboard?access_key=c0a80baded29bb989c06458a8d410a13&coin=BTC",
			expectError: false,
		},
		{
			name:        "Invalid URL",
			observerURL: "invalid-url",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pool := pools.Pool{ObserverURL: tc.observerURL}
			provider, err := pool.NewProviderInterface()
			if tc.expectError {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}

			err = provider.ValidateURL(pool.ObserverURL)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// func TestPool_ScrapeTotals(t *testing.T) {
// 	testPools := []pools.Pool{
// 		{
// 			ObserverURL: "https://www.viabtc.com/observer/dashboard?access_key=c0a80baded29bb989c06458a8d410a13&coin=BTC",
// 			Name:        "viabtc pool",
// 			OwnerID:     "d0d4ad7e-6c89-4a44-8e58-2b0d01170c2f",
// 		},
// 		{
// 			ObserverURL: "https://www.f2pool.com/mining-user/6f2e4214d79688ab7697aca78f243d50?user_name=adamantiumnewvz",
// 			Name:        "f2pool",
// 			OwnerID:     "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
// 		},
// 	}

// 	for _, pool := range testPools {
// 		provider, err := pool.NewProviderInterface()
// 		assert.NoError(t, err)

// 		totals, err := provider.ScrapeTotals(pool.ObserverURL)
// 		assert.NoError(t, err)
// 		assert.NotNil(t, totals, "Totals should not be nil")
// 	}

// }

func Test_ScrapeRewards(t *testing.T) {
	testPools := []pools.Pool{
		{
			ObserverURL: "https://www.viabtc.com/observer/dashboard?access_key=c0a80baded29bb989c06458a8d410a13&coin=BTC",
			Name:        "viabtc pool",
			UserID:      "d0d4ad7e-6c89-4a44-8e58-2b0d01170c2f",
		},
		{
			ObserverURL: "https://www.f2pool.com/mining-user/6f2e4214d79688ab7697aca78f243d50?user_name=adamantiumnewvz",
			Name:        "f2pool",
			UserID:      "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
		},
	}

	for _, pool := range testPools {
		fmt.Println("scraping rewards for", pool.Name)
		provider, err := pool.NewProviderInterface()
		assert.NoError(t, err)

		rewards, err := provider.ScrapeDailyRewards(pool.ObserverURL, pool.ID)
		assert.NoError(t, err)
		assert.NotNil(t, rewards, "Rewards should not be nil")
		fmt.Printf("Rewards: %+v\n", rewards)
	}

}
