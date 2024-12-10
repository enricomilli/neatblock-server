package pools

type Pool struct {
	ID          string  `db:"id" json:"id"`
	ObserverURL string  `db:"pool_url" json:"pool_url"`
	Owner       string  `db:"user_id" json:"user_id"`
	Name        string  `db:"name" json:"name"`
	BTCRevenue  float64 `db:"total_btc_mined" json:"total_btc_mined"`
}

type PoolReward struct {
	Date   string
	Reward float64 // BTC reward
	TxFee  float64 // Transaction fee rewarded
	Payout float64 // Total payout
}
