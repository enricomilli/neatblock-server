package wallets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func TestWalletInfo(w http.ResponseWriter, r *http.Request) {

	url := "https://bitcoin.drpc.org"
	testAddr := "bc1q5dzz74g3k8keachlkv6r0jwk3lgeezdedplamr"

	// params := []any{6, 9999999, []string{testAddr}, true, map[string]interface{}{"minimumAmount": 0.005}}

	params := map[string]any{
		"minconf":        1,
		"maxconf":        999999,
		"addresses":      []string{testAddr},
		"include_unsafe": true,
	}

	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      01301,
		"method":  "listunspent",
		"params":  params,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println(result)
	fmt.Fprintf(w, "%+v", result)
}
