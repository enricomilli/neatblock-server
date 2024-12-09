package poolproviders

type ViaBTC struct {
	TotalsEndpoint  string
	RewardsEndpoint string
}

func (provider *ViaBTC) CompanyName() string {
	return "VIABTC"
}

func (provider *ViaBTC) ScrapeTotals(observerurl string) (float64, error) {

	return 0.11, nil
}

func (provider *ViaBTC) ScrapeDailyRewards(observerurl string) ([]string, error) {

	return []string{}, nil
}
