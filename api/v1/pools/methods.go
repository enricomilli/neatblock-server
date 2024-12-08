package pools

type PoolProvider interface {
	GetDailyReturnsEndpoint() string
	GetReturnsEndpoint() string
}
