package bitcoinrpc

type RPC interface {
	NewRPCClient() *RPC
	GetWalletBalance(walletAddr string) float64
}
