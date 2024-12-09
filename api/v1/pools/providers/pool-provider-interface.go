package poolproviders

// The PoolProvider is an interface that works with all pools provider websites
type PoolProvider interface {
	CompanyName() string
	ScrapeTotals(observerUrl string) (float64, error)        // todo create a unified type for totals
	ScrapeDailyRewards(observerUrl string) ([]string, error) // todo create a unified type for a reward
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
