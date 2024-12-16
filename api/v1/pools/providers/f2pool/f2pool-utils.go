package f2pool

import (
	"fmt"
	"net/url"
	"strings"
)

func (provider *F2Pool) CompanyName() string {
	return "F2POOL"
}

// one url can have multiple user_name:
// example url #1: https://www.f2pool.com/mining-user/6f2e4214d79688ab7697aca78f243d50?user_name=adamantiumnewvz
// example url #2: https://www.f2pool.com/mining-user/6f2e4214d79688ab7697aca78f243d50?user_name=adamantiumnuno
func (provider *F2Pool) ValidateURL(observerURL string) error {
	if observerURL == "" {
		return fmt.Errorf("URL cannot be empty")
	}

	currURL, err := url.Parse(observerURL)
	if err != nil {
		return fmt.Errorf("could not parse url: %v", err)
	}

	// Check if domain is f2pool.com
	if !strings.Contains(currURL.Host, "f2pool.com") {
		return fmt.Errorf("invalid domain: must be f2pool.com")
	}

	// Check if path starts with /mining-user
	if !strings.HasPrefix(currURL.Path, "/mining-user") {
		return fmt.Errorf("invalid path: must start with /mining-user")
	}

	// Check if user_name parameter exists
	// queryParams := currURL.Query()
	// if !queryParams.Has("user_name") {
	// 	return fmt.Errorf("missing required parameter: user_name")
	// }

	// Check if user_name is not empty
	// userName := queryParams.Get("user_name")
	// if userName == "" {
	// 	return fmt.Errorf("user_name parameter cannot be empty")
	// }

	return nil
}
