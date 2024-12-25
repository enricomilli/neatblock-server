package pools

import "database/sql"

// row from the pools table
type Pool struct {
	ID                      string          `db:"id" json:"id"`
	ObserverURL             string          `db:"pool_url" json:"pool_url"`
	Status                  string          `db:"status" json:"status"`
	OwnerID                 string          `db:"user_id" json:"user_id"`
	Name                    string          `db:"name" json:"name"`
	BoughtHashrate          sql.NullFloat64 `db:"bought_hashrate" json:"bought_hashrate"`                   // added by user
	HighestHashrateAchieved float64         `db:"highest_achieve_hashrate" json:"highest_achieve_hashrate"` // scraped from pool data
	TotalBtcPayout          float64         `db:"total_btc_payout" json:"total_btc_payout"`                 // total btc paid out from mining
	TotalBtcMined           float64         `db:"total_btc_mined" json:"total_btc_mined"`                   // total mined + tx fees from mining
	TotalBtcSold            sql.NullFloat64 `db:"total_btc_sold" json:"total_btc_sold"`                     // null till they add realized gains
	TotalBtcExpensed        sql.NullFloat64 `db:"total_btc_expensed" json:"total_btc_expensed"`             // null till they add expenses
	TotalUsdGain            sql.NullFloat64 `db:"total_usd_gain" json:"total_usd_gain"`                     // null till a realized gain is added
	TotalUsdExpensed        sql.NullFloat64 `db:"total_usd_expensed" json:"total_usd_expensed"`             // null till they add an expense they paid in usd, meaning that it from outside this investment
	RevenueShare            sql.NullFloat64 `db:"revenue_share" json:"revenue_share"`                       // possibly null
	UptimePercent           float64         `db:"uptime_percent" json:"uptime_percent"`                     // possibly null because it requires hashrate, or we can make an estimate to begin with
	InitialInvestment       sql.NullFloat64 `db:"initial_investment" json:"initial_investment"`             // possibly null
	COC                     sql.NullFloat64 `db:"coc" json:"coc"`                                           // possibly null
	ROI                     sql.NullFloat64 `db:"roi" json:"roi"`                                           // possibly null
	BtcPriceAtDeployment    sql.NullFloat64 `db:"btc_price_at_deployment"`                                  // null till deployment date is added, or can confirm that the first reward is the deployment date
	DeploymentDate          sql.NullString  `db:"deployment_date" json:"deployment_date"`                   // null till deployment date is added, or can confirm that the first reward is the deployment date
	CreatedAt               string          `db:"created_at"`
	UpdatedAt               string          `db:"updated_at"`
}
