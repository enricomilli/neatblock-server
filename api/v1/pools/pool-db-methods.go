package pools

import (
	"fmt"
	"os"

	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
	"github.com/enricomilli/neat-server/db"
	"github.com/supabase-community/supabase-go"
)

// This method stores the current state of
// the pool struct into the db
func (pool *Pool) StorePoolStructState() error {

	// sbClient, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SERVICE_KEY"), &supabase.ClientOptions{})
	// if err != nil {
	// 	return fmt.Errorf("could not init supabase client: %w", err)
	// }

	// _, _, err = sbClient.From("pools").Upsert(pool, "id", "*", "exact").Execute()
	// if err != nil {
	// 	return fmt.Errorf("could not execute pool upsert: %w", err)
	// }

	database, err := db.NewClient()
	if err != nil {
		return fmt.Errorf("could not init db: %v", err)
	}

	query, values := db.BuildInsertQuery("pools", pool)

	_, err = database.NamedExec(query, values)
	if err != nil {
		return fmt.Errorf("could not store current pool state: %v", err)
	}

	return nil
}

// Gets all the rewards with the pool uid as a reference
func (pool *Pool) GetAllRewards() ([]poolproviders.MiningReward, error) {

	rewards := []poolproviders.MiningReward{}

	database, err := db.NewClient()
	if err != nil {
		return rewards, err
	}

	query := `
		select * from rewards where pool_id = $1
	`

	err = database.Select(&rewards, query, pool.ID)
	if err != nil {
		return rewards, fmt.Errorf("could not get all rewards: %v", err)
	}

	return rewards, nil
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
