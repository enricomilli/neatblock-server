package viabtc

type ViaBTCTotalsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AccountBalance     string `json:"account_balance"`
		Profit24Hour       string `json:"profit_24hour"`
		ProfitTotal        string `json:"profit_total"`
		Hashrate10Min      string `json:"hashrate_10min"`
		Hashrate1Hour      string `json:"hashrate_1hour"`
		Hashrate1Day       string `json:"hashrate_1day"`
		TotalActive        int    `json:"total_active"`
		TotalUnactive      int    `json:"total_unactive"`
		TotalInvalid       int    `json:"total_invalid"`
		HashUnit           string `json:"hash_unit"`
		Target24Hour       string `json:"target_24hour"`
		GiftProfit24HourFb string `json:"gift_profit_24hour_fb"`
		GiftProfitTotalFb  string `json:"gift_profit_total_fb"`
	} `json:"data"`
}

type ViaBTCRewardsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		TotalPage int  `json:"total_page"`
		Total     int  `json:"total"`
		HasNext   bool `json:"has_next"`
		CurrPage  int  `json:"curr_page"`
		Count     int  `json:"count"`
		Data      []struct {
			Coin        string `json:"coin"`
			Date        string `json:"date"`
			Hashrate    string `json:"hashrate"`
			PayoutBTC   string `json:"total_profit"`
			RewardBTC   string `json:"pps_profit"`
			RewardTxFee string `json:"pplns_profit"`
			SoloProfit  string `json:"solo_profit"`
			PpsPlusRate string `json:"pps_plus_rate"`
			UnitOutput  string `json:"unit_output"`
			HashUnit    string `json:"hash_unit"`
			Currency    string `json:"currency"`
			Price       string `json:"price"`
			TotalMoney  string `json:"total_money"`
		} `json:"data"`
	} `json:"data"`
}
