package pools

import (
	"fmt"
	"os"

	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
	"github.com/supabase-community/supabase-go"
)

// This method stores the current state of
// the pool struct into the db
func (pool *Pool) StorePoolStructState() error {

	sbClient, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SERVICE_KEY"), &supabase.ClientOptions{})
	if err != nil {
		return fmt.Errorf("could not init supabase client: %w", err)
	}

	_, _, err = sbClient.From("pools").Upsert(pool, "id", "*", "exact").Execute()
	if err != nil {
		return fmt.Errorf("could not execute pool upsert: %w", err)
	}

	return nil
}

func (pool *Pool) GetAllRewards() ([]poolproviders.MiningReward, error) {

	return []poolproviders.MiningReward{}, nil
}

func (pool *Pool) StoreRewards([]poolproviders.MiningReward) error {

	return nil
}
