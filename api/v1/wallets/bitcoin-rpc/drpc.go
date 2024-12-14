package bitcoinrpc

type DRPC struct{}

func (rpc *DRPC) NewRPCClient() *DRPC {
	return &DRPC{}
}
