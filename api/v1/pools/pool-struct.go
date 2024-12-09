package pools

type Pool struct {
	UID         string
	ObserverURL string
	Owner       string
	BTCRevenue  float64
	BTCProfit   float64
}

type PoolReward struct {
	Date   string
	Reward float64 // BTC reward
	TxFee  float64 // Transaction fee rewarded
	Payout float64 // Total payout
}
