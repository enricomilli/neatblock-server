package f2pool

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type F2PoolRewardsResponse struct {
	Status string           `json:"status"`
	Data   F2PoolRewardData `json:"data"`
}

type F2PoolRewardData struct {
	IncomeData       []F2PoolIncomeData     `json:"income_data"`
	FilterCommentMap F2PoolFilterCommentMap `json:"filter_comment_map"`
}

type F2PoolIncomeData struct {
	HashRate      HashRateField `json:"hash_rate"`
	CreatedAt     int64         `json:"created_at"`
	Comment       string        `json:"comment"`
	FilterComment string        `json:"filter_comment"`
	Type          string        `json:"type"`
	Amount        float64       `json:"amount"`
	Txfee         float64       `json:"txfee"`
	CurrencyCode  string        `json:"currency_code"`
	Difficulty    string        `json:"difficulty"`
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
	CreatedAt    int     `json:"created_at"`
	Address      string  `json:"address"`
	TxID         string  `json:"txid"`
	Chain        string  `json:"chain"`
	CurrencyCode string  `json:"currency_code"`
	Amount       string  `json:"amount"`
	Tax          float64 `json:"tax"`
	Type         string  `json:"type"`
	Comment      string  `json:"comment"`
}

type HashRateField float64

func (h *HashRateField) UnmarshalJSON(data []byte) error {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch v := raw.(type) {
	case float64:
		*h = HashRateField(v)
	case string:
		// Remove any non-numeric characters except decimal point
		numStr := strings.Split(v, " ")[0] // Take only the number part before space
		f, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return fmt.Errorf("failed to parse string to float64: %v", err)
		}
		*h = HashRateField(f)
	case int:
		*h = HashRateField(float64(v))
	default:
		return fmt.Errorf("unexpected type for HashRate: %T", v)
	}

	return nil
}
