package poolproviders

// The PoolProvider is an interface that works with all pools provider websites
type PoolProvider interface {
	CompanyName() string
	ValidateURL(observerURL string) error
	ScrapeTotals(observerURL string) (MiningTotals, error)         // todo create a unified type for totals
	ScrapeDailyRewards(observerURL string) ([]MiningReward, error) // todo create a unified type for a reward
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
	TotalBtcProfit float64
}

type MiningReward struct {
	Date         string
	Hashrate     float64
	TxFee        float64
	Reward       float64
	RewardPlusTx float64
	Payout       float64
}
