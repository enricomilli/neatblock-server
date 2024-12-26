package f2pool

import (
	"fmt"
	"testing"

	"github.com/enricomilli/neat-server/api/v1/pools/providers/f2pool"
	"github.com/stretchr/testify/assert"
)

// TODO: find more pool links and test them
func TestF2Pool_TotalsEndpointBuilder(t *testing.T) {
	provider := &f2pool.F2Pool{}
	testURL := "https://www.f2pool.com/mining-user/6f2e4214d79688ab7697aca78f243d50?user_name=adamantiumnewvz"

	totalsEndpoint, err := provider.GetTotalsEndpoint(testURL)
	assert.NoError(t, err)
	assert.NotNil(t, totalsEndpoint)

	fmt.Println("Test pool's endpoint:", totalsEndpoint)
}

func TestF2Pool_ScrapeTotals(t *testing.T) {
	provider := &f2pool.F2Pool{}
	testURL := "https://www.f2pool.com/mining-user/6f2e4214d79688ab7697aca78f243d50?user_name=adamantiumnewvz"

	totals, err := provider.ScrapeTotals(testURL)
	assert.NoError(t, err)
	assert.NotNil(t, totals)

	fmt.Printf("f2pool totals: %+v\n", totals)
}
