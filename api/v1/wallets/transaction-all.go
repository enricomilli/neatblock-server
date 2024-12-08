package wallets

import (
	"fmt"
	"net/http"
	"os"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	apiutil "github.com/enricomilli/neat-server/api/api-utils"
	v1types "github.com/enricomilli/neat-server/api/v1/types"
)

func HandleAllTransactions(w http.ResponseWriter, r *http.Request) {

	body := v1types.WalletAllTransactionsRequest{}
	err := apiutil.StrictParseJSON(r, &body)
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusBadRequest, "Error parsing body:", err)
		return
	}

	// Configure RPC connection
	connCfg := &rpcclient.ConnConfig{
		Host:         os.Getenv("BTCD_ENDPOINT"),
		User:         os.Getenv("USERNAME"),
		Pass:         os.Getenv("PASSWORD"),
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	// Create new client
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusInternalServerError, "rpcclient error:", err)
		return
	}
	defer client.Shutdown()

	// Bitcoin address to check
	address, err := btcutil.DecodeAddress(body.WalletAddr, &chaincfg.MainNetParams)
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusBadRequest, "Invalid address:", err)
		return
	}

	// Get transactions for the address
	transactions, err := client.SearchRawTransactions(address, 0, 1000, true, nil)
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusInternalServerError, "Error fetching transactions:", err)
		return
	}

	// Print transaction details
	for _, tx := range transactions {
		fmt.Printf("Transaction ID: %s\n", tx.TxHash())
		fmt.Printf("------------------------\n")
	}
}
