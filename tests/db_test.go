package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/enricomilli/neat-server/api/v1/pools"
	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
	"github.com/enricomilli/neat-server/db"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Ensure environment variables are set for tests
	requiredEnvVars := []string{
		"SUPABASE_URL",
		"SUPABASE_PORT",
		"SUPABASE_USER",
		"SUPABASE_PASSWORD",
		"SUPABASE_DB_NAME",
	}

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("could not load env vars:", err)
	}

	for _, env := range requiredEnvVars {
		if os.Getenv(env) == "" {
			panic("Required environment variable not set: " + env)
		}
	}

	os.Exit(m.Run())
}

func TestDatabaseConnection(t *testing.T) {

	db, err := db.NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Test the connection with a simple query
	var result int
	err = db.Get(&result, "SELECT 1")
	assert.NoError(t, err)
	assert.Equal(t, 1, result)
}

func TestPoolsTableExists(t *testing.T) {
	db, err := db.NewClient()
	assert.NoError(t, err)

	var exists bool
	err = db.Get(&exists, `
		SELECT EXISTS (
			SELECT FROM information_schema.tables
			WHERE table_schema = 'public'
			AND table_name = 'pools'
		)
	`)
	assert.NoError(t, err)
	assert.True(t, exists, "Pools table should exist")
}

func TestPoolRewardsTableExists(t *testing.T) {
	db, err := db.NewClient()
	assert.NoError(t, err)

	var exists bool
	err = db.Get(&exists, `
		SELECT EXISTS (
			SELECT FROM information_schema.tables
			WHERE table_schema = 'public'
			AND table_name = 'pool_rewards'
		)
	`)
	assert.NoError(t, err)
	assert.True(t, exists, "Pool rewards table should exist")
}

func TestGettingAllRewards(t *testing.T) {
	db, err := db.NewClient()
	assert.NoError(t, err)

	rewardsList := []poolproviders.MiningReward{}

	query := `
	select * from pool_rewards order by date ASC
	`

	err = db.Select(&rewardsList, query)
	assert.NoError(t, err)
}

func TestGettingAllPools(t *testing.T) {
	db, err := db.NewClient()
	assert.NoError(t, err)

	poolsList := []pools.Pool{}

	query := `
	select * from pools order by initial_investment ASC
	`

	err = db.Select(&poolsList, query)
	assert.NoError(t, err)

	fmt.Printf("all pools: %+v\n", poolsList)
}
