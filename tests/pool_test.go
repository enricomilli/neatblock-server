package tests

import (
	"fmt"
	"testing"

	"github.com/enricomilli/neat-server/api/v1/pools"
	"github.com/stretchr/testify/assert"
)

func TestPoolFunctions(t *testing.T) {
	testPool := pools.Pool{
		ObserverURL: "https://www.viabtc.com/observer/dashboard?access_key=c0a80baded29bb989c06458a8d410a13&coin=BTC",
		Name:        "Test pool",
		OwnerID:     "d0d4ad7e-6c89-4a44-8e58-2b0d01170c2f",
	}

	provider, err := testPool.NewProviderInterface()
	assert.NoError(t, err)

	err = provider.ValidateURL(testPool.ObserverURL)
	assert.NoError(t, err)

	totals, err := provider.ScrapeTotals(testPool.ObserverURL)
	assert.NoError(t, err)

	fmt.Printf("%s Totals: %+v\n", testPool.Name, totals)
}
