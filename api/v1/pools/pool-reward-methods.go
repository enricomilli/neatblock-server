package pools

import (
	"fmt"

	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
	"github.com/enricomilli/neat-server/db"
)

// Gets all the rewards with the pool uid as a reference
func (pool *Pool) GetAllRewards() ([]poolproviders.MiningReward, error) {

	rewards := []poolproviders.MiningReward{}

	database, err := db.NewClient()
	if err != nil {
		return rewards, err
	}

	query := `
		select * from pool_rewards where pool_id = $1
	`

	err = database.Select(&rewards, query, pool.ID)
	if err != nil {
		return rewards, fmt.Errorf("could not get all rewards: %v", err)
	}

	return rewards, nil
}

// This function will store the list of rewards its passed, will error if the reward already exists
func (pool *Pool) StoreRewards(newRewards []poolproviders.MiningReward) error {
	// sbClient, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SERVICE_KEY"), &supabase.ClientOptions{})
	// if err != nil {
	// 	return fmt.Errorf("could not init supabase client: %w", err)
	// }

	// for _, reward := range newRewards {
	// 	_, _, err = sbClient.From("rewards").Insert(reward, false, "id", "*", "exact").Execute()
	// 	if err != nil {
	// 		return fmt.Errorf("could not execute pool upsert: %w", err)
	// 	}
	// }

	database, err := db.NewClient()
	if err != nil {
		return fmt.Errorf("could not init db client: %v", err)
	}

	query := `
		INSERT INTO pool_rewards (
		    id,
		    pool_id,
		    date,
		    hashrate,
		    btc_reward,
		    btc_tx_fee,
		    total,
		    payout
		)
		VALUES (
		    :id,
		    :pool_id,
		    :date,
		    :hashrate,
		    :btc_reward,
		    :btc_tx_fee,
		    :total,
		    :payout
		)
		ON CONFLICT (id) DO UPDATE SET
		    pool_id = EXCLUDED.pool_id,
		    date = EXCLUDED.date,
		    hashrate = EXCLUDED.hashrate,
		    btc_reward = EXCLUDED.btc_reward,
		    btc_tx_fee = EXCLUDED.btc_tx_fee,
		    total = EXCLUDED.total,
		    payout = EXCLUDED.payout,
		    updated_at = EXCLUDED.updated_at
    `

	for _, reward := range newRewards {
		_, err = database.NamedExec(query, reward)
		if err != nil {
			return fmt.Errorf("could not upsert reward: %v", err)
		}
	}

	return nil
}
