package tests

import (
	"fmt"
	"testing"

	"github.com/enricomilli/neat-server/api/v1/pools"
	"github.com/enricomilli/neat-server/db"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_PoolGetAllRewards(t *testing.T) {

	godotenv.Load("../../.env")

	database, err := db.NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, database)

	query := `
	select * from pools where id = $1;
	`
	testID := "828ff931-5255-4087-b4ba-5692aeb6364c"

	pool := &pools.Pool{}

	// The get method is for fetching a specific row
	err = database.Get(pool, query, testID)
	assert.NoError(t, err)

	allRewards, err := pool.GetAllRewards()
	assert.NoError(t, err)
	assert.NotNil(t, allRewards)

	fmt.Printf("Pools rewards: %+v\n", allRewards)
}
