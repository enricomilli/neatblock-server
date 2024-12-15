package pools

import (
	"fmt"
	"net/url"
	"strings"

	poolproviders "github.com/enricomilli/neat-server/api/v1/pools/providers"
	"github.com/enricomilli/neat-server/api/v1/pools/providers/viabtc"
	"golang.org/x/net/publicsuffix"
)

func (pool *Pool) GetProvider() (poolproviders.SupportedProvider, error) {
	poolURL, err := url.Parse(pool.ObserverURL)
	if err != nil {
		return "", fmt.Errorf("could not parse URL: %w", err)
	}

	domain, err := publicsuffix.EffectiveTLDPlusOne(poolURL.Hostname())
	if err != nil {
		return "", fmt.Errorf("could not extract domain: %w", err)
	}

	splitDomain := strings.Split(domain, ".")
	if len(splitDomain) < 1 {
		return "", fmt.Errorf("invalid domain format: %v", domain)
	}

	provider := poolproviders.SupportedProvider(strings.ToUpper(splitDomain[0]))

	if !provider.IsValid() {
		return "", fmt.Errorf("provider %s is not supported", provider)
	}

	return provider, nil
}

// A pool provider is an interface to interact with pool hosts
func (pool *Pool) NewProviderInterface() (poolproviders.PoolProvider, error) {

	provider, err := pool.GetProvider()
	if err != nil {
		return nil, fmt.Errorf("Error extracting provider: %w", err)
	}

	switch provider {
	case poolproviders.EnumViaBTC:
		return &viabtc.ViaBTC{}, nil
	case poolproviders.EnumF2Pool:
		return &viabtc.ViaBTC{}, nil
	default:
		return nil, fmt.Errorf("Pool provider %s is not supported.", provider)
	}

}
