package poolproviders

// The PoolProvider is an interface that works with all pools provider websites
type PoolProvider interface {
	CompanyName() string
	ValidateURL(observerURL string) error
	ScrapeTotals(observerURL string) (MiningTotals, error) // todo create a unified type for totals
	// requires pool id to add as a reference to each reward
	ScrapeDailyRewards(observerURL string, poolID string) ([]MiningReward, error)
}

type SupportedProvider string

const (
	EnumViaBTC SupportedProvider = "VIABTC"
	EnumF2Pool SupportedProvider = "F2POOL"
)

func (p SupportedProvider) IsValid() bool {
	switch p {
	case EnumViaBTC, EnumF2Pool:
		return true
	default:
		return false
	}
}

type MiningTotals struct {
	TotalBtcPayout float64
	TotalBtcMined  float64 `db:"total_btc_mined" json:"total_btc_mined"` // total mined + tx fees from mining
}

// row from rewards table
type MiningReward struct {
	ID        string  `db:"id"`
	PoolID    string  `db:"pool_id"`
	Date      string  `db:"date"`
	Hashrate  float64 `db:"hashrate"`
	BtcReward float64 `db:"btc_reward"` // BTC reward
	BtcTxFee  float64 `db:"btc_tx_fee"` // Transaction fee rewarded
	Total     float64 `db:"total"`      // BtcTxFee + BtcReward
	Payout    float64 `db:"payout"`     // Total payout
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
}
