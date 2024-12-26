package f2pool

import (
	"fmt"
	"net/http"
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

// The home page of f2pool dashboard as we are parsing the html
func (provider *F2Pool) GetTotalsEndpoint(observerURL string) (string, error) {
	url, err := url.Parse(observerURL)
	if err != nil {
		return observerURL, err
	}

	splitPath := strings.Split(url.EscapedPath(), "/")

	accessKey := splitPath[2]
	accessKey = strings.TrimSpace(accessKey)
	if accessKey == "mining-user" || accessKey == "" {
		return observerURL, fmt.Errorf("f2pool url not parsed correctly format seems to have changed")
	}

	queryString := "?"
	userName := url.Query().Get("user_name")

	if userName != "" {
		queryString += "user_name=" + userName
	}

	return "https://www.f2pool.com/mining-user/" + accessKey + queryString, nil
}

func (provider *F2Pool) GetRewardsEndpoint(observerURL string) string {
	//?user_name=adamantiumnewvz&params=user_name=adamantiumnewvz&currency_code=btc&account=adamantiumnewvz&action=load_payout_history_income
	currURL, _ := url.Parse(observerURL)

	queries := currURL.Query()
	path := currURL.Path

	userName := queries.Get("user_name")
	queryStr := "?" + "user_name=" + userName + "&params=user_name=" + userName + "&currency_code=btc&action=load_payout_history_income"

	return "https://www.f2pool.com" + path + queryStr
}

// currency code should be variable
func (provider *F2Pool) GetPayoutsEndpoint(observerURL string) string {
	currURL, _ := url.Parse(observerURL)

	queries := currURL.Query()
	path := currURL.Path

	userName := queries.Get("user_name")
	queryStr := "?" + "user_name=" + userName + "&params=user_name=" + userName + "&currency_code=btc&action=load_payout_history_outcome"

	return "https://www.f2pool.com" + path + queryStr
}

func addTotalsHeaders(req *http.Request) {
	req.Header.Set("authority", "www.f2pool.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en_US")
	req.Header.Set("cookie", "lang=en_US")
	req.Header.Set("platform", "web")
	req.Header.Set("sec-ch-ua", "'Not_A Brand';v='8', 'Chromium';v='120', 'Brave';v='120'")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "'macOS'")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
}

func addRewardsHeaders(req *http.Request, observerURL string) {
	headers := map[string]string{
		"accept":             "*/*",
		"accept-language":    "en-US,en;q=0.9",
		"cache-control":      "no-cache",
		"dnt":                "1",
		"pragma":             "no-cache",
		"priority":           "u=1, i",
		"referer":            observerURL,
		"sec-ch-ua":          `"Chromium";v="131", "Not_A Brand";v="24"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"macOS"`,
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
		"x-requested-with":   "XMLHttpRequest",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
