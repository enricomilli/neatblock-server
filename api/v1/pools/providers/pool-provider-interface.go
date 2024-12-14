package poolproviders

// The PoolProvider is an interface that works with all pools provider websites
type PoolProvider interface {
	CompanyName() string
	ValidateURL(observerURL string) error
	ScrapeTotals(observerURL string) (NeatblockTotals, error)         // todo create a unified type for totals
	ScrapeDailyRewards(observerURL string) ([]NeatblockReward, error) // todo create a unified type for a reward
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

type NeatblockTotals struct {
	TotalBtcProfit float64
}

type NeatblockReward struct {
	Date   string
	TxFee  float64
	Reward float64
	Payout float64
}
