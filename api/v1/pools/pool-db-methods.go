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

// Gets all the rewards with the pool uid as a reference
func (pool *Pool) GetAllRewards() ([]poolproviders.MiningReward, error) {
	fmt.Println("getting all rewards for pool:", pool.Name, "with id:", pool.ID)

	return []poolproviders.MiningReward{}, nil
}

// This function will store the list of rewards its passed, will error if the reward already exists
func (pool *Pool) StoreRewards(newRewards []poolproviders.MiningReward) error {
	sbClient, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SERVICE_KEY"), &supabase.ClientOptions{})
	if err != nil {
		return fmt.Errorf("could not init supabase client: %w", err)
	}

	for _, reward := range newRewards {
		_, _, err = sbClient.From("rewards").Insert(reward, false, "id", "*", "exact").Execute()
		if err != nil {
			return fmt.Errorf("could not execute pool upsert: %w", err)
		}
	}

	return nil
}
