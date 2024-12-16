package f2pool

type F2PoolRewardsResponse struct {
	Status string           `json:"status"`
	Data   F2PoolRewardData `json:"data"`
}

type F2PoolRewardData struct {
	IncomeData       []F2PoolIncomeData     `json:"income_data"`
	FilterCommentMap F2PoolFilterCommentMap `json:"filter_comment_map"`
}

type F2PoolIncomeData struct {
	HashRate      string  `json:"hash_rate"`
	CreatedAt     int64   `json:"created_at"`
	Comment       string  `json:"comment"`
	FilterComment string  `json:"filter_comment"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	Txfee         float64 `json:"txfee"`
	CurrencyCode  string  `json:"currency_code"`
	Difficulty    string  `json:"difficulty"`
}

type F2PoolFilterCommentMap struct {
	CurrencyCode string            `json:"currency_code"`
	FilterMap    map[string]string `json:"filter_map"`
}

type F2PoolPayoutResponse struct {
	Status string             `json:"status"`
	Data   []F2PoolPayoutData `json:"data"`
}

type F2PoolPayoutData struct {
	CreatedAt    int64   `json:"created_at"`
	Address      string  `json:"address"`
	TxID         string  `json:"txid"`
	Chain        string  `json:"chain"`
	CurrencyCode string  `json:"currency_code"`
	Amount       string  `json:"amount"`
	Tax          float64 `json:"tax"`
	Type         string  `json:"type"`
	Comment      string  `json:"comment"`
}
